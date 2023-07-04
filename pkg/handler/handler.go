package handler

import (
	"github.com/2hard4me/simple-rest-api/pkg/service"
	"github.com/gin-gonic/gin"
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
	{
		api.POST("/create_user", h.CreateUser)
		api.GET("/users", h.GetUsers)
		api.GET("/get_user/:id", h.GetUserByID)
		api.PUT("/update_user/:id", h.UpdateUser)
		api.DELETE("/delete_user/:id", h.DeleteUser)
	}

	return router
}