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
	CronJob  map[string]string
}

func NewDBInfo(db *modles.Database) *dBInfo {
	en, err := xorm.NewEngine("mysql", db.User+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.DBName+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Logger.Error("配置数据库引擎失败", err)
		return nil
	}
	jobmap := make(map[string]string)
	jobmap[db.DBName] = db.BackupCycle
	return &dBInfo{
		engine:   en,
		database: db,
		CronJob:  jobmap,
	}
}

func (d *dBInfo) StartBak() {
	c := cron.New(cron.WithSeconds())
	fmt.Println(d.CronJob)
	fmt.Println(d.database.DBName)
	cid, err := c.AddJob(d.database.BackupCycle, BakHandler{DbName: d.database.DBName})
	if err != nil {
		log.Logger.Error(err)
		return
	}
	log.Logger.Info(cid)
	c.Start()
}
