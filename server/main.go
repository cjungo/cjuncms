package main

import (
	"log"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"

	_ "github.com/cjungo/cjuncms/docs"
)

func init() {
	if err := cjungo.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if app, err := cjungo.NewApplication(func(container cjungo.DiContainer) error {
		if err := container.Provides(
			cjungo.LoadLoggerConfFromEnv,     // 加载日志配置
			cjungo.LoadHttpServerConfFromEnv, // 加载服务器配置
			db.LoadMySqlConfFormEnv,          // 加载数据库配置
			ext.NewStorageManager,            // 存储管理
			misc.ProvideMysqlForWeb(),        // 提供数据库
			misc.NewJwtClaimsManager,         // Jwt 管理器
			misc.ProvidePermitManager(),      // 提供权限管理器
			misc.ProvideMachineWatcher,       // 机器监控器
			ext.ProvideMessageController(&ext.MessageControllerProviderConf[string]{ // 提供消息控制器
				TokenAccess: func(ctx cjungo.HttpContext) (string, error) {
					return ctx.GetReqID(), nil
				},
			}),
			route, // 提供路由
		); err != nil {
			return err
		}

		// 提供控制器
		return container.Provides(controllerProviders...)
	}); err != nil {
		log.Fatalln(err)
	} else {
		app.BeforeRun = func(dc cjungo.DiContainer) error {
			return dc.Invoke(misc.MachineTick)
		}
		if err := app.Run(); err != nil {
			log.Fatalln(err)
		}
	}
}
