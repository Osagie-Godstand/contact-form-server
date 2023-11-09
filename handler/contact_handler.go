package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/Osagie-Godstand/contact-form/db"
	"github.com/Osagie-Godstand/contact-form/internal/data"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ContactHandler struct {
	contactsRepository db.ContactsRepository
}

func NewContactRequestHandler(contactsRepository db.ContactsRepository) *ContactHandler {
	return &ContactHandler{
		contactsRepository: contactsRepository,
	}
}

func (r *ContactHandler) CreateContactHandler(context *fiber.Ctx) error {
	contact := data.ContactRequest{}

	err := context.BodyParser(&contact)

	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	err = r.contactsRepository.InsertContactRequest(&contact)
	if err != nil {
		return fiber.ErrBadRequest
	}

	// Send email via SMTP
	if err := sendEmail(contact.Message); err != nil {
		log.Println("Failed to send email:", err)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Contact has been added successfully"})
	return nil
}

func sendEmail(message string) error {
	from := "dgodstand@gmail.com"
	to := "dgodstand@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	smtpPassword := os.Getenv("SMTP_PASSWORD") // generate your personal SMTP password using Google 2-factor authentication.
	auth := smtp.PlainAuth("", from, smtpPassword, smtpHost)

	subject := "CONTACT_REQUEST"
	body := fmt.Sprintf("Message: %s", message)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}

func (r *ContactHandler) GetContactsHandler(context *fiber.Ctx) error {
	contactsinList, err := r.contactsRepository.GetContactsReceived()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"contacts": contactsinList})
	return nil
}

func (r *ContactHandler) GetContactByIDHandler(context *fiber.Ctx) error {
	contactID, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	contact, err := r.contactsRepository.GetContactReceivedByID(contactID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}
		return fiber.ErrBadRequest
	}

	context.Status(http.StatusOK).JSON(contact)
	return nil
}

func (r *ContactHandler) DeleteContactByIDHandler(context *fiber.Ctx) error {
	contactID, err := uuid.Parse(context.Params("id"))
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	err = r.contactsRepository.DeleteContactReceivedByID(contactID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		}
		return fiber.ErrBadRequest
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Contact has been deleted successfully"})
	return nil
}
