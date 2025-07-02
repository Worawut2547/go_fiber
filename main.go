package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/books", getBooks)
	app.Get("/book/:id",getBook)
	app.Post("/book" ,createBook)
	app.Put("/book/:id",updateBook)
	app.Delete("/book/:id",deleteBook)

	fmt.Printf("Server is Running on Ports 6000\n")
	app.Listen(":6000")
}




