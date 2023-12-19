package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jdwillmsen/fiber-crm-basic/database"
	"github.com/jdwillmsen/fiber-crm-basic/lead"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=localhost port=9500 user=postgres dbname=leads password=my-secret-pw sslmode=disable")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic("failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
