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

	contentType := helpers.GetContentType(ctx)
	Variant := models.Variant{}

	variantUUID := ctx.Param("variantUUID")

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Variant)
	} else {
		ctx.ShouldBind(&Variant)
	}

	// Retrieve existing variant from the database
	var getVariant models.Variant
	if err := db.Model(&getVariant).Where("uuid = ?", variantUUID).First(&getVariant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Update the Variant struct with retrieved data
	Variant.ID = uint(getVariant.ID)
	Variant.ProductId = getVariant.ProductId

	// Update the variant record in the database
	updateData := models.Variant{
		VariantName:  Variant.VariantName,
		Quantity: Variant.Quantity,
	}

	if err := db.Model(&Variant).Where("uuid = ?", variantUUID).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Retrieve the updated variant data from the database
	var updatedVariant models.Variant
	if err := db.Where("uuid = ?", variantUUID).Preload("Admin").First(&updatedVariant).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not retrieve updated variant data",
			"message": err.Error(),
		})
		return
	}

	// Respond with the updated variant data
	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedVariant,
	})
}

func GetAllVariants(ctx *gin.Context) {
	db := database.GetDB()

	results := []models.Variant{}

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
