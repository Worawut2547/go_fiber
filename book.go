package main

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book {
	{ID: 1, Title: "Jack th Ripper", Author: "Jackson"},
	{ID: 2, Title: "Ironman", Author: "Marvel"},
	{ID: 3, Title: "Focus", Author: "John"},
	{ID: 4, Title: "John Wick", Author: "JJ"},
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId , err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _ , book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books , *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	id , err := strconv.Atoi(c.Params("id"))

	if err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i,book := range books{
		if book.ID == id{
			book.Title = bookUpdate.Title
			book.Author = bookUpdate.Author
			books[i] = book
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error {
	id , err := strconv.Atoi(c.Params("id"))

	if err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookDelete := new(Book)

	if err := c.BodyParser(bookDelete); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i,book := range books{
		if book.ID == id{
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}