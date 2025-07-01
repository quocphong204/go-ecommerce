package handler

import (
	"encoding/json"
	"net/http"

	"go-ecommerce/internal/config"
	"go-ecommerce/internal/model"
	"go-ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	svc service.ProductService
}

func NewProductHandler(svc service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

// GetAll godoc
// @Summary Get all products
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Router /products [get]
func (h *ProductHandler) GetAll(c *gin.Context) {
	const cacheKey = "products:all"

	// Try get from Redis cache
	if cached, err := config.RedisClient.Get(config.RedisCtx, cacheKey).Result(); err == nil {
		c.Data(http.StatusOK, "application/json", []byte(cached))
		return
	}

	// Cache miss → get from DB
	products, err := h.svc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}

	// Cache result
	data, _ := json.Marshal(products)
	_ = config.RedisClient.Set(config.RedisCtx, cacheKey, data, 0).Err()

	c.JSON(http.StatusOK, products)
}

// Create godoc
// @Summary Create new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product data"
// @Success 201 {object} model.Product
// @Router /products [post]
func (h *ProductHandler) Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.svc.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	// Invalidate cache
	_ = config.RedisClient.Del(config.RedisCtx, "products:all").Err()

	c.JSON(http.StatusCreated, product)
}

// Update godoc
// @Summary Update a product
// @Tags admin-products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Updated data"
// @Success 200 {object} model.Product
// @Router /admin/products/{id} [put]
func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.svc.Update(id, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	// Invalidate cache
	_ = config.RedisClient.Del(config.RedisCtx, "products:all").Err()

	c.JSON(http.StatusOK, product)
}

// Delete godoc
// @Summary Delete a product
// @Tags admin-products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {string} string "product deleted"
// @Router /admin/products/{id} [delete]
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	// Invalidate cache
	_ = config.RedisClient.Del(config.RedisCtx, "products:all").Err()

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
