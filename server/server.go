package server

import (
	"fmt"
	"github.com/noovertime7/mysqlbak/controllers"
	"github.com/noovertime7/mysqlbak/modles"
)

func Start() {
	for _, db := range modles.SystemConf.Database {
		dbi := controllers.NewDBInfo(&db)
		fmt.Println(db.DBName)
		dbi.StartBak()
	}
}
