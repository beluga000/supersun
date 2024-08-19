package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"sunny.ksw.kr/inits"

	_ "sunny.ksw.kr/work/pro1/main/docs"          // Swagger docs
	HANDLER "sunny.ksw.kr/work/pro1/main/handler" // 핸들러 패키지
)

func init() {
	mongo_uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	mongo_db := "local"
	inits.MongoInit(mongo_uri, mongo_db)
}

// @title Fiber Example API
// @version 1.0
// @description API 설명
// @host localhost:8090
// @BasePath /api/v1
func main() {
	engine := html.New(os.Getenv("VPATH"), ".html")

	app := fiber.New(fiber.Config{
		AppName:                  "big_money",
		CaseSensitive:            true,
		StrictRouting:            true,
		ServerHeader:             "Fiber",
		ProxyHeader:              "X-Forwarded-For",
		EnableTrustedProxyCheck:  true,
		BodyLimit:                1024 * 1024 * 1000,
		Views:                    engine,
		EnableSplittingOnParsers: true,
	})

	// CORS 설정
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:9000, http://localhost:8080",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,name",
		AllowMethods:     "GET,POST,HEAD,DELETE,PATCH",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Disposition",
	}))

	// Swagger UI 초기화
	app.Get("/swagger/*", swagger.HandlerDefault) // Swagger 핸들러 등록

	apiV1 := app.Group("/api/v1") // API 그룹

	// 기본 엔드포인트
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// 핸들러 등록
	HANDLER.Member(apiV1)
	HANDLER.Card(apiV1) // 카드 핸들러 등록
	HANDLER.Bank(apiV1)

	app.Listen(":8090") // 서버 실행
}
