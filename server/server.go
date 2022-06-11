package server

import (
	"github.com/noovertime7/mysqlbak/controllers"
	"github.com/noovertime7/mysqlbak/modles"
)

func Start() {
	for _, db := range modles.SystemConf.Database {
		dbi := controllers.NewDBInfo(&db)
		dbi.StartBak()
	}
	select {}
}
