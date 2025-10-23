package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/caiiomp/vehicle-resale-auth/src/core/useCases/auth"
	"github.com/caiiomp/vehicle-resale-auth/src/core/useCases/user"
	_ "github.com/caiiomp/vehicle-resale-auth/src/docs"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation/authApi"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation/userApi"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/mongodb/userRepository"
)

func main() {
	var (
		mongoURI      = os.Getenv("MONGO_URI")
		mongoDatabase = os.Getenv("MONGO_DATABASE")

		jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

		validate = validator.New()
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("could not initialize mongodb client: %v", err)
	}

	if err = mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Collections
	collection := mongoClient.Database(mongoDatabase).Collection("users")

	// Repositories
	userRepository := userRepository.NewUserRepository(collection)

	// Services
	userService := user.NewUserService(validate, userRepository)
	authService := auth.NewAuthService(userRepository, jwtSecretKey)

	app := presentation.SetupServer()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userApi.RegisterUserRoutes(app, userService)
	authApi.RegisterAuthRoutes(app, authService)

	if err = app.Run(":8080"); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}
