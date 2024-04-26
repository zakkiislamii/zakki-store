package controllers

import (
	"log"
	"net/http"
	"strconv"
	"zakki-store/models"
	"zakki-store/structs"

	"github.com/gin-gonic/gin"
)

func IndexTokoBaju(c *gin.Context) {
	db := models.DB

	TokoBaju, err := models.GetAllTokoBaju(db)
	if err != nil {
		log.Printf("Failed to get TokoBaju data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get TokoBaju data"})
		return
	}

	c.JSON(http.StatusOK, TokoBaju)
}

func CreateTokoBaju(c *gin.Context) {
	db := models.DB

	var TokoBaju structs.TokoBaju
	if err := c.ShouldBindJSON(&TokoBaju); err != nil {
		log.Printf("Invalid JSON format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := models.InsertTokoBaju(db, TokoBaju); err != nil {
		log.Printf("Failed to insert TokoBaju: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert TokoBaju"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert TokoBaju"})
}

func UpdateTokoBaju(c *gin.Context) {
	db := models.DB

	var TokoBaju structs.TokoBaju
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.ShouldBindJSON(&TokoBaju); err != nil {
		log.Printf("Invalid JSON format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	TokoBaju.IdTokoBaju = int(id)
	if err := models.UpdateTokoBaju(db, TokoBaju); err != nil {
		log.Printf("Failed to update TokoBaju: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update TokoBaju"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Update TokoBaju"})
}

func DeleteTokoBaju(c *gin.Context) {
	db := models.DB

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := models.DeleteTokoBaju(db, id); err != nil {
		log.Printf("Failed to delete TokoBaju: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete TokoBaju"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Delete TokoBaju"})
}
