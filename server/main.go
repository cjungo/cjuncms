package main

import (
	"log"

	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"

	_ "github.com/cjungo/cjuncms/docs"
)

func init() {
	if err := cjungo.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if app, err := cjungo.NewApplication(func(container cjungo.DiContainer) error {
		// 加载日志配置
		if err := container.Provide(cjungo.LoadLoggerConfFromEnv); err != nil {
			return err
		}
		// 加载服务器配置
		if err := container.Provide(cjungo.LoadHttpServerConfFromEnv); err != nil {
			return err
		}
		// 加载数据库配置
		if err := container.Provide(db.LoadMySqlConfFormEnv); err != nil {
			return err
		}

		// 提供权限管理器
		if err := container.Provide(mid.NewPermitManager(func(ctx cjungo.HttpContext) (mid.PermitProof[string, misc.EmployeeToken], error) {
			claims := &misc.JwtClaims{}
			if _, err := ext.ParseJwtToken(ctx, claims); err != nil {
				return nil, &cjungo.ApiError{
					Code:     misc.API_ERR_TOKEN_INVALID,
					Message:  "TOKEN 无效",
					HttpCode: 400,
					Reason:   err,
				}
			}
			return claims, nil
		})); err != nil {
			return err
		}
		if err := container.Provide(misc.NewJwtClaimsManager); err != nil {
			return err
		}

		// 提供数据库
		if err := container.Provide(misc.ProvideMysqlForWeb()); err != nil {
			return err
		}

		// 提供控制器
		if err := container.Provides(controllerProviders...); err != nil {
			return err
		}
		// 提供路由
		if err := container.Provide(route); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Fatalln(err)
	} else {
		if err := app.Run(); err != nil {
			log.Fatalln(err)
		}
	}
}
