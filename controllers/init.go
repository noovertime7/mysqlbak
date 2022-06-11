package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/noovertime7/mysqlbak/modles"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/robfig/cron"
)

type dBInfo struct {
	engine   *xorm.Engine
	database *modles.Database
	Cron     *cron.Cron
}

func NewDBInfo(db *modles.Database) *dBInfo {
	c := cron.New()
	//if err := c.AddFunc(db.BackupCycle, func() {
	//	fmt.Println("测试")
	//});err != nil {
	//	log.Logger.Error("创建定时任务失败")
	//}
	err := c.AddJob(db.BackupCycle, BakHandler{Bakdatabase: db})
	if err != nil {
		log.Logger.Error(err)
		return nil
	}
	en, err := xorm.NewEngine("mysql", db.User+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.DBName+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Logger.Error("配置数据库引擎失败", err)
		return nil
	}
	return &dBInfo{
		engine:   en,
		database: db,
		Cron:     c,
	}
}

func (d *dBInfo) StartBak() {
	d.Cron.Start()
	select {}
}
