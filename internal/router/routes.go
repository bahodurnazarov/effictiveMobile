package router

import (
	_ "github.com/bahodurnazarov/effictiveMobile/docs"
	"github.com/bahodurnazarov/effictiveMobile/internal/app/handler"
	"github.com/bahodurnazarov/effictiveMobile/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title Music Info API
// @version 0.0.1
// @description This is a Music Info API server.
// @host http://localhost:8081
// @BasePath /api/
func Init(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ping", Ping)
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8081/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "works!",
			"message": "route not found!"})
	})

	api := r.Group("/api")
	{
		api.POST("/song", h.CreateSong)
		api.DELETE("/song/:id", h.DeleteSong)
		api.GET("/song/:id", h.GetSongByID)
		api.PUT("/song/:id", h.UpdateSong)
		//api.GET("/songs", h.GetAllSongs)
	}

	return r
}

// PingExample ping
// @Summary Ping
// @Schemes
// @Description do ping
// @Tags Routes
// @Accept json
// @Produce json
// @Success 200
// @Router /ping [get]
// Init foutes function
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
