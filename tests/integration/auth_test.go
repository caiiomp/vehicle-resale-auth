//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiiomp/vehicle-resale-auth/src/core/responses"
	"github.com/caiiomp/vehicle-resale-auth/src/core/useCases/auth"
	"github.com/caiiomp/vehicle-resale-auth/src/core/useCases/user"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation/authApi"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation/userApi"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/memory/userRepository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	userRepository := userRepository.NewUserRepository()

	authService := auth.NewAuthService(userRepository, "123")
	userService := user.NewUserService(validator.New(), userRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

	authApi.RegisterAuthRoutes(app, authService)
	userApi.RegisterUserRoutes(app, userService)

	payload := map[string]any{
		"name":     "John Doe",
		"email":    "john.doe@email.com",
		"password": "123456",
	}

	rawPayload, _ := json.Marshal(payload)
	body := bytes.NewReader(rawPayload)

	req, _ := http.NewRequest(http.MethodPost, "/users", body)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	payload = map[string]any{
		"email":    "john.doe@email.com",
		"password": "123456",
	}

	rawPayload, _ = json.Marshal(payload)
	body = bytes.NewReader(rawPayload)

	req, _ = http.NewRequest(http.MethodPost, "/login", body)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response responses.Login
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response.AccessToken)
	assert.NotZero(t, response.ExpiresIn)
}
