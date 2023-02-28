package main

import(
	"fmt"
	"go-fiber-crm-basic/database"
	"go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBconn, err =gorm.Open("sqlite3", "Go_lead.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection open to db")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}