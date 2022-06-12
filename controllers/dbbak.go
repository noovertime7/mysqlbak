package controllers

import (
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/robfig/cron"
	"time"
)

type Baker interface {
	Run()
	AfterBakHandler() error
}

type BakHandler struct {
	Cron       *cron.Cron
	DbName     string
	dBInfo     *dBInfo
	Status     bool
	IsDingSend bool
	IsOssSend  bool
}

//func NewBaker( db *modles.Database) *BakHandler{
//  return &BakHandler{bakdatabase: db}
//}

func (b *BakHandler) Run() {
	//log.Logger.Info("开始备份", b.DbName)
	log.Logger.Info("开始备份", b.dBInfo.database.DBName)
	time.Sleep(2 * time.Second)
	log.Logger.Info("备份成功", b.DbName)
	log.Logger.Debugf("当前数据库:%s,开关状态:%v %v", b.DbName, b.IsDingSend, b.IsOssSend)
	b.Status = true
	err := b.AfterBakHandler()
	if err != nil {
		log.Logger.Error(err)
		return
	}
	//SendDingMsg()
}

func (b *BakHandler) AfterBakHandler() error {
	if b.Status {
		if b.IsDingSend {
			SendDingMsg(b.DbName + "备份成功")
			return nil
		}
	}
	return nil
}
