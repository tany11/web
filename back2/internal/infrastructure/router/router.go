package router

import (
	"time"

	"back2/internal/interface/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SetupRouter configures all the routes for the application
func SetupRouter(
	engine *gin.Engine,
	castHandler *handler.CastHandler,
	customerHandler *handler.CustomerHandler,
	groupHandler *handler.GroupHandler,
	orderHandler *handler.OrderHandler,
	staffHandler *handler.StaffHandler,
) {
	// Apply global middleware
	engine.Use(RecordUaAndTime)

	// API v1 routes
	v1 := engine.Group("/api/v1")
	{
		// Customer routes
		v1.POST("/customers", customerHandler.Create)
		v1.GET("/customers", customerHandler.GetAll)
		v1.GET("/customers/:id", customerHandler.Get)
		v1.GET("/customers/phone/:phone", customerHandler.GetByPhone)
		v1.PUT("/customers/:id", customerHandler.Get)
		v1.DELETE("/customers/:id", customerHandler.Delete)

		// Order routes
		v1.POST("/orders", orderHandler.Create)
		v1.GET("/orders", orderHandler.GetAll)
		v1.GET("/orders/:id", orderHandler.Get)
		v1.PUT("/orders/:id", orderHandler.Update)
		v1.DELETE("/orders/:id", orderHandler.Delete)

		// Group routes
		v1.POST("/groups", groupHandler.Create)
		v1.GET("/groups", groupHandler.GetAll)
		v1.GET("/groups/:id", groupHandler.Get)
		v1.PUT("/groups/:id", groupHandler.Update)
		v1.DELETE("/groups/:id", groupHandler.Delete)

		// Cast routes
		v1.POST("/casts", castHandler.Create)
		v1.GET("/casts", castHandler.GetAll)
		v1.GET("/casts/:id", castHandler.Get)
		v1.PUT("/casts/:id", castHandler.Update)
		v1.DELETE("/casts/:id", castHandler.Delete)

		// Staff routes
		v1.POST("/staff", staffHandler.Create)
		v1.GET("/staff", staffHandler.GetAll)
		v1.GET("/staff/:id", staffHandler.Get)
		v1.PUT("/staff/:id", staffHandler.Update)
		v1.DELETE("/staff/:id", staffHandler.Delete)
	}
}

// RecordUaAndTime is a middleware to log the user agent and request time
func RecordUaAndTime(c *gin.Context) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	start := time.Now()
	ua := c.GetHeader("User-Agent")
	c.Next()

	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("UA", ua),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Since(start)),
	)
}
