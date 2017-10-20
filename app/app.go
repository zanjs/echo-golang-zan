package app

import (
	"flag"
	"fmt"
	"os"
	Router "zan/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lunny/log"
)

var e *echo.Echo

// GetE is 初始化
func GetE() *echo.Echo {
	return e
}

// Start 运行启动
func Start() {
	e = echo.New()

	loggerConfig := middleware.DefaultLoggerConfig

	f, err := os.OpenFile("api.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("cannot open 'api.log', (%s)", err.Error())
		flag.Usage()
		os.Exit(-1)
	}
	defer f.Close()

	fmt.Println(f)
	fmt.Println(loggerConfig)

	loggerConfig.Output = f
	// Middleware
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Router.Register(e)
	Router.Port(e)
}
