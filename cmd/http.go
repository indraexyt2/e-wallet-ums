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
	dependency := dependencyInject()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Failed to set trusted proxies", err)
	}

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)

	err = r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started")
}

type Dependency struct {
	RegisterAPI *api.RegisterHandler
	LoginAPI    *api.LoginHandler
}

func dependencyInject() Dependency {
	userRepo := &repository.UserRepository{DB: helpers.DB}

	registerSvc := &services.RegisterService{RegisterRepo: userRepo}
	registerApi := &api.RegisterHandler{RegisterService: registerSvc}

	loginSvc := &services.LoginService{UserRepo: userRepo}
	loginApi := &api.LoginHandler{LoginService: loginSvc}

	return Dependency{
		RegisterAPI: registerApi,
		LoginAPI:    loginApi,
	}
}
