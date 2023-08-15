package main

import (
	"go-rest/controllers"
	"go-rest/initializers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/wpcodevo/golang-fiber/controllers"
	_ "github.com/wpcodevo/golang-fiber/initializers"
)

//	func init() {
//		config, err := initializers.LoadConfig(".")
//		if err != nil {
//			log.Fatalln("Failed to load env variable \n", err.Error())
//		}
//		initializers.ConnectDB(&config)
//	}
func main() {
	initializers.ConnectDB()

	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
	}))

	micro.Route("/notes", func(router fiber.Router) {
		router.Post("/", controllers.CreateNoteHandler)
		router.Get("", controllers.FindNotes)
	})
	micro.Route("/notes/:noteId", func(router fiber.Router) {
		router.Get("", controllers.FindNoteById)
		router.Delete("", controllers.DeleteNote)
		router.Put("", controllers.UpdateNote)
	})

	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber and GORM",
		})
	})
	log.Fatal(app.Listen(":9000"))
}
