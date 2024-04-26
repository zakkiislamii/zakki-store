package controllers

import (
	"log"
	"net/http"
	"strconv"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexPabrik(c *gin.Context) {
	db := models.DB

	Pabrik, err := models.GetAllPabrik(db)
	if err != nil {
		log.Printf("Failed to get Pabrik data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Pabrik data"})
		return
	}

	c.JSON(http.StatusOK, Pabrik)
}

func CreatePabrik(c *gin.Context) {
	db := models.DB

	var Pabrik structs.Pabrik
	if err := c.ShouldBindJSON(&Pabrik); err != nil {
		log.Printf("Invalid JSON format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := models.InsertPabrik(db, Pabrik); err != nil {
		log.Printf("Failed to insert Pabrik: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert Pabrik"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Pabrik"})
}

func UpdatePabrik(c *gin.Context) {
	db := models.DB

	var Pabrik structs.Pabrik
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.ShouldBindJSON(&Pabrik); err != nil {
		log.Printf("Invalid JSON format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	Pabrik.IdPabrik = int(id)
	if err := models.UpdatePabrik(db, Pabrik); err != nil {
		log.Printf("Failed to update Pabrik: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Pabrik"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Update Pabrik"})
}

func DeletePabrik(c *gin.Context) {
	db := models.DB

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := models.DeletePabrik(db, id); err != nil {
		log.Printf("Failed to delete Pabrik: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Pabrik"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Delete Pabrik"})
}
