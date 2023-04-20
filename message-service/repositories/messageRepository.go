package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"messages/domain"

	_ "github.com/go-sql-driver/mysql"
)

type MessageRepository interface {
	FindAll() ( []*domain.Message, error)
	FindById(id string) (*domain.Message, error)
	FindMessageByLecturerEmail(lecturerEmail string) ([]*domain.Message, error)
}

type DefaultMessageRepository struct {
	db *sql.DB
}

func (ch DefaultMessageRepository) FindAll() ([]*domain.Message, error) {
	findall_sql := "SELECT * FROM message"
	rows, err := ch.db.Query(findall_sql)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	messages := make([]*domain.Message, 0)
	for rows.Next() {
		var message domain.Message
		err = rows.Scan(&message.MessageID, &message.LecturerEmail, &message.ToEmail, &message.Content)

		if err != nil {
			return nil, fmt.Errorf("finding messages: %w", err)
		}

		messages = append(messages, &message)
	}

	return messages, nil
}

func (ch DefaultMessageRepository) FindById(ID string) (*domain.Message, error) {
	var message domain.Message

	err := ch.db.QueryRow("SELECT * FROM message WHERE messageID=?", ID).Scan(&message.MessageID, &message.LecturerEmail, &message.ToEmail, &message.Content)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error: %w", err)
		} else {
			log.Println("Error scanning rows ById" + err.Error())
			return nil, fmt.Errorf("error: %w", err)
		}
	}

	return &message, nil
}

func (ch DefaultMessageRepository) FindMessageByLecturerEmail(lecturerEmail string) ([]*domain.Message, error) {
	var messages []*domain.Message

	rows, err := ch.db.Query("SELECT * FROM message WHERE lecturerEmail=?", lecturerEmail)

	if err != nil {
		return nil, fmt.Errorf("creating classroom: %w", err)
	}
	
	defer rows.Close()
	for rows.Next() {
		var message domain.Message

		err := rows.Scan(&message.MessageID, &message.LecturerEmail, &message.ToEmail, &message.Content)

		if err != nil {
			return nil, fmt.Errorf("creating classroom: %w", err)
		}
		messages = append(messages, &message)
	}
	return messages, nil
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return DefaultMessageRepository{db}
}
