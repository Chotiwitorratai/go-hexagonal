package main

import (
	"go-hexagonal/database"
	"go-hexagonal/handler"
	"go-hexagonal/repository"
	router "go-hexagonal/routes"
	"go-hexagonal/service"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Use(logger.New())
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	cr := repository.NewCustomerRepositoryDB(db)
	ir := repository.NewItemRepositoryDB(db, redisClient)
	or := repository.NewOrderRepositoryDB(db)
	cs := service.NewCustomerService(cr)
	is := service.NewItemService(ir)
	os := service.NewOrderService(or)
	h := handler.NewHandler(cs, is, os)
	router.SetupRouter(app, h)
	log.Fatal(app.Listen(":3004"))
	// app.Listen(":3000")
}
