package handler

import (
	"github.com/bahodurnazarov/effictiveMobile/pkg/models"
	logger "github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// CreateSong creates a new song
// @Summary Create a new song
// @Schemes
// @Description Adds a new song to the database using the provided JSON payload.
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body models.RequestSong true "Song details"
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
// @Success 200 {object} models.RequestSong
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

// UpdateSong godoc
// @Summary Update an existing song
// @Description Updates the details of a song by its ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param song body models.RequestSong true "Updated song details"
// @Success 200 {object} models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/song/{id} [put]
func (h *Handler) UpdateSong(c *gin.Context) {
	songID := c.Param("id")
	log.Println("id", songID)
	if songID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "song ID is required"})
		return
	}

	var newSong models.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	logger.InfoLogger.Println(newSong)

	song, err := h.Service.UpdateSong(songID, newSong)
	if err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	logger.InfoLogger.Println(song)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": song})
}

// GetAllSongs godoc
// @Summary Get all songs with pagination
// @Description Retrieves a list of songs with pagination support
// @Tags Songs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of records per page" default(10)
// @Success 200 {object} []models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/songs [get]
func (h *Handler) GetAllSongs(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit < 1 {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid limit"})
		return
	}
	logger.InfoLogger.Println("page %s and limit %s", page, limit)
	offset := (page - 1) * limit

	songs, err := h.Service.GetAllSongs(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Failed to fetch songs"})
		return
	}

	logger.InfoLogger.Println("songs :", songs)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": songs})

}

// GetSongsWithFilter retrieves a list of songs with optional filters.
// @Summary Get Songs with Filters
// @Description Retrieve songs by applying filters and pagination.
// @Tags Songs
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param limit query int false "Number of records per page (default is 20)"
// @Param group_name query string false "Filter by group name (partial match)"
// @Param song_name query string false "Filter by song name (partial match)"
// @Param release_date query string false "Filter by release date (format: YYYY-MM-DD)"
// @Param song_text query string false "Filter by song text (partial match)"
// @Param link query string false "Filter by link (partial match)"
// @Success 200 {object} []models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/songs/filter [get]
func (h *Handler) GetSongsWithFilter(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit < 1 {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid limit"})
		return
	}
	logger.InfoLogger.Println("page %s and limit %s", page, limit)
	offset := (page - 1) * limit

	var filters struct {
		Group       string    `form:"group_name"`
		SongName    string    `form:"song_name"`
		ReleaseDate time.Time `form:"release_date"`
		SongText    string    `form:"song_text"`
		Link        string    `form:"link"`
	}

	if err = c.ShouldBindQuery(&filters); err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid query parameters"})
		return
	}
	var requestBody models.RequestSong
	requestBody.Group = filters.Group
	requestBody.SongText = filters.SongText
	requestBody.ReleaseDate = filters.ReleaseDate
	requestBody.SongName = filters.SongName
	requestBody.Link = filters.Link

	songs, err := h.Service.GetSongsWithFilter(requestBody, limit, offset)
	if err != nil {
		logger.DebugLogger.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": "Could not fetch songs"})
		return
	}

	logger.InfoLogger.Println("songs :", songs)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": songs})
}
