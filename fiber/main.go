package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_"github.com/jigneng1/fiber-test/docs"
	
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("load .env error")
	}

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	//Login
	app.Post("/login", login)

	// Initialize in-memory data
	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	//handle middle all routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Use(checkMiddleware)

	// CRUD routes
	app.Get("/book", getBooks)
	app.Get("/book/:id", getBook)
	app.Post("/book", createBook)
	app.Put("/book/:id", updateBook)

	// File upload
	app.Post("/upload", uploadFile)

	//get Env config
	app.Get("/api/config", getConfig)

	// run server on port
	// Use the environment variable for the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	app.Listen(":" + port)

}

// Dummy user for example

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var MemberUsers = User{
	Email:    "user@example.com",
	Password: "password1234",
}
