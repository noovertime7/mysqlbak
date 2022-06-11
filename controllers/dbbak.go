package controllers

import (
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/robfig/cron"
	"time"
)

type Baker interface {
	Run()
}

type BakHandler struct {
	Cron   *cron.Cron
	DbName string
}

//func NewBaker( db *modles.Database) *BakHandler{
//  return &BakHandler{bakdatabase: db}
//}

func (b BakHandler) Run() {
	log.Logger.Info("开始备份", b.DbName)
	time.Sleep(2 * time.Second)
	log.Logger.Info("备份成功", b.DbName)
}
