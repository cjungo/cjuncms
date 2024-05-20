package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cjungo/cjuncms/controller"
	"github.com/cjungo/cjuncms/entity"
	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/db"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	_ "github.com/cjungo/cjuncms/docs"
)

func route(
	logger *zerolog.Logger,
	router cjungo.HttpRouter,
	permitManager *mid.PermitManager[string, misc.EmployeeToken],
	captchaController *ext.CaptchaController,
	signController *controller.SignController,
	employeeController *controller.EmployeeController,
) (http.Handler, error) {
	here, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	publicDir := filepath.Join(here, "public")
	logger.Info().Str("dir", publicDir).Msg("静态目录")

	router.Static("/", publicDir)

	// 验证码
	captchaGroup := router.Group("/captcha")
	captchaGroup.GET("/math", captchaController.GenerateMath)

	// 登录
	signGroup := router.Group("/sign")
	signGroup.POST("/in", signController.SignIn)
	signGroup.GET("/out", signController.SignOut)

	// 接口 ==================================================
	apiGroup := router.Group("/api")

	employeeGroup := apiGroup.Group("/employee")
	employeeGroup.PUT("/add", employeeController.Add)

	return router.GetHandler(), nil
}

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
				return nil, err
			}
			return claims, nil
		})); err != nil {
			return err
		}

		// 提供数据库
		if err := container.Provide(db.NewMySqlHandle(func(mysql *db.MySql) error {
			entity.Use(mysql.DB)
			return mysql.Transaction(func(tx *gorm.DB) error {
				if err := misc.EnsurePermissions(tx); err != nil {
					return err
				}

				if err := misc.EnsureAdmin(tx); err != nil {
					return err
				}

				return nil
			})
		})); err != nil {
			return err
		}

		// 提供控制器
		if err := container.ProvideController([]any{
			ext.NewCaptchaController,
			controller.NewLoginController,
			controller.NewEmployeeController,
		}); err != nil {
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
