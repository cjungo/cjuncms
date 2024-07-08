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
	controller.NewPassController,
	// controller.NewShellController,
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
	passController *controller.PassController,
) (http.Handler, error) {
	here, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// 静态目录
	publicDir := filepath.Join(here, "public")
	logger.Info().Str("dir", publicDir).Str("action", "静态目录").Msg("[HTTP]")
	router.Static("/", publicDir)

	// 验证码
	captchaGroup := router.Group("/captcha")
	captchaGroup.GET("/math", captchaController.GenerateMath)

	// 登录
	signGroup := router.Group("/sign")
	signGroup.POST("/in", signController.SignIn)
	signGroup.GET("/out", signController.SignOut)
	signGroup.GET("/renewal", signController.SignRenewal, permitManager.Permit("default"))

	// 接口 ==================================================
	apiGroup := router.Group("/api")

	employeeGroup := apiGroup.Group("/employee")
	employeeGroup.GET("/query", employeeController.Query, permitManager.Permit("employee_find"))
	employeeGroup.PUT("/add", employeeController.Add, permitManager.Permit("employee_edit"))
	employeeGroup.POST("/edit", employeeController.Edit, permitManager.Permit("employee_edit"))
	employeeGroup.DELETE("/drop", employeeController.Drop, permitManager.Permit("employee_edit"))

	projectGroup := apiGroup.Group("/project")
	projectGroup.GET("/query", projectController.Query, permitManager.Permit("project_find"))
	projectGroup.PUT("/add", projectController.Add, permitManager.Permit("project_edit"))
	projectGroup.POST("/edit", projectController.Edit, permitManager.Permit("project_edit"))
	projectGroup.DELETE("/drop", projectController.Drop, permitManager.Permit("project_edit"))
	passGroup := apiGroup.Group("/pass")
	passGroup.GET("/query", passController.Query, permitManager.Permit("project_find"))
	passGroup.PUT("/add", passController.Add, permitManager.Permit("project_edit"))

	machineGroup := apiGroup.Group("/machine", permitManager.Permit("default"))
	machineGroup.GET("/cpu/info", machineController.PeekCpuInfo)
	machineGroup.GET("/cpu/times", machineController.PeekCpuTimes)
	machineGroup.GET("/cpu/times/list", machineController.ListCpuTimes)
	machineGroup.GET("/virtual-memory", machineController.PeekVirtualMemory)
	machineGroup.GET("/virtual-memory/list", machineController.ListVirtualMemory)
	machineGroup.GET("/disk", machineController.PeekDiskUsage)
	machineGroup.GET("/processes", machineController.PeekProcesses)
	machineGroup.GET("/processes/list", machineController.ListProcesses)

	return router.GetHandler(), nil
}
