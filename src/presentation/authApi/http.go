package authApi

import (
	"net/http"

	interfaces "github.com/caiiomp/vehicle-resale-auth/src/core/_interfaces"
	"github.com/caiiomp/vehicle-resale-auth/src/core/responses"
	"github.com/gin-gonic/gin"
)

type authApi struct {
	authService interfaces.AuthService
}

func RegisterAuthRoutes(app *gin.Engine, authService interfaces.AuthService) {
	service := authApi{
		authService: authService,
	}

	app.POST("/login", service.login)
}

// Create godoc
// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Produce json
// @Param login body authApi.loginRequest true "Body"
// @Success 200 {object} responses.Login
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /login [post]
func (ref *authApi) login(ctx *gin.Context) {
	var request loginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	auth, err := ref.authService.Login(ctx, request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if auth == nil {
		ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "email and password does not match",
		})
		return
	}

	response := responses.LoginFromDomain(*auth)
	ctx.JSON(http.StatusOK, response)
}
