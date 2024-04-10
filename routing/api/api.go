package api

import (
	"nfs002/gallery/auth/middleware"
	h "nfs002/gallery/routing/api/handlers"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

func ConfigureApiRoutes(r *gin.RouterGroup) {

	// Wrap the http handler with gin adapter
	r.Use(adapter.Wrap(middleware.EnsureValidToken()))

	r.GET("/host", middleware.ValidateScope("read:host"), h.GetHost)

	r.POST("/host", h.CreateHost)
	r.PUT("/host", h.UpdateHost)
	r.DELETE("/host", h.DeleteHost)

	r.GET("/artist", h.GetArtist)
	r.POST("/artist", h.CreateArtist)
	r.PUT("/artist", h.UpdateArtist)
	r.DELETE("/artist", h.DeleteArtist)

	r.GET("/exhibition", h.GetExhibition)
	r.POST("/exhibition", h.CreateExhibition)
	r.PUT("/exhibition", h.UpdateExhibition)
	r.DELETE("/exhibition", h.DeleteExhibition)

	r.GET("/user", h.GetUser)
	r.POST("/user", h.CreateUser)
	r.PUT("/user", h.UpdateUser)
	r.DELETE("/user", h.DeleteUser)
}
