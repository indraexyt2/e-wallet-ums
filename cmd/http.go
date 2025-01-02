package cmd

import (
	"e-wallet-ums/external"
	"e-wallet-ums/helpers"
	"e-wallet-ums/internal/api"
	"e-wallet-ums/internal/interfaces"
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

	userV1WithAuth := userV1.Use()
	userV1WithAuth.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.Logout)
	userV1WithAuth.PUT("/refresh-token", dependency.MiddleWareRefreshToken, dependency.RefreshTokenAPI.RefreshToken)

	err = r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started")
}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	RegisterAPI     interfaces.IRegisterHandler
	LoginAPI        interfaces.ILoginHandler
	LogoutAPI       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler

	TokenValidationAPI *api.TokenValidationHandler
}

func dependencyInject() Dependency {
	userRepo := &repository.UserRepository{DB: helpers.DB}
	ext := &external.External{}

	registerSvc := &services.RegisterService{
		RegisterRepo: userRepo,
		External:     ext,
	}
	registerApi := &api.RegisterHandler{RegisterService: registerSvc}

	loginSvc := &services.LoginService{UserRepo: userRepo}
	loginApi := &api.LoginHandler{LoginService: loginSvc}

	logoutSvc := &services.LogoutService{UserRepo: userRepo}
	logoutApi := &api.LogoutHandler{LogoutService: logoutSvc}

	refreshTokenSvc := &services.RefreshTokenService{UserRepo: userRepo}
	refreshTokenApi := &api.RefreshTokenHandler{RefreshTokenService: refreshTokenSvc}

	tokenValidationSvc := &services.TokenValidationService{UserRepo: userRepo}
	tokenValidationApi := &api.TokenValidationHandler{TokenValidationService: tokenValidationSvc}

	return Dependency{
		UserRepository:  userRepo,
		RegisterAPI:     registerApi,
		LoginAPI:        loginApi,
		LogoutAPI:       logoutApi,
		RefreshTokenAPI: refreshTokenApi,

		TokenValidationAPI: tokenValidationApi,
	}
}
