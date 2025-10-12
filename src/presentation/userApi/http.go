package userApi

import (
	"net/http"

	"github.com/caiiomp/vehicle-resale-auth/src/core/responses"
	"github.com/caiiomp/vehicle-resale-auth/src/core/useCases/user"
	"github.com/gin-gonic/gin"
)

type userApi struct {
	userService user.UserService
}

func RegisterUserRoutes(app *gin.Engine, userService user.UserService) {
	service := userApi{
		userService: userService,
	}

	app.POST("/users", service.create)
	app.GET("/users", service.search)
	app.GET("/users/:user_id", service.get)
}

func (ref *userApi) create(ctx *gin.Context) {
	var request createUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ref.userService.Create(ctx, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.UserFromDomain(*user)
	ctx.JSON(http.StatusCreated, response)
}

func (ref *userApi) search(ctx *gin.Context) {
	users, err := ref.userService.Search(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]userResponse, len(users))

	for i, user := range users {
		response[i] = userResponseFromDomain(user)
	}

	ctx.JSON(http.StatusOK, response)
}

func (ref *userApi) get(ctx *gin.Context) {
	var uri userURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ref.userService.GetByID(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := userResponseFromDomain(*user)
	ctx.JSON(http.StatusOK, response)
}
