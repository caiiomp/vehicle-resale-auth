package authApi

import (
	"net/http"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/useCases/auth"
	"github.com/caiiomp/vehicle-resale-auth/src/core/responses"
	"github.com/gin-gonic/gin"
)

type authApi struct {
	authService auth.AuthService
}

func RegisterAuthRoutes(app *gin.Engine, authService auth.AuthService) {
	service := authApi{
		authService: authService,
	}

	app.POST("/login", service.login)
}

func (ref *authApi) login(ctx *gin.Context) {
	var request loginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := ref.authService.Login(ctx, request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if auth == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "email and password does not match"})
		return
	}

	response := responses.LoginResponseFromDomain(*auth)
	ctx.JSON(http.StatusOK, response)
}
