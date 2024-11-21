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
