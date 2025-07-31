package controllers

import (
    "net/http"
    "strconv"
    "reffil_finder/models"
    "reffil_finder/config"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// CreatePromotion handles POST /promotions
func CreatePromotion(c *gin.Context) {
    var promotion models.Promotion

    if err := c.ShouldBindJSON(&promotion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    sellerID := uint(c.GetFloat64("user_id")) // ✅ aman dari panic

    // Cek role juga kalau perlu
    role := c.GetString("role")
    if role != "penjual" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can create promotions"})
        return
    }

    promotion.SellerID = sellerID

    if err := config.DB.Create(&promotion).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan promo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":   "Promo berhasil ditambahkan",
        "promotion": promotion,
    })
}


// GetPromotionsBySeller handles GET /promotions
func GetPromotionsBySeller(c *gin.Context) {
    role := c.MustGet("role").(string)
    if role != "penjual" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can view their promotions"})
        return
    }

    sellerID := uint(c.GetFloat64("user_id"))
    var promos []models.Promotion

    if err := config.DB.Where("seller_id = ?", sellerID).Find(&promos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, promos)
}

// UpdatePromotion handles PUT /promotions/:id
func UpdatePromotion(c *gin.Context) {
    role := c.MustGet("role").(string)
    if role != "penjual" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can update promotions"})
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var promo models.Promotion
    if err := config.DB.First(&promo, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    sellerID := uint(c.GetFloat64("user_id"))
    if promo.SellerID != sellerID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to update this promotion"})
        return
    }

    if err := c.ShouldBindJSON(&promo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    promo.SellerID = sellerID // make sure seller_id stays correct

    if err := config.DB.Save(&promo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, promo)
}

// DeletePromotion handles DELETE /promotions/:id
func DeletePromotion(c *gin.Context) {
    role := c.MustGet("role").(string)
    if role != "penjual" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can delete promotions"})
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var promo models.Promotion
    if err := config.DB.First(&promo, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Promotion not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    sellerID := uint(c.GetFloat64("user_id"))
    if promo.SellerID != sellerID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to delete this promotion"})
        return
    }

    if err := config.DB.Delete(&promo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Promotion deleted successfully"})
}
