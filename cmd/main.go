package main

import (
	"github.com/noovertime7/mysqlbak/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
