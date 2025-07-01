package handler

import (
	"encoding/json"
	"net/http"

	"go-ecommerce/internal/config"
	"go-ecommerce/internal/model"
	"go-ecommerce/internal/service"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderHandler struct {
	svc *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

// CreateOrder godoc
// @Summary Create order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body model.Order true "Order data"
// @Success 201 {object} model.Order
// @Router /orders [post]
func (h *OrderHandler) Create(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID := c.GetUint("user_id")
	order.UserID = userID

	if err := h.svc.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	// Gửi message tới RabbitMQ
	orderJSON, _ := json.Marshal(order)
	_ = config.MQChannel.Publish(
		"order_exchange",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        orderJSON,
		},
	)

	c.JSON(http.StatusCreated, order)
}

// GetUserOrders godoc
// @Summary Get current user's orders
// @Tags orders
// @Produce json
// @Success 200 {array} model.Order
// @Router /orders [get]
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID := c.GetUint("user_id")
	orders, err := h.svc.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetByID godoc
// @Summary Get order by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} model.Order
// @Router /orders/{id} [get]
func (h *OrderHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	order, err := h.svc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// GetAll godoc
// @Summary Get all orders (admin only)
// @Tags admin-orders
// @Produce json
// @Success 200 {array} model.Order
// @Router /admin/orders [get]
func (h *OrderHandler) GetAll(c *gin.Context) {
	orders, err := h.svc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// UpdateStatus godoc
// @Summary Update order status (admin only)
// @Tags admin-orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body map[string]string true "New status"
// @Success 200 {object} string
// @Router /admin/orders/{id} [put]
func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var payload struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil || payload.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	if err := h.svc.UpdateStatus(id, payload.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "status updated"})
}

// Delete godoc
// @Summary Delete an order (admin only)
// @Tags admin-orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} string
// @Router /admin/orders/{id} [delete]
func (h *OrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
