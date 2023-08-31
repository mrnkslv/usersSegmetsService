package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrnkslv/user-segmentation-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")

	slugs := api.Group("/slugs")
	{
		slugs.POST("/", h.createSlug)
		slugs.DELETE("/", h.deleteSlug)
	}
	users := api.Group("/users")
	{
		users.POST("/", h.addUserToSlug)     //slugid
		users.GET("/", h.getActiveSlugsByID) //userid
	}
	return router
}
