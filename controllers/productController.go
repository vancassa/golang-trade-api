package controllers

import (
	"net/http"
	"trade-api/database"
	"trade-api/helpers"
	"trade-api/models"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Product := models.Product{}
	adminID := uint(adminData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.AdminId = adminID
	newUUID := uuid.New()
	Product.UUID = newUUID.String()

	err := db.Debug().Create(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Product := models.Product{}

	productUUID := ctx.Param("productUUID")
	adminID := uint(adminData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	// Retrieve existing product from the database
	var getProduct models.Product
	if err := db.Model(&getProduct).Where("uuid = ?", productUUID).First(&getProduct).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Update the Product struct with retrieved data
	Product.ID = uint(getProduct.ID)
	Product.AdminId = adminID

	// Update the product record in the database
	updateData := models.Product{
		Name:  Product.Name,
		File: Product.File,
	}

	if err := db.Model(&Product).Where("uuid = ?", productUUID).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Retrieve the updated product data from the database
	var updatedProduct models.Product
	if err := db.Where("uuid = ?", productUUID).Preload("Admin").First(&updatedProduct).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not retrieve updated product data",
			"message": err.Error(),
		})
		return
	}

	// Respond with the updated product data
	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedProduct,
	})
}

func GetAllProducts(ctx *gin.Context) {
	db := database.GetDB()

	results := []models.Product{}

	err := db.Debug().Preload("Admin").Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}
