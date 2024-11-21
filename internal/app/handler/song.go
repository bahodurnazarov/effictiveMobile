package handler

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateSong(c *gin.Context) {
	var requestBody models.Song
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	logger.InfoLogger.Println(requestBody)

	newSong, err := h.Service.CreateSong(requestBody)
	if err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	logger.InfoLogger.Println("New Song Created Successfully")
	c.JSON(http.StatusCreated, newSong)
}

func (h *Handler) DeleteSong(c *gin.Context) {
	songID := c.Param("id")
	logger.InfoLogger.Println("id", songID)
	if songID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "song ID is required"})
		return
	}

	err := h.Service.DeleteSong(songID)
	if err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	logger.InfoLogger.Println("song deleted successfully")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "song deleted successfully"})
}
