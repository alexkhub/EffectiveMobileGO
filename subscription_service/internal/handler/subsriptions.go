package handler

import (
	"context"
	subscriptionservice "effective_mobile"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CreateSubscriptionHandler godoc
// @Summary Создать подписку
// @Description Создаёт новую подписку пользователю
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body subscriptionservice.CreateSubscription true "Subscription"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscription [post]
func (h *Handler) CreateSubscriptionHandler(c *gin.Context) {
	var request subscriptionservice.CreateSubscription

	reqId, ok := c.Get("req_id")
	if !ok {
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "CreateSubscriptionHandler",
	}).Debug()

	if err := c.BindJSON(&request); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		c.Set("message", err.Error())
		return
	}
	serviceCtx := context.WithValue(c, "req_id", reqId)

	response, err := h.service.CreateSubscriptionService(serviceCtx, request)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}
	c.Set("message", fmt.Sprintf("create subscription - %s", response))

	c.JSON(http.StatusOK, gin.H{
		"id": response,
	})
}

// GetSubscriptionHandler godoc
// @Summary Получить подписку
// @Description Возвращает подписку по ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} subscriptionservice.Subscription
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscription/{id} [get]
func (h *Handler) GetSubscriptionHandler(c *gin.Context) {
	reqId, ok := c.Get("req_id")
	if !ok {
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "GetSubscriptionHandler",
	}).Debug()

	subId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}
	serviceCtx := context.WithValue(c, "req_id", reqId)
	response, err := h.service.GetSubscriptionService(serviceCtx, subId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListSubscriptionHandler godoc
// @Summary Список подписок
// @Description Возвращает все подписки
// @Tags subscriptions
// @Produce json
// @Success 200 {object} []subscriptionservice.Subscription
// @Failure 500 {object} map[string]string
// @Router /subscription [get]
func (h *Handler) ListSubscriptionHandler(c *gin.Context) {
	reqId, ok := c.Get("req_id")
	if !ok {
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "ListSubscriptionHandler",
	}).Debug()

	serviceCtx := context.WithValue(c, "req_id", reqId)

	response, err := h.service.ListSubscriptionService(serviceCtx)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}
	c.JSON(http.StatusOK, getListSubscriptionResponse{
		Data: response,
	})

}

// DeleteSubscriptionHandler godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Param id path string true "Subscription ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscription/{id} [delete]
func (h *Handler) DeleteSubscriptionHandler(c *gin.Context) {
	reqId, ok := c.Get("req_id")
	if !ok {
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "DeleteSubscriptionHandler",
	}).Debug()

	subId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}
	serviceCtx := context.WithValue(c, "req_id", reqId)
	err = h.service.DeleteSubscriptionService(serviceCtx, subId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}
	c.Set("message", fmt.Sprintf("subsription delete - %s", subId))
	c.JSON(http.StatusNoContent, nil)
}

// UpdateSubscriptionHandler godoc
// @Summary Обновить подписку
// @Description Обновляет существующую подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param subscription body subscriptionservice.UpdateSubscription true "Данные для обновления"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscription/{id} [patch]
func (h *Handler) UpdateSubscriptionHandler(c *gin.Context) {
	var request subscriptionservice.UpdateSubscription

	reqId, ok := c.Get("req_id")
	if !ok {
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "UpdateSubscriptionHandler",
	}).Debug()

	if err := c.BindJSON(&request); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		c.Set("message", err.Error())
		return
	}

	subId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}

	request.Id = subId
	serviceCtx := context.WithValue(c, "req_id", reqId)

	err = h.service.UpdateSubscriptionService(serviceCtx, request)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}
	c.Set("message", fmt.Sprintf("update subscription - %s", subId))

	c.JSON(http.StatusOK, gin.H{
		"id": subId,
	})
}

// TotalPriceHandler godoc
// @Summary Общая сумма подписок
// @Description Считает сумму всех подписок за период с фильтрами
// @Tags subscriptions
// @Produce json
// @Param date_from query string false "Начало периода (MM-YYYY)"
// @Param date_to query string false "Конец периода (MM-YYYY)"
// @Param user_id query string false "UUID пользователя"
// @Param service_name query string false "Название сервиса"
// @Success 200 {object} map[string]int
// @Failure 500 {object} map[string]string
// @Router /subscription/total [get]
func (h *Handler) TotalPriceHandler(c *gin.Context) {
	filters := subscriptionservice.FilterSubscription{}
	reqId, ok := c.Get("req_id")
	if !ok {
		reqId = "none"
	}
	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "TotalPriceHandler",
	}).Debug()

	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")
	userId := c.Query("user_id")
	serviceName := c.Query("service_name")

	if dateFrom == "" {
		filters.StartDateFrom.SetValid("01-1970")
	} else {
		filters.StartDateFrom.SetValid(dateFrom)
	}

	if dateTo == "" {
		now := time.Now().Format("01-2006")
		filters.StartDateTo.SetValid(now)
	} else {
		filters.StartDateTo.SetValid(dateTo)
	}

	if userId != "" {
		filters.UserId.SetValid(userId)
	}
	if serviceName != "" {
		filters.ServiceName.SetValid(serviceName)
	}

	logrus.WithFields(logrus.Fields{
		"req_id": reqId,
		"method": "TotalPriceHandler",
		"data":   filters,
	}).Debug()
	serviceCtx := context.WithValue(c, "req_id", reqId)

	total, err := h.service.TotalPriceService(serviceCtx, filters)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		c.Set("message", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})

}
