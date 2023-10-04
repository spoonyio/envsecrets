package auth

import (
	"github.com/labstack/echo/v4"
)

func AddRoutes(sg *echo.Group) {

	group := sg.Group("/auth")

	group.GET("/qr", GenerateQRHandler)
	group.POST("/mfa", ToggleMFAHandler)
	group.DELETE("/mfa", ToggleMFAHandler)
	group.POST("/signin", SigninHandler)
	group.POST("/signup", SignupHandler)
	group.POST("/update-password", UpdatePasswordHandler)
}
