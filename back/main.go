package main

import (
	middleware "back/middlewares"
	"back/models"
	"back/routers"
	"back/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// データベース接続の初期化
	dbEngine, err := initDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	service.SetDbEngine(dbEngine)
	models.DBEngine = service.GetDbEngine()

	engine := gin.Default()
	// ミドルウェア
	engine.Use(middleware.RecordUaAndTime)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Vueアプリのデフォルトポート
	engine.Use(cors.New(config))

	// ルーターのセットアップ
	routers.SetupRouter(engine)

	// サーバーの起動
	if err := engine.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func initDatabase() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", "user:password@/dbname?charset=utf8")
}
