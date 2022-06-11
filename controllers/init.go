package controllers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/mysqlbak/modles"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/robfig/cron/v3"
)

type dBInfo struct {
	engine   *xorm.Engine
	database *modles.Database
	CronJob  map[string]*cron.Cron
}

func NewDBInfo(db *modles.Database) *dBInfo {
	en, err := xorm.NewEngine("mysql", db.User+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.DBName+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Logger.Error("配置数据库引擎失败", err)
		return nil
	}
	fmt.Println(db.DingConf.IsDingSend)
	c := cron.New(cron.WithSeconds())
	jobmap := make(map[string]*cron.Cron)
	jobmap[db.DBName] = c
	return &dBInfo{
		engine:   en,
		database: db,
		CronJob:  jobmap,
	}
}

func (d *dBInfo) StartBak() {
	//c := cron.New(cron.WithSeconds())
	for dbname, c := range d.CronJob {
		cid, err := c.AddJob(d.database.BackupCycle,
			&BakHandler{DbName: dbname, dBInfo: d, IsDingSend: d.database.DingConf.IsDingSend,
				IsOssSend: d.database.OssConf.IsOssSave})
		if err != nil {
			log.Logger.Error("定时任务添加失败", err)
			return
		}
		log.Logger.Infof("启动任务:%d,备份数据库:%s,备份周期:%s,数据保留周期:%d天", cid, dbname, d.database.BackupCycle, d.database.KeepNumber)
		c.Start()
	}
}
