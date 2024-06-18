package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/cjungo/cjuncms/controller"
	"github.com/cjungo/cjuncms/misc"
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
	"github.com/cjungo/cjungo/mid"
	"github.com/rs/zerolog"
)

var controllerProviders = []any{
	ext.NewCaptchaController,
	controller.NewMachineController,
	controller.NewSignController,
	controller.NewEmployeeController,
	controller.NewProjectController,
}

func route(
	logger *zerolog.Logger,
	router cjungo.HttpRouter,
	permitManager *mid.PermitManager[string, misc.EmployeeToken],
	captchaController *ext.CaptchaController,
	signController *controller.SignController,
	machineController *controller.MachineController,
	employeeController *controller.EmployeeController,
	projectController *controller.ProjectController,
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
	employeeGroup.PUT("/add", employeeController.Add, permitManager.Permit("employee_edit"))

	machineGroup := apiGroup.Group("/machine", permitManager.Permit("able"))
	machineGroup.GET("/cpu/info", machineController.PeekCpuInfo)
	machineGroup.GET("/cpu/times", machineController.PeekCpuTimes)
	machineGroup.GET("/virtual-memory", machineController.PeekVirtualMemory)

	return router.GetHandler(), nil
}
