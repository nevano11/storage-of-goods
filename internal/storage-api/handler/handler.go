package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nevano11/storage-of-goods/internal/storage-api/service"
)

type HttpHandler struct {
	service *service.Service
}

func NewHttpHandler(service *service.Service) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		storage := api.Group("/storage")
		{
			storage.POST("/reserve", h.reserveProducts)
			storage.POST("/free", h.freeReserved)

			storage.GET("/goods-in-stock", h.getGoodsInStockByCode)
		}
	}
	return router
}
