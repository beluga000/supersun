package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"sunny.ksw.kr/inits"

	HANDLER "sunny.ksw.kr/work/pro1/main/handler"
)

func init() {

	// log.Print("Server is running on port : " + os.Getenv("PORT"))

	//  MONGO_DEFAULT_URI =  "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	//	MONGO_DEFAULT_DB = "local"
	// mongo init
	mongo_uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	mongo_db := "local"

	inits.MongoInit(mongo_uri, mongo_db)
	//inits.MongoInit(mongo_uri, mongo_db)

}

func main() { //@note

	// session.Init()
	// jwt_custom.Init()
	// log.Print(os.Getenv("VPATH"))
	engine := html.New(os.Getenv("VPATH"), ".html")

	app := fiber.New(fiber.Config{
		// Prefork: true,
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
	/* cors ********************************************************************* */

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:9000, http://localhost:8080",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,name",
		AllowMethods:     "GET,POST,HEAD,DELETE,PATCH",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Disposition",
	}))

	apiV1 := app.Group("/api/v1") // /api

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	HANDLER.Member(apiV1)

	HANDLER.Card(apiV1)

	HANDLER.Bank(apiV1)

	//log.Print("Server is running on port : " + os.Getenv("PORT"))

	app.Listen(":8090") // 변경된 부분

}
