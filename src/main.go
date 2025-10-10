package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/caiiomp/vehicle-resale-auth/src/core/domain/useCases/user"
	"github.com/caiiomp/vehicle-resale-auth/src/presentation/userApi"
	"github.com/caiiomp/vehicle-resale-auth/src/repository/userRepository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		mongoURI      = os.Getenv("MONGO_URI")
		mongoDatabase = os.Getenv("MONGO_DATABASE")
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

	collection := mongoClient.Database(mongoDatabase).Collection("vehicles")

	userRepository := userRepository.NewUserRepository(collection)
	userService := user.NewUserRepository(userRepository)

	app := gin.Default()

	userApi.RegisterUserRoutes(app, userService)

	if err = app.Run(":4000"); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}
