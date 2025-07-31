package controllers

import (
	"net/http"
	"reffil_finder/config"
	"reffil_finder/models"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))

	var carts []models.Cart
	if err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"carts": carts})
}

func AddToCart(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah produk ada
	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	// Cek stok
	if product.ProductStock < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stok produk tidak mencukupi"})
		return
	}

	// Cek apakah produk sudah ada di cart
	var existingCart models.Cart
	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&existingCart).Error; err == nil {
		// Update quantity
		existingCart.Quantity += input.Quantity
		if err := config.DB.Save(&existingCart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan produk ke cart"})
			return
		}
	} else {
		// Tambah item baru
		newCart := models.Cart{
			UserID:    userID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		if err := config.DB.Create(&newCart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan produk ke cart"})
			return
		}
	}

	// Kurangi stok produk
	product.ProductStock -= input.Quantity
	config.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil ditambahkan ke cart"})
}

func UpdateCartItem(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	cartID := c.Param("id")

	var input struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	if err := config.DB.Preload("Product").First(&cart, cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item cart tidak ditemukan"})
		return
	}

	if cart.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tidak diizinkan"})
		return
	}

	selisih := input.Quantity - cart.Quantity

	// Cek apakah stok cukup (jika menambah)
	if selisih > 0 && cart.Product.ProductStock < selisih {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stok produk tidak mencukupi"})
		return
	}

	// Update stok produk
	cart.Product.ProductStock -= selisih
	config.DB.Save(&cart.Product)

	cart.Quantity = input.Quantity
	if err := config.DB.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengubah jumlah produk di cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jumlah produk di cart berhasil diperbarui"})
}

func DeleteCartItem(c *gin.Context) {
	userID := uint(c.GetFloat64("user_id"))
	cartID := c.Param("id")

	var cart models.Cart
	if err := config.DB.Preload("Product").First(&cart, cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item cart tidak ditemukan"})
		return
	}

	if cart.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tidak diizinkan"})
		return
	}

	// Kembalikan stok produk
	cart.Product.ProductStock += cart.Quantity
	config.DB.Save(&cart.Product)

	if err := config.DB.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus item dari cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item berhasil dihapus dari cart"})
}
