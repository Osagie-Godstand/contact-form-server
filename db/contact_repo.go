package db

import (
	"database/sql"
	"fmt"

	"github.com/Osagie-Godstand/contact-form/internal/data"
	"github.com/google/uuid"
)

type ContactsRepository interface {
	InsertContactRequest(contact *data.ContactRequest) error
	GetContactsReceived() ([]data.ContactsReceived, error)
	GetContactReceivedByID(contactID uuid.UUID) (*data.ContactsReceived, error)
	DeleteContactReceivedByID(contactID uuid.UUID) error
}

type ContactsPostgresRepository struct {
	DB *sql.DB
}

func NewPostgresContactsRepository(db *sql.DB) ContactsRepository {
	return &ContactsPostgresRepository{DB: db}
}

func (r *ContactsPostgresRepository) InsertContactRequest(contact *data.ContactRequest) error {
	contactID := data.NewUUID()
	query := `
        INSERT INTO contacts_received (id, email, name, phonenumber, message)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.DB.Exec(query, contactID, contact.Email, contact.Name, contact.PhoneNumber, contact.Message)
	if err != nil {
		return fmt.Errorf("failed to insert contact request: %v", err)
	}

	return nil
}

func (r *ContactsPostgresRepository) GetContactsReceived() ([]data.ContactsReceived, error) {
	query := "SELECT id, email, name, phonenumber, message FROM your_contacts_table"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contactsinList := []data.ContactsReceived{}

	for rows.Next() {
		var contact data.ContactsReceived
		if err := rows.Scan(&contact.ID, &contact.Email, &contact.Name, &contact.PhoneNumber, &contact.Message); err != nil {
			return nil, err
		}
		contactsinList = append(contactsinList, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contactsinList, nil
}

func (r *ContactsPostgresRepository) GetContactReceivedByID(contactID uuid.UUID) (*data.ContactsReceived, error) {
	query := "SELECT id, email, name, phonenumber, message FROM your_contacts_table WHERE id = $1"
	row := r.DB.QueryRow(query, contactID)

	var contact data.ContactsReceived
	err := row.Scan(&contact.ID, &contact.Email, &contact.Name, &contact.PhoneNumber, &contact.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &contact, nil
}

func (r *ContactsPostgresRepository) DeleteContactReceivedByID(contactID uuid.UUID) error {
	query := "DELETE FROM your_contacts_table WHERE id = $1"
	result, err := r.DB.Exec(query, contactID)
	if err != nil {
		return err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numRows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
