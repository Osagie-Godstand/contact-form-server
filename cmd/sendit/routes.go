package main

import (
	"database/sql"

	"github.com/Osagie-Godstand/contact-form/db"
	"github.com/Osagie-Godstand/contact-form/handler"
	"github.com/gofiber/fiber/v2"
)

func initializeRouter(dbConn *sql.DB) *fiber.App {
	app := fiber.New()

	contactsRepository := db.NewPostgresContactsRepository(dbConn)
	contactHandler := handler.NewContactRequestHandler(contactsRepository)

	app.Post("/create_contact", contactHandler.CreateContactHandler)
	app.Get("/get_contacts", contactHandler.GetContactsHandler)
	app.Get("/get_contact/:id", contactHandler.GetContactByIDHandler)
	app.Delete("/delete_contact/:id", contactHandler.DeleteContactByIDHandler)

	return app
}
