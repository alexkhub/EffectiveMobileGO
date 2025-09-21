package handler

import (
	"effective_mobile/internal/middleware"
	"effective_mobile/internal/service"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "effective_mobile/docs"
)


type Handler struct{
	service *service.Sevice
}

func NewHandler(service *service.Sevice) *Handler{
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	
	router := gin.New()
	router.Use(middleware.Logging())
	router.MaxMultipartMemory = 15 << 20
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		subscription := api.Group("/subscription")
		{
			subscription.POST("", h.CreateSubscriptionHandler)
			subscription.GET("/:id", h.GetSubscriptionHandler)
			subscription.GET("", h.ListSubscriptionHandler)
			subscription.DELETE("/:id", h.DeleteSubscriptionHandler)
			subscription.PATCH("/:id", h.UpdateSubscriptionHandler)
			subscription.GET("/total", h.TotalPriceHandler)
		}
	}
	
	return router

}