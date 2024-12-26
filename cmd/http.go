package cmd

import (
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/api"
	"e-wallet-ums/internal/repository"
	"e-wallet-ums/internal/services"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeHTTP() {
	registerRepo := repository.RegisterRepository{DB: helpers.DB}
	registerSvc := services.RegisterService{RegisterRepo: &registerRepo}
	registerApi := api.RegisterHandler{RegisterService: &registerSvc}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Failed to set trusted proxies", err)
	}

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerApi.Register)

	err = r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started")
}
