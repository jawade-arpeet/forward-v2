package v1

import (
	"forward/internal/handler"

	"github.com/labstack/echo/v5"
)

func mountAuthRouter(
	routerGrp *echo.Group,
	hdlr *handler.AuthHandler,
) {
	authGrp := routerGrp.Group("/auth")

	authGrp.GET("/sign-up", hdlr.SignUp)
	authGrp.GET("/sign-in", hdlr.SignIn)
}
