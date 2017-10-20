package router

import "github.com/labstack/echo"

// Port is 占用服务端口号
func Port(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":1323"))
}
