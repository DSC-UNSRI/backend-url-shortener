package main

import (
	"database/sql"
	"log"

	"github.com/DSC-UNSRI/backend-url-shortener/handlers"
	"github.com/DSC-UNSRI/backend-url-shortener/repository"
	"github.com/DSC-UNSRI/backend-url-shortener/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := getDBConnection()
	linkRepo := repository.NewLinkRepository(db)
	linkUc := usecase.NewLinkUsecase(*linkRepo)
	linkHandler := handlers.NewLinkHandler(*linkUc)

	app := fiber.New()

	app.Get("/links", linkHandler.GetAllLinks)
	app.Get("/:shortened_link", linkHandler.Redirect)
	app.Listen(":8000")
}

func getDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "nabiel:nabiel@tcp(localhost)/backend-url-shortener?parseTime=true")
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}

	return db
}
