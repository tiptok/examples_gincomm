package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/tiptok/gocomm/config"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gocomm/pkg/orm"
	"github.com/tiptok/gocomm/pkg/redis"

	"github.com/tiptok/examples_gincomm/routers"
)

func init() {
	config.NewViperConfig("yaml", "conf/app-dev.yaml")
}

func main() {
	log.InitGinLog(config.Logger{
		Filename: "app.log",
		Level:    "7",
	})

	err := redis.InitWithDb(100, config.Default.String("redis_url"), config.Default.String("redis_auth"), "0")
	if err != nil {
		log.Fatal(err)
	}
	orm.NewBeeormEngine(config.Mysql{
		DataSource: config.Default.String("mysql_url"),
		MaxIdle:    100,
		MaxOpen:    100,
	})
	defer func() {
		log.Info("app on stop!")
	}()
	routers.InitRouter(config.Default.String("listen_addr"))
}
