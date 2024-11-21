package router

import (
	"github.com/bahodurnazarov/effictiveMobile/internal/app/handler"
	"github.com/bahodurnazarov/effictiveMobile/pkg/middleware"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

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
