package data

import "github.com/google/uuid"

type ContactRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
	Message     string `json:"message"`
}

type ContactsReceived struct {
	ID          uuid.UUID `json:"id"`
	Email       *string   `json:"email"`
	Name        *string   `json:"name"`
	PhoneNumber *string   `json:"phonenumber"`
	Message     *string   `json:"message"`
}

func NewUUID() uuid.UUID {
	return uuid.New()
}

type ContactID struct {
	ID uuid.UUID `json:"id"`
}
