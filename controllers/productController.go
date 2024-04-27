package controllers

import (
	"mime/multipart"
	"net/http"
	"trade-api/database"
	"trade-api/helpers"
	"trade-api/models"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ProductRequest struct {
	Name  	string                   	`form:"name" binding:"required"`
	Image  	*multipart.FileHeader 		`form:"file"`
}

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	var Product models.Product;
	var productReq ProductRequest;

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)

	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the filename without extension
	fileName := helpers.RemoveExtension(productReq.Image.Filename)

	uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Product.Name = productReq.Name;
	Product.ImageUrl = uploadResult;
	Product.AdminId =  uint(adminData["id"].(float64))
	Product.UUID = uuid.New().String()

	err = db.Debug().Create(&Product).Error
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

	var Product models.Product
	var productReq ProductRequest

	productUUID := ctx.Param("productUUID")

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&productReq); err != nil {
			ctx.JSON(http.StatusNotAcceptable, gin.H{
				"error":   "Bad request",
				"message": err.Error()})
			return
		}
	} else {
		if err := ctx.ShouldBind(&productReq); err != nil {
			ctx.JSON(http.StatusNotAcceptable, gin.H{
				"error":   "Bad request",
				"message": err.Error()})
			return
		}
	}

	// Retrieve existing product from the database
	if err := db.Model(&Product).Where("uuid = ?", productUUID).First(&Product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// If there is a new image, upload to cloudinary and update product image url
	if productReq.Image != nil {
		// Extract the filename without extension
		fileName := helpers.RemoveExtension(productReq.Image.Filename)

		uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Product.ImageUrl = uploadResult
	}

	Product.Name = productReq.Name

	// Update DB with new product data
	if err := db.Save(&Product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Respond with the updated product data
	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func GetAllProducts(ctx *gin.Context) {
	db := database.GetDB()

	results := []models.Product{}

	search := ctx.Query("search") + "%"

	err := db.Debug().Where("name LIKE ?", search).Find(&results).Error
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

func GetProductByUUID(ctx *gin.Context) {
	db := database.GetDB()
	productUUID := ctx.Param("productUUID")

	var Product models.Product

	err := db.Debug().Where("uuid = ?", productUUID).First(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	productUUID := ctx.Param("productUUID")

	var Product models.Product
	var Variant models.Variant

	// Get product
	if err := db.Debug().Where("uuid = ?", productUUID).First(&Product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	// Remove variant
	if err := db.Debug().Where("product_id = ?", Product.ID).Delete(&Variant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Remove product
	if 	err := db.Debug().Where("uuid = ?", productUUID).Delete(&Product).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}