package controllers

import (
	"net/http"
	"trade-api/database"
	"trade-api/helpers"
	"trade-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VariantRequest struct {
	VariantName string `form:"variant_name" json:"variant_name" binding:"required"`
	Quantity    uint   `form:"quantity" json:"quantity" binding:"required"`
	ProductID   string `form:"product_id" json:"product_id"`
}

func CreateVariant(ctx *gin.Context) {
	db := database.GetDB()


	var Variant models.Variant;
	var variantReq VariantRequest;
	var Product models.Product;

	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&variantReq)
	} else {
		ctx.ShouldBind(&variantReq)
	}

	// Get product by UUID
	if err := db.Where("uuid = ?", variantReq.ProductID).First(&Product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	newUUID := uuid.New()
	Variant.UUID = newUUID.String()
	Variant.VariantName = variantReq.VariantName
	Variant.Quantity = variantReq.Quantity
	Variant.ProductId = Product.ID
	Variant.Product = &Product

	// Create variant
	err := db.Debug().Create(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}

func UpdateVariant(ctx *gin.Context) {
	db := database.GetDB()

	var Variant models.Variant
	var variantReq VariantRequest

	variantUUID := ctx.Param("variantUUID")

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&variantReq)
	} else {
		ctx.ShouldBind(&variantReq)
	}

	// Retrieve existing variant from the database
	if err := db.Model(&Variant).Where("uuid = ?", variantUUID).First(&Variant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Update the Variant struct with retrieved data
	Variant.VariantName = variantReq.VariantName
	Variant.Quantity = variantReq.Quantity

	if err := db.Save(Variant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Respond with the updated variant data
	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}

func GetAllVariants(ctx *gin.Context) {
	db := database.GetDB()

	results := []models.Variant{}

	err := db.Debug().Preload("Product").Find(&results).Error
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


func GetVariantByUUID(ctx *gin.Context) {
	db := database.GetDB()
	variantUUID := ctx.Param("variantUUID")

	var Variant models.Variant

	err := db.Debug().Preload("Product").Where("uuid = ?", variantUUID).First(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}
