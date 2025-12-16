package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"seriusbrand/backend/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// CreateOrder creates a new order
func (h *Handler) CreateOrder(c *gin.Context) {
	var req struct {
		CustomerName string `json:"customer_name" binding:"required"`
		Whatsapp     string `json:"whatsapp" binding:"required"`
		ProductName  string `json:"product_name" binding:"required"`
		PackageType  string `json:"package_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate price based on package
	price := 10000
	switch req.PackageType {
	case "starter":
		price = 10000
	case "growth":
		price = 25000
	case "pro":
		price = 50000
	}

	order := models.Order{
		CustomerName: req.CustomerName,
		Whatsapp:     req.Whatsapp,
		ProductName:  req.ProductName,
		PackageType:  req.PackageType,
		Price:        price,
		Status:       "pending",
	}

	if err := h.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// UploadPaymentProof uploads payment proof for an order
func (h *Handler) UploadPaymentProof(c *gin.Context) {
	orderID := c.Param("id")

	file, err := c.FormFile("payment_proof")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
	uploadDir := "./uploads/proofs"
	savePath := fmt.Sprintf("%s/%s", uploadDir, filename)

	// Ensure directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Save file
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Update order
	fileURL := fmt.Sprintf("/uploads/proofs/%s", filename)
	if err := h.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("payment_proof_url", fileURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment proof uploaded", "file_url": fileURL})
}

// ListOrders lists all orders (admin)
func (h *Handler) ListOrders(c *gin.Context) {
	var orders []models.Order
	if err := h.DB.Order("created_at DESC").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// UpdateOrderStatus updates order status (admin)
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated"})
}

// GetUmkmPage gets UMKM page by slug
func (h *Handler) GetUmkmPage(c *gin.Context) {
	slug := c.Param("slug")

	var page models.UmkmPage
	if err := h.DB.Where("slug = ?", slug).First(&page).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, page)
}

// CreateUmkmPage creates a new UMKM page (admin)
func (h *Handler) CreateUmkmPage(c *gin.Context) {
	var req struct {
		OrderID         uint   `form:"order_id" binding:"required"`
		Slug            string `form:"slug" binding:"required"`
		Photos          string `form:"photos"`
		Description     string `form:"description"`
		Price           string `form:"price"`
		WhatsappLink    string `form:"whatsapp_link"`
		MarketplaceLink string `form:"marketplace_link"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	videoURL := ""
	file, err := c.FormFile("video_file")
	if err == nil {
		// Generate unique filename
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
		uploadDir := "./uploads/videos"
		savePath := fmt.Sprintf("%s/%s", uploadDir, filename)

		// Ensure directory exists
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		// Save file
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video file"})
			return
		}
		videoURL = fmt.Sprintf("/uploads/videos/%s", filename)
	}

	page := models.UmkmPage{
		OrderID:         req.OrderID,
		Slug:            req.Slug,
		VideoURL:        videoURL,
		Photos:          req.Photos,
		Description:     req.Description,
		Price:           req.Price,
		WhatsappLink:    req.WhatsappLink,
		MarketplaceLink: req.MarketplaceLink,
	}

	if err := h.DB.Create(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page"})
		return
	}

	c.JSON(http.StatusCreated, page)
}

// UpdateUmkmPage updates an existing UMKM page (admin)
func (h *Handler) UpdateUmkmPage(c *gin.Context) {
	pageID := c.Param("id")

	var req struct {
		Slug            string `form:"slug"`
		Photos          string `form:"photos"`
		Description     string `form:"description"`
		Price           string `form:"price"`
		WhatsappLink    string `form:"whatsapp_link"`
		MarketplaceLink string `form:"marketplace_link"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find existing page
	var page models.UmkmPage
	if err := h.DB.First(&page, pageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	// Handle video upload if provided
	file, err := c.FormFile("video_file")
	if err == nil {
		// Generate unique filename
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
		uploadDir := "./uploads/videos"
		savePath := fmt.Sprintf("%s/%s", uploadDir, filename)

		// Ensure directory exists
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		// Save new file
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video file"})
			return
		}

		// Update video URL
		page.VideoURL = fmt.Sprintf("/uploads/videos/%s", filename)
	}

	// Update other fields
	page.Slug = req.Slug
	page.Description = req.Description
	page.Price = req.Price
	page.WhatsappLink = req.WhatsappLink
	page.MarketplaceLink = req.MarketplaceLink

	// Only update photos if provided and valid JSON
	if req.Photos != "" && req.Photos != "[]" {
		page.Photos = req.Photos
	}

	if err := h.DB.Save(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update page"})
		return
	}

	c.JSON(http.StatusOK, page)
}

// ListUmkmPages lists all UMKM pages (public)
func (h *Handler) ListUmkmPages(c *gin.Context) {
	var pages []models.UmkmPage
	// Preload the Order to get the product name if needed, but for now just fetching pages
	if err := h.DB.Order("created_at DESC").Find(&pages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pages"})
		return
	}

	c.JSON(http.StatusOK, pages)
}
