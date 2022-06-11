package controllers

import (
	"github.com/noovertime7/mysqlbak/modles"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"time"
)

type Baker interface {
	Run()
}

type BakHandler struct {
	Bakdatabase *modles.Database
}

//func NewBaker( db *modles.Database) *BakHandler{
//  return &BakHandler{bakdatabase: db}
//}

func (b BakHandler) Run() {
	log.Logger.Info("开始备份", b.Bakdatabase.DBName)
	time.Sleep(2 * time.Second)
	log.Logger.Info("备份成功")
}
