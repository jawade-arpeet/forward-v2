package handler

import (
	"forward/internal/dto"
	"forward/internal/service"
	"net/http"

	"github.com/labstack/echo/v5"
)

type PlatformHandler struct {
	platformSvc *service.PlatformService
}

func newPlatformHandler(platformSvc *service.PlatformService) *PlatformHandler {
	return &PlatformHandler{platformSvc: platformSvc}
}

func (h *PlatformHandler) CreatePlatform(ctx *echo.Context) error {
	var req dto.CreatePlatformRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "invalid payload"},
		)
	}

	if err := h.platformSvc.CreatePlatform(
		ctx.Request().Context(),
		req.Name,
		req.LogoURL,
	); err != nil {
		return err
	}

	return ctx.JSON(
		http.StatusCreated,
		map[string]string{"message": "successfully created platform"},
	)
}

func (h *PlatformHandler) GetAllPlatforms(ctx *echo.Context) error {
	platforms, err := h.platformSvc.GetAllPlatforms(ctx.Request().Context())
	if err != nil {
		return err
	}

	resp := make([]*dto.CreatePlatformResponse, 0, len(platforms))
	for _, p := range platforms {
		resp = append(
			resp,
			dto.NewCreatePlatformResponse(p.ID, p.Name, p.LogoURL),
		)
	}

	return ctx.JSON(http.StatusOK, resp)
}
