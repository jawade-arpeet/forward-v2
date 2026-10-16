package handler

import (
	"forward/internal/config"
	"forward/internal/constants"
	"forward/internal/dto"
	"forward/internal/service"
	"net/http"

	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	authSvc *service.AuthService
}

func newAuthHandler(authSvc *service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

func (h *AuthHandler) SetAuthCookie(ctx *echo.Context, accID string) {
	cfg := config.GetServerConfig()

	ctx.SetCookie(&http.Cookie{
		Name:     "forward_auth",
		Value:    accID,
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
		Secure:   cfg.RunEnv != constants.EnvDev,
		SameSite: http.SameSiteLaxMode,
	})
}

func (h *AuthHandler) SignUp(ctx *echo.Context) error {
	var req dto.SignUpRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "invalid payload"},
		)
	}

	if err := h.authSvc.SignUp(
		ctx.Request().Context(),
		req.Email,
		req.Password,
	); err != nil {
		return err
	}

	return ctx.JSON(
		http.StatusCreated,
		map[string]string{"message": "account sucessfully signed up"},
	)
}

func (h *AuthHandler) SignIn(ctx *echo.Context) error {
	var req dto.SignInRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "invalid payload"},
		)
	}

	accID, err := h.authSvc.SignIn(
		ctx.Request().Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		return err
	}

	h.SetAuthCookie(ctx, accID)

	return ctx.JSON(
		http.StatusOK,
		map[string]string{"message": "account sucessfully signed in"},
	)
}
