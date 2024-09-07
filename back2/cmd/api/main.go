package main

import (
	"log"
	"os"

	"back2/internal/infrastructure/database"
	"back2/internal/infrastructure/router"
	"back2/internal/interface/handler"
	"back2/internal/interface/repository"
	"back2/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// データベース接続の初期化
	dbEngine, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}

	// データベースの同期
	if err := database.SyncDatabase(dbEngine); err != nil {
		log.Fatalf("Failed to synchronize database: %v", err)
	}

	// リポジトリの初期化
	castRepo := repository.NewXormCastRepository(dbEngine)
	customerRepo := repository.NewXormCustomerOrderRepository(dbEngine)
	groupRepo := repository.NewXormGroupRepository(dbEngine)
	orderRepo := repository.NewXormOrderRepository(dbEngine)
	customerOrderRepo := repository.NewXormCustomerOrderRepository(dbEngine)
	staffRepo := repository.NewXormStaffRepository(dbEngine)
	storeRepo := repository.NewXormStoreRepository(dbEngine)
	mediaRepo := repository.NewXormMediaRepository(dbEngine)
	masterRepo := repository.NewXormMasterRepository(dbEngine)
	tipsRepo := repository.NewXormTipsRepository(dbEngine)
	// ユースケースの初期化
	castUseCase := usecase.NewCastUseCase(castRepo)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo, orderRepo)
	groupUseCase := usecase.NewGroupUseCase(groupRepo)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, customerOrderRepo, storeRepo)
	staffUseCase := usecase.NewStaffUseCase(staffRepo)
	storeUseCase := usecase.NewStoreUseCase(storeRepo)
	mediaUseCase := usecase.NewMediaUseCase(mediaRepo)
	masterUseCase := usecase.NewMasterUseCase(masterRepo)
	tipsUseCase := usecase.NewTipsUseCase(tipsRepo)
	// ハンドラーの初期化
	castHandler := handler.NewCastHandler(castUseCase)
	customerHandler := handler.NewCustomerHandler(customerUseCase, orderUseCase)
	groupHandler := handler.NewGroupHandler(groupUseCase)
	orderHandler := handler.NewOrderHandler(orderUseCase)
	staffHandler := handler.NewStaffHandler(staffUseCase)
	authHandler := handler.NewAuthHandler(staffUseCase)
	storeHandler := handler.NewStoreHandler(storeUseCase)
	mediaHandler := handler.NewMediaHandler(mediaUseCase)
	masterHandler := handler.NewMasterHandler(masterUseCase)
	tipsHandler := handler.NewTipsHandler(tipsUseCase)
	// Ginエンジンの設定
	engine := gin.Default()
	config := cors.DefaultConfig()
	frontUrl := os.Getenv("FRONT_URL")
	log.Println(frontUrl)
	config.AllowOrigins = []string{frontUrl} // フロントエンドのオリジンを指定
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	engine.Use(cors.New(config))

	// ルーターのセットアップ
	router.SetupRouter(engine,
		castHandler,
		customerHandler,
		groupHandler,
		orderHandler,
		staffHandler,
		authHandler,
		storeHandler,
		mediaHandler,
		masterHandler, tipsHandler)

	// サーバーの起動
	if err := engine.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
