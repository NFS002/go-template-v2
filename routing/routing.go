package routing

import (
	api "nfs002/template/v2/routing/api"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Serve frontend static files
	router.Use(static.Serve("/static", static.LocalFile("./static/", true)))

	// Configure API routes
	apiGroup := router.Group("/api")
	api.ConfigureApiRoutes(apiGroup)

	return router
}
