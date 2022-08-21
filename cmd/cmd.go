package cmd

import (
	"behappy-url-shortener/src/model"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/builtin"
)

var RunServer = &gcli.Command{
	Name: "run",
	Desc: "启动短地址服务",
	Config: func(c *gcli.Command) {
		c.StrOpt(&model.RunOpts.Port, "port", "p", "3000", "Port number for the application")
		c.StrOpt(&model.RunOpts.RedisHost, "redis-host", "", "localhost", "Redis Server hostname")
		c.StrOpt(&model.RunOpts.RedisPort, "redis-Port", "", "6379", "Redis Server Port number")
		c.StrOpt(&model.RunOpts.RedisPass, "redis-pass", "", "", "Redis Server password")
		c.IntOpt(&model.RunOpts.RedisDb, "redis-db", "", 15, "Redis DB index")
	},
	Examples: `<cyan>{$binName} {$cmd} -u http://127.0.0.1:3000 -p 3000 --redis-host localhost --redis-Port 6379 --redis-pass 123456 --redis-db 15</>`,
	Func: func(c *gcli.Command, args []string) error {
		gcli.Println("感谢您的使用...")
		model.RedisInit()
		startApp()
		return nil
	},
}

func Command() {
	app := gcli.NewApp()
	app.Version = "1.0.0"
	app.Name = "BeHappy Url Shortener"
	app.Logo = &gcli.Logo{
		Text:  LOGO,
		Style: "info",
	}
	app.Desc = "基于go+redis实现的短地址服务"
	app.Add(RunServer)
	app.Add(builtin.GenAutoComplete())
	app.Run(nil)
}
