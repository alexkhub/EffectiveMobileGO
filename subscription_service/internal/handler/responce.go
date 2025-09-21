package handler

import (
	subscriptionservice "effective_mobile"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorMessage(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type getListSubscriptionResponse struct {
	Data []subscriptionservice.Subscription `json:"data"`
}