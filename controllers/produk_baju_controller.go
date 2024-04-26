package controllers

import (
	"log"
	"net/http"
	"strconv"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexProdukBaju(c *gin.Context) {
	db := models.DB

	ProdukBaju, err := models.GetAllProdukBaju(db)
	if err != nil {
		log.Printf("Failed to get Produk Baju data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ProdukBaju data"})
		return
	}

	c.JSON(http.StatusOK, ProdukBaju)
}

func CreateProdukBaju(c *gin.Context) {
	db := models.DB

	var ProdukBaju structs.ProdukBaju
	if err := c.ShouldBindJSON(&ProdukBaju); err != nil {
		log.Printf("Invalid JSON format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := models.InsertProdukBaju(db, ProdukBaju); err != nil {
		log.Printf("Failed to insert ProdukBaju: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert ProdukBaju"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert ProdukBaju"})
}

func UpdateProdukBaju(c *gin.Context) {
	db := models.DB

	var ProdukBaju structs.ProdukBaju
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.ShouldBindJSON(&ProdukBaju); err != nil {
		log.Printf("Invalid JSON format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	ProdukBaju.IdProduk = int(id)
	if err := models.UpdateProdukBaju(db, ProdukBaju); err != nil {
		log.Printf("Failed to update ProdukBaju: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ProdukBaju"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Update ProdukBaju"})
}

func DeleteProdukBaju(c *gin.Context) {
	db := models.DB

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := models.DeleteProdukBaju(db, id); err != nil {
		log.Printf("Failed to delete ProdukBaju: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ProdukBaju"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Delete ProdukBaju"})
}
