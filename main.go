package main

import (
	"net/http"

	"github.com/eriicafes/filedb"
	"github.com/eriicafes/go-api-starter/controller"
	"github.com/eriicafes/go-api-starter/middlewares"
	"github.com/eriicafes/go-api-starter/modules/auth"
	"github.com/eriicafes/go-api-starter/modules/buckets"
	"github.com/gin-gonic/gin"
)

var (
	database          = filedb.New("store")
	authRepository    = auth.NewAuthRepository(database)
	authService       = auth.NewAuthService(authRepository)
	authController    = auth.NewAuthController(authService)
	bucketsRepository = buckets.NewBucketsRepository(database)
	bucketsService    = buckets.NewBucketsService(bucketsRepository)
	bucketsController = buckets.NewBucketsController(bucketsService)
)

func main() {
	router := gin.Default()

	router.Use(middlewares.UseCors())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	router.POST("/auth/signin", authController.SignIn)
	router.POST("/auth/signout", middlewares.UseAuth(authService), authController.SignOut)
	router.GET("/auth/profile", middlewares.UseAuth(authService), authController.Profile)

	binder := controller.NewBinder(router)

	binder.Bind("/buckets", bucketsController, middlewares.UseAuth(authService))

	router.Run()
}
