package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/privatesquare/bkst-go-utils/utils/config"
	"github.com/privatesquare/bkst-go-utils/utils/httputils"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/interfaces/db/cassandra"
	"github.com/privatesquare/bkst-oauth-api/interfaces/rest"
	"github.com/privatesquare/bkst-oauth-api/services"
	"os"
)

const (
	defaultWebServerPort    = "8080"
	dbClusterCreationErrMsg = "Unable to create db cluster"
	usingExternalDbMsg      = "Using external %s database listening on %s:%s"
	apiServerStartingMsg    = "Starting the API server..."
	apiServerStartedMsg     = "The API server has started and is listening on %s"
	apiServerStartupErrMsg  = "Unable to run the web server"

	apiHealthPath = "/health"
)

func StartApp() {
	r := httputils.NewRouter()
	setupRoutes(r)
	dbConnect()

	logger.Info(apiServerStartingMsg)
	logger.Info(fmt.Sprintf(apiServerStartedMsg, defaultWebServerPort))
	if err := r.Run(":8080"); err != nil {
		logger.Error(apiServerStartupErrMsg, err)
		os.Exit(1)
	}
}

func dbConnect() {
	cfg := &cassandra.Cfg{}
	if err := config.Load(cfg); err != nil {
		logger.Error(err.Error(), err)
		os.Exit(1)
	}

	logger.Info(fmt.Sprintf(usingExternalDbMsg, cfg.DBDriver, cfg.DBHost, cfg.DBPort))
	if err := cassandra.NewCluster(*cfg); err != nil {
		logger.Error(dbClusterCreationErrMsg, err)
		os.Exit(1)
	}

	session, err := cassandra.NewSession()
	if err != nil {
		logger.Error(err.Error(), err)
		os.Exit(1)
	}
	session.Close()
}

func setupRoutes(r *gin.Engine) *gin.Engine {
	r.GET(apiHealthPath, httputils.Health)
	r.GET("/oauth/access_token/:id", rest.NewAccessTokenHandler(services.NewAccessTokenService(cassandra.NewAccessTokenStore())).GetById)
	return r
}
