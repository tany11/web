package routers

import (
	"back2/controllers"
	"back2/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		// CustomerControllerのルーティング
		customerController := controllers.CustomerController{DB: service.DbEngine}
		v1.POST("/customers", customerController.CreateCustomer)
		v1.GET("/customers", customerController.GetAllCustomers)
		v1.GET("/customers/:id", customerController.GetCustomer)
		v1.GET("/customers/phone/:phone", customerController.GetCustomerPhone)
		v1.PUT("/customers/:id", customerController.UpdateCustomer)
		v1.DELETE("/customers/:id", customerController.DeleteCustomer)

		// OrderControllerのルーティング
		orderController := controllers.OrderController{DB: service.DbEngine}
		v1.POST("/orders", orderController.CreateOrder)
		v1.GET("/orders", orderController.GetAllOrders)
		v1.GET("/orders/:email", orderController.GetOrder)
		v1.PUT("/orders/:id", orderController.UpdateOrder)
		v1.DELETE("/orders/:id", orderController.DeleteOrder)

		// StoreControllerのルーティング
		storeController := controllers.StoreController{DB: service.DbEngine}
		v1.POST("/stores", storeController.CreateStore)
		v1.GET("/stores", storeController.GetAllStores)
		v1.GET("/stores/:id", storeController.GetStore)
		v1.PUT("/stores/:id", storeController.UpdateStore)
		v1.DELETE("/stores/:id", storeController.DeleteStore)

		// GroupLoginControllerのルーティング
		groupController := controllers.GroupLoginController{DB: service.DbEngine}
		v1.POST("/groups", groupController.CreateGroupLogin)
		v1.GET("/groups", groupController.GetAllGroupLogins)
		v1.GET("/groups/:id", groupController.GetGroupLogin)
		v1.PUT("/groups/:id", groupController.UpdateGroupLogin)
		v1.DELETE("/groups/:id", groupController.DeleteGroupLogin)

		// CastLoginControllerのルーティング
		castService := service.NewCastLoginService(service.DbEngine)
		castController := controllers.NewCastLoginController(castService)
		v1.POST("/casts", castController.CreateCastLogin)
		v1.GET("/casts", castController.GetAllCastLogins)
		v1.GET("/casts/:id", castController.GetCastLogin)
		v1.PUT("/casts/:id", castController.UpdateCastLogin)
		v1.DELETE("/casts/:id", castController.DeleteCastLogin)

		// StaffLoginControllerのルーティング
		staffController := controllers.StaffLoginController{DB: service.DbEngine}
		v1.POST("/staff", staffController.CreateStaffLogin)
		v1.GET("/staff", staffController.GetAllStaffLogins)
		v1.GET("/staff/:id", staffController.GetStaffLogin)
		v1.PUT("/staff/:id", staffController.UpdateStaffLogin)
		v1.DELETE("/staff/:id", staffController.DeleteStaffLogin)
	}

}
