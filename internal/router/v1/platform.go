package v1

import (
	"forward/internal/handler"

	"github.com/labstack/echo/v5"
)

func mountPlatformRouter(
	routerGrp *echo.Group,
	hdlr *handler.PlatformHandler,
) {
	platformGrp := routerGrp.Group("/platforms")

	platformGrp.POST("", hdlr.CreatePlatform)
	platformGrp.GET("", hdlr.GetAllPlatforms)
}
