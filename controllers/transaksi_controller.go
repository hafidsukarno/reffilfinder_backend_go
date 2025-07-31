package controllers

import (
	"reffil_finder/models"
	"reffil_finder/config"
	"time"
	"gorm.io/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST /transaksi
func CreateTransaksi(c *gin.Context) {
	var input struct {
		PaymentMethod string `json:"payment_method" binding:"required"`
		Address       string `json:"address" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uint(c.GetFloat64("user_id"))

	transaksi := models.Transaksi{
		UserID:        userID,
		PaymentMethod: input.PaymentMethod,
		Address:       input.Address,
		TransaksiDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	config.DB.Create(&transaksi)

	var carts []models.Cart
	config.DB.Where("user_id = ?", userID).Find(&carts)

	for _, cart := range carts {
		detail := models.DetailTransaksi{
			TransaksiID: transaksi.ID,
			ProductID:   cart.ProductID,
			Quantity:    cart.Quantity,
			Status:      "pending",
		}
		config.DB.Create(&detail)
		config.DB.Model(&models.Product{}).Where("id = ?", cart.ProductID).
			UpdateColumn("product_stock", gorm.Expr("product_stock - ?", cart.Quantity))
	}

	config.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	c.JSON(http.StatusOK, gin.H{"message": "Checkout berhasil", "transaksi": transaksi})
}

// GET semua transaksi user
func GetTransaksiUser(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	var transaksi []models.Transaksi
	config.DB.Preload("DetailTransaksis").Where("user_id = ?", userID).Find(&transaksi)
	c.JSON(http.StatusOK, transaksi)
}

// PUT status detail transaksi (ubah status jadi berhasil/batal)
func UpdateStatusDetail(c *gin.Context) {
	id := c.Param("id")
	var detail models.DetailTransaksi
	if err := config.DB.First(&detail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Detail transaksi tidak ditemukan"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail.Status = input.Status
	config.DB.Save(&detail)

	c.JSON(http.StatusOK, detail)
}

// DELETE transaksi (dan semua detail-nya)
func DeleteTransaksi(c *gin.Context) {
	id := c.Param("id")

	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	config.DB.Where("transaksi_id = ?", transaksi.ID).Delete(&models.DetailTransaksi{})
	config.DB.Delete(&transaksi)

	c.JSON(http.StatusOK, gin.H{"message": "Transaksi dan detail berhasil dihapus"})
}
