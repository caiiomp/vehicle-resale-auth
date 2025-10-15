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

// Create godoc
// @Summary Create User
// @Description Create an user
// @Tags User
// @Accept json
// @Produce json
// @Param user body userApi.createUserRequest true "Body"
// @Success 201 {object} responses.User
// @Failure 204 {object} responses.ErrorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /users [post]
func (ref *userApi) create(ctx *gin.Context) {
	var request createUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := ref.userService.Create(ctx, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.UserFromDomain(*user)
	ctx.JSON(http.StatusCreated, response)
}

// Create godoc
// @Summary Search users
// @Description Seach users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} responses.User
// @Failure 500 {object} responses.ErrorResponse
// @Router /users [get]
func (ref *userApi) search(ctx *gin.Context) {
	users, err := ref.userService.Search(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := make([]responses.User, len(users))

	for i, user := range users {
		response[i] = responses.UserFromDomain(user)
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Get User
// @Description Get an user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} responses.User
// @Failure 204 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /users/{user_id} [get]
func (ref *userApi) get(ctx *gin.Context) {
	var uri userURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := ref.userService.GetByID(ctx, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.UserFromDomain(*user)
	ctx.JSON(http.StatusOK, response)
}
