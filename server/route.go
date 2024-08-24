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
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

var controllerProviders = []any{
	ext.NewCaptchaController,
	controller.NewMachineController,
	controller.NewSignController,
	controller.NewEmployeeController,
	controller.NewProjectController,
	controller.NewPassController,
	controller.NewDemandController,
	controller.NewDepartmentController,
	// controller.NewShellController,
}

func route(
	logger *zerolog.Logger,
	router cjungo.HttpRouter,
	storageManager *ext.StorageManager,
	permitManager *mid.PermitManager[string, misc.EmployeeToken],
	captchaController *ext.CaptchaController,
	messageController *misc.MessageController,
	shellMessageController *misc.ShellMessageController,
	signController *controller.SignController,
	machineController *controller.MachineController,
	employeeController *controller.EmployeeController,
	projectController *controller.ProjectController,
	passController *controller.PassController,
	demandController *controller.DemandController,
	departmentController *controller.DepartmentController,
) (http.Handler, error) {
	here, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// 静态目录
	publicDir := filepath.Join(here, "public")
	logger.Info().
		Str("public", publicDir).
		Str("action", "静态目录").
		Msg("[HTTP]")
	router.Static("/", publicDir)

	// 存储
	uploadDir := filepath.Join(here, "upload")
	logger.Info().
		Str("upload", uploadDir).
		Str("action", "私有存储").
		Msg("[HTTP]")
	storageManager.Route(router, &ext.StorageConf{
		PathPrefix:       "/upload",
		Dir:              uploadDir,
		UploadMiddleware: []echo.MiddlewareFunc{permitManager.Permit(misc.PERMIT_DEFAULT)},
		IndexMiddleware:  []echo.MiddlewareFunc{permitManager.Permit(misc.PERMIT_DEFAULT)},
		// TODO 访问权限
		// QueryMiddleware:  []echo.MiddlewareFunc{permitManager.Permit(misc.PERMIT_DEFAULT)},
	})

	// 消息
	router.GET("/msg", messageController.Dispatch, permitManager.Permit(misc.PERMIT_DEFAULT))

	// 验证码
	captchaGroup := router.Group("/captcha")
	captchaGroup.GET("/math", captchaController.GenerateMath)

	// 登录
	signGroup := router.Group("/sign")
	signGroup.POST("/in", signController.SignIn)
	signGroup.POST("/out", signController.SignOut)
	signGroup.GET("/renewal", signController.SignRenewal, permitManager.Permit(misc.PERMIT_DEFAULT))
	signGroup.GET("/profile", signController.GetProfile, permitManager.Permit(misc.PERMIT_DEFAULT))
	signGroup.POST("/profile", signController.SetProfile, permitManager.Permit(misc.PERMIT_DEFAULT))

	// TODO 鉴权, EventSource 不能修改头部，只能通过 QueryString 或者 Cookie 传递 jwt
	// sse 接口
	sseGroup := router.Group("/sse")
	// machineSseGroup := sseGroup.Group("/machine", permitManager.Permit(misc.PERMIT_DEFAULT))
	machineSseGroup := sseGroup.Group("/machine")
	machineSseGroup.SSE("/cpu/info", machineController.WatchCpuInfo)
	machineSseGroup.SSE("/cpu/times", machineController.WatchCpuTimes)
	machineSseGroup.SSE("/cpu/timeline", machineController.WatchCpuTimesTimeline)
	machineSseGroup.SSE("/virtual-memory", machineController.WatchVirtualMemory)
	machineSseGroup.SSE("/disk", machineController.WatchDiskUsage)

	// 接口 ==================================================
	apiGroup := router.Group("/api", permitManager.Permit(misc.PERMIT_DEFAULT))

	// 员工
	employeeGroup := apiGroup.Group("/employee", permitManager.Permit(misc.PERMIT_EMPLOYEE))
	employeeGroup.POST("/query", employeeController.Query)
	employeeGroup.POST("/add", employeeController.Add, permitManager.Permit(misc.PERMIT_EMPLOYEE_EDIT))
	employeeGroup.POST("/edit", employeeController.Edit, permitManager.Permit(misc.PERMIT_EMPLOYEE_EDIT))
	employeeGroup.DELETE("/drop", employeeController.Drop, permitManager.Permit(misc.PERMIT_EMPLOYEE_EDIT))

	// 部门
	departmentGroup := apiGroup.Group("/department", permitManager.Permit(misc.PERMIT_EMPLOYEE))
	departmentGroup.POST("/query", departmentController.Query)
	departmentGroup.POST("/add", departmentController.Add, permitManager.Permit(misc.PERMIT_EMPLOYEE_DEPARTMENT_EDIT))
	departmentGroup.POST("/edit", departmentController.Edit, permitManager.Permit(misc.PERMIT_EMPLOYEE_DEPARTMENT_EDIT))
	departmentGroup.POST("/drop", departmentController.Drop, permitManager.Permit(misc.PERMIT_EMPLOYEE_DEPARTMENT_EDIT))

	// 项目
	projectGroup := apiGroup.Group("/project", permitManager.Permit(misc.PERMIT_PROJECT))
	projectGroup.POST("/query", projectController.Query)
	projectGroup.POST("/add", projectController.Add, permitManager.Permit(misc.PERMIT_PROJECT_EDIT))
	projectGroup.POST("/edit", projectController.Edit, permitManager.Permit(misc.PERMIT_PROJECT_EDIT))
	projectGroup.DELETE("/drop", projectController.Drop, permitManager.Permit(misc.PERMIT_PROJECT_EDIT))
	passGroup := apiGroup.Group("/pass", permitManager.Permit(misc.PERMIT_PROJECT, misc.PERMIT_PASS))
	passGroup.POST("/query", passController.Query)
	passGroup.POST("/add", passController.Add, permitManager.Permit(misc.PERMIT_PASS_EDIT))

	// 机器
	machineGroup := apiGroup.Group("/machine", permitManager.Permit(misc.PERMIT_DEFAULT))
	machineGroup.GET("/cpu/info", machineController.PeekCpuInfo)
	machineGroup.GET("/cpu/times", machineController.PeekCpuTimes)
	machineGroup.POST("/cpu/times/list", machineController.ListCpuTimes)
	machineGroup.GET("/virtual-memory", machineController.PeekVirtualMemory)
	machineGroup.GET("/virtual-memory/list", machineController.ListVirtualMemory)
	machineGroup.GET("/disk", machineController.PeekDiskUsage)
	machineGroup.GET("/processes", machineController.PeekProcesses)
	machineGroup.GET("/processes/list", machineController.ListProcesses)

	// 命令行
	shellGroup := apiGroup.Group("/shell", permitManager.Permit(misc.PERMIT_DEFAULT))
	shellGroup.POST("/msg", shellMessageController.Dispatch)

	return router.GetHandler(), nil
}
