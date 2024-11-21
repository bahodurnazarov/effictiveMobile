package handler

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateSong creates a new song
// @Summary Create a new song
// @Schemes
// @Description Adds a new song to the database using the provided JSON payload.
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Song details"
// @Success 201 {object} models.Song "The created song object"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/song [post]
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

// DeleteSong deletes a song by ID
// @Summary Delete a song
// @Schemes
// @Description Deletes a song from the database using its unique identifier.
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/song/{id} [delete]
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

// GetSongByID godoc
// @Summary Get a song by ID
// @Description Retrieve a song using its unique ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 200 {object} models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/song/{id} [get]
func (h *Handler) GetSongByID(c *gin.Context) {
	songID := c.Param("id")
	log.Println("id", songID)
	if songID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "song ID is required"})
		return
	}

	song, err := h.Service.GetSongByID(songID)
	if err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	logger.InfoLogger.Println(song)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": song})
}
