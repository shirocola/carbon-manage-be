package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/shirocola/carbon-manage-be/internal/infrastructure"
	"github.com/shirocola/carbon-manage-be/internal/interfaces"
	"github.com/shirocola/carbon-manage-be/internal/usecase"
	"github.com/shirocola/carbon-manage-be/pkg/config"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to the database
	db, err := sql.Open("postgres", config.GetDatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the repository and use case
	repo := infrastructure.NewPostgresCarbonFootprintRepository(db)
	usecase := usecase.NewCarbonFootprintUsecase(repo)
	handler := interfaces.NewCarbonFootprintHandler(usecase)

	// Set up the Echo server and routes
	e := echo.New()
	e.POST("/carbon-footprints", handler.CreateCarbonFootprint)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
