package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/privatesquare/bkst-go-utils/utils/httputils"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/controllers"
	"github.com/privatesquare/bkst-oauth-api/repository/db"
	"github.com/privatesquare/bkst-oauth-api/services"
	"os"
)

const (
	defaultWebServerPort   = "8080"
	apiServerStartingMsg   = "Starting the API server..."
	apiServerStartedMsg    = "The API server has started and is listening on %s"
	apiServerStartupErrMsg = "Unable to run the web server"

	apiHealthPath     = "/health"
)

func StartApp() {
	r := NewRouter()
	SetupRoutes(r)

	logger.Info(apiServerStartingMsg)
	logger.Info(fmt.Sprintf(apiServerStartedMsg, defaultWebServerPort))
	if err := r.Run(":8080"); err != nil {
		logger.Error(apiServerStartupErrMsg, err)
		os.Exit(1)
	}
}

func NewRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinZap())
	r.Use(gin.Recovery())
	r.NoRoute(httputils.NoRoute)
	r.HandleMethodNotAllowed = true
	r.NoMethod(httputils.MethodNotAllowed)

	return r
}

func SetupRoutes(r *gin.Engine) *gin.Engine {
	r.GET(apiHealthPath, httputils.Health)
	r.GET("/oauth/access_token/:id", controllers.NewAccessTokenHandler(services.New(db.New())).GetById)
	return r
}
