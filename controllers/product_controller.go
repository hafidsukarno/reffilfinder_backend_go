package controllers

import (
	"reffil_finder/models"
	"reffil_finder/config"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func GetAllProducts(c *gin.Context) {
	role := c.GetString("role")
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan di token"})
		return
	}

	userIDFloat, ok := userIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID format tidak valid"})
		return
	}
	userID := uint(userIDFloat)


	if role != "penjual" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya penjual yang boleh mengakses produk"})
		return
	}

	var products []models.Product
	if err := config.DB.Where("seller_id = ?", userID).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "message": "Produk berhasil dimunculkan",
        // "seller_id": userID,
        "products": products,
    })
}

func CreateProduct(c *gin.Context) {
	role := c.GetString("role")
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan di token"})
		return
	}

	userIDFloat, ok := userIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID format tidak valid"})
		return
	}
	userID := uint(userIDFloat)


	if role != "penjual" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya penjual yang boleh menambah produk"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if product.ProductName == "" || product.ProductPrice < 0 || product.ProductStock < 0 {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid: nama tidak boleh kosong, harga/stok tidak boleh negatif"})
	return
	}

	product.SellerID = uint(userID)

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "message": "Produk berhasil ditambahkan",
        "product": gin.H{
            "id":                 product.ID,
            "product_name":       product.ProductName,
            "product_description": product.ProductDescription,
            "product_price":      product.ProductPrice,
            "product_stock":      product.ProductStock,
            "product_image":      product.ProductImage,
            "seller_id":          product.SellerID,
        },
    })
}

func UpdateProduct(c *gin.Context) {
		role := c.GetString("role")
		userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan di token"})
		return
	}

	userIDFloat, ok := userIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID format tidak valid"})
		return
	}
	userID := uint(userIDFloat)


	if role != "penjual" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya penjual yang boleh mengubah produk"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	if product.SellerID != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tidak bisa mengubah produk orang lain"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui produk"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	role := c.GetString("role")
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan di token"})
		return
	}

	userIDFloat, ok := userIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID format tidak valid"})
		return
	}
	userID := uint(userIDFloat)


	if role != "penjual" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya penjual yang boleh menghapus produk"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	if product.SellerID != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tidak bisa menghapus produk orang lain"})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}
