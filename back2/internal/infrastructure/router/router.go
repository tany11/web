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
	authHandler *handler.AuthHandler,
	storeHandler *handler.StoreHandler,
	mediaHandler *handler.MediaHandler,
	masterHandler *handler.MasterHandler,
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
		v1.GET("/customers/list", customerHandler.List)
		v1.PUT("/customers/:id", customerHandler.Get)
		v1.DELETE("/customers/:id", customerHandler.Delete)

		// Order routes
		v1.POST("/orders", orderHandler.Create)
		v1.PUT("/orders/:id", orderHandler.Update)
		v1.PUT("/orders/:id/completion", orderHandler.UpdateCompletionFlg)
		v1.PUT("/orders/:id/delete", orderHandler.DeleteFlg)
		v1.GET("/orders", orderHandler.GetAll)
		v1.GET("/orders/:id", orderHandler.Get)
		v1.DELETE("/orders/:id", orderHandler.Delete)
		v1.GET("/orders/reserved", orderHandler.ListReserved)

		// Group routes
		v1.POST("/group", groupHandler.Create)
		v1.GET("/group", groupHandler.GetAll)
		v1.GET("/group/:id", groupHandler.Get)
		v1.PUT("/group/:id", groupHandler.Update)
		v1.DELETE("/group/:id", groupHandler.Delete)

		// Cast routes
		v1.POST("/cast", castHandler.Create)
		v1.GET("/cast", castHandler.GetAll)
		v1.GET("/cast/:id", castHandler.Get)
		v1.PUT("/cast/:id", castHandler.Update)
		v1.DELETE("/cast/:id", castHandler.Delete)
		v1.GET("/cast/dropdown", castHandler.ListForDropdown)

		// Staff routes
		v1.POST("/staff", staffHandler.Create)
		v1.GET("/staff", staffHandler.GetAll)
		v1.GET("/staff/:id", staffHandler.Get)
		v1.PUT("/staff/:id", staffHandler.Update)
		v1.DELETE("/staff/:id", staffHandler.Delete)
		v1.GET("/staff/dropdown", staffHandler.ListForDropdown)

		// Store routes
		v1.POST("/store", storeHandler.Create)
		v1.GET("/store", storeHandler.GetAll)
		v1.GET("/store/:id", storeHandler.Get)
		v1.PUT("/store/:id", storeHandler.Update)
		v1.DELETE("/store/:id", storeHandler.Delete)
		v1.GET("/store/dropdown", storeHandler.ListForDropdown)

		// Media routes
		v1.POST("/media", mediaHandler.Create)
		v1.GET("/media", mediaHandler.GetAll)
		v1.GET("/media/:id", mediaHandler.Get)
		v1.PUT("/media/:id", mediaHandler.Update)
		v1.DELETE("/media/:id", mediaHandler.Delete)
		v1.GET("/media/dropdown", mediaHandler.ListForDropdown)

		// Master routes
		v1.POST("/master", masterHandler.Create)
		v1.GET("/master", masterHandler.GetAll)
		v1.GET("/master/:id", masterHandler.Get)
		v1.PUT("/master/:id", masterHandler.Update)
		v1.DELETE("/master/:id", masterHandler.Delete)

		// Login endpoint added
		v1.POST("/login", authHandler.Login)
		v1.POST("/logout", authHandler.Logout)
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
