package router

import (
	"time"

	"back2/internal/infrastructure/middleware"
	"back2/internal/infrastructure/websocket"
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
	tipsHandler *handler.TipsHandler,
	wsServer *websocket.Server,
) {
	// Apply global middleware
	engine.Use(RecordUaAndTime)
	engine.Use(middleware.ErrorHandler()) // 追加

	public := engine.Group("/api/v1")
	{
		public.POST("/login", authHandler.Login)
		public.POST("/logout", authHandler.Logout)
	}

	// WebSocket用のルートを設定
	engine.GET("/api/v1/ws", func(c *gin.Context) {
		wsServer.Serve(c.Writer, c.Request)
	})

	// 認証が必要なルートはこの後に設定
	protected := engine.Group("/api/v1")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		// Customer routes
		protected.POST("/customers", customerHandler.Create)
		protected.GET("/customers", customerHandler.GetAll)
		protected.GET("/customers/:id", customerHandler.Get)
		protected.GET("/customers/phone/:phone", customerHandler.GetByPhone)
		protected.GET("/customers/detail/:phone", customerHandler.GetDetail)
		protected.GET("/customers/list", customerHandler.List)
		protected.GET("/customers/search", customerHandler.GetSearchList)
		protected.PUT("/customers/:id", customerHandler.Get)
		protected.DELETE("/customers/:id", customerHandler.Delete)

		// Order routes
		protected.POST("/orders", orderHandler.Create)
		protected.PUT("/orders/:id", orderHandler.Update)
		protected.PUT("/orders/:id/completion", orderHandler.UpdateCompletionFlg)
		protected.PUT("/orders/:id/delete", orderHandler.DeleteFlg)
		protected.GET("/orders", orderHandler.GetAll)
		protected.GET("/orders/:id", orderHandler.Get)
		protected.DELETE("/orders/:id", orderHandler.Delete)
		protected.GET("/orders/reserved", orderHandler.ListReserved)
		protected.GET("/orders/scheduled", orderHandler.ListSchedule)
		protected.PUT("/orders/:id/schedule", orderHandler.UpdateSchedule)

		// Group routes
		protected.POST("/group", groupHandler.Create)
		protected.GET("/group", groupHandler.GetAll)
		protected.GET("/group/:id", groupHandler.Get)
		protected.PUT("/group/:id", groupHandler.Update)
		protected.DELETE("/group/:id", groupHandler.Delete)

		// Cast routes
		protected.POST("/cast", castHandler.Create)
		protected.GET("/cast", castHandler.GetAll)
		protected.GET("/cast/:id", castHandler.Get)
		protected.PUT("/cast/:id", castHandler.Update)
		protected.DELETE("/cast/:id", castHandler.Delete)
		protected.GET("/cast/dropdown", castHandler.ListForDropdown)
		protected.PUT("/cast/:id/working", castHandler.UpdateWorkingFlg)

		// Staff routes
		protected.POST("/staff", staffHandler.Create)
		protected.GET("/staff", staffHandler.GetAll)
		protected.GET("/staff/:id", staffHandler.Get)
		protected.PUT("/staff/:id", staffHandler.Update)
		protected.DELETE("/staff/:id", staffHandler.Delete)
		protected.GET("/staff/dropdown", staffHandler.ListForDropdown)

		// Store routes
		protected.POST("/store", storeHandler.Create)
		protected.GET("/store", storeHandler.GetAll)
		protected.GET("/store/:id", storeHandler.Get)
		protected.PUT("/store/:id", storeHandler.Update)
		protected.DELETE("/store/:id", storeHandler.Delete)
		protected.GET("/store/dropdown", storeHandler.ListForDropdown)

		// Media routes
		protected.POST("/media", mediaHandler.Create)
		protected.GET("/media", mediaHandler.GetAll)
		protected.GET("/media/:id", mediaHandler.Get)
		protected.PUT("/media/:id", mediaHandler.Update)
		protected.DELETE("/media/:id", mediaHandler.Delete)
		protected.GET("/media/dropdown", mediaHandler.ListForDropdown)

		// Master routes
		protected.POST("/master", masterHandler.Create)
		protected.GET("/master", masterHandler.GetAll)
		protected.GET("/master/:id", masterHandler.Get)
		protected.PUT("/master/:id", masterHandler.Update)
		protected.DELETE("/master/:id", masterHandler.Delete)
		protected.GET("/master/usage", masterHandler.ListForUsage)

		// Tips routes
		protected.POST("/tips", tipsHandler.Create)
		protected.GET("/tips", tipsHandler.List)
		protected.GET("/tips/schedule", tipsHandler.ListSchedule)
		protected.PUT("/tips/:id", tipsHandler.Update)
		protected.DELETE("/tips/:id", tipsHandler.Delete)
		protected.PUT("/tips/:id/schedule", tipsHandler.UpdateMemo)
		protected.PUT("/tips/:id/completion", tipsHandler.UpdateCompletionFlg)
		protected.PUT("/tips/:id/delete", tipsHandler.DeleteFlg)
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
