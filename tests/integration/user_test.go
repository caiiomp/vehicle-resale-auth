//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiiomp/vehicle-resale-auth/src/core/responses"
	"github.com/caiiomp/vehicle-resale-auth/src/core/useCases/user"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation/userApi"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/memory/userRepository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	userRepository := userRepository.NewUserRepository()

	userService := user.NewUserService(validator.New(), userRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

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

	var response responses.User
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response.ID)
	assert.Equal(t, "John Doe", response.Name)
	assert.Equal(t, "john.doe@email.com", response.Email)
}

func TestSearchUsers(t *testing.T) {
	userRepository := userRepository.NewUserRepository()

	userService := user.NewUserService(validator.New(), userRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

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

	var response responses.User
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	userID := response.ID

	req, _ = http.NewRequest(http.MethodGet, "/users", body)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var users []responses.User
	err = json.Unmarshal(resp.Body.Bytes(), &users)
	require.NoError(t, err)

	assert.NotEmpty(t, response.ID)
	assert.Equal(t, userID, users[0].ID)
	assert.Equal(t, "John Doe", users[0].Name)
	assert.Equal(t, "john.doe@email.com", users[0].Email)
}

func TestGetUser(t *testing.T) {
	userRepository := userRepository.NewUserRepository()

	userService := user.NewUserService(validator.New(), userRepository)

	gin.SetMode(gin.TestMode)

	app := presentation.SetupServer()

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

	var response responses.User
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	userID := response.ID

	req, _ = http.NewRequest(http.MethodGet, "/users/"+userID, body)
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal(resp.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response.ID)
	assert.Equal(t, "John Doe", response.Name)
	assert.Equal(t, "john.doe@email.com", response.Email)
}
