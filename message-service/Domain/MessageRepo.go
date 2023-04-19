package Domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MessageRepository interface {
	FindAll() ([]Message, error)
	FindById(id string) (*Message, *AppError)
	FindMessageByLecturerEmail(lecturerEmail string) ([]Message, *AppError)
}

type MessageRepoDB struct {
	db *sql.DB
}

func (ch MessageRepoDB) FindAll() ([]Message, error) {
	findall_sql := "SELECT * FROM message"
	rows, err := ch.db.Query(findall_sql)
	if err != nil {
		log.Println("error executing sql")
	}
	messages := make([]Message, 0)
	for rows.Next() {
		var message Message
		err = rows.Scan(&message.MessageID, &message.LecturerEmail, &message.ToEmail, &message.Content)
		if err != nil {
			log.Println("Error scanning rows" + err.Error())
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (ch MessageRepoDB) FindById(ID string) (*Message, *AppError) {
	var message Message
	err := ch.db.QueryRow("SELECT * FROM message WHERE messageID=?", ID).Scan(&message.MessageID, &message.LecturerEmail, &message.ToEmail, &message.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundError("Message not found")
		} else {
			log.Println("Error scanning rows ById" + err.Error())
			return nil, UnexpectedError("Unexpected db error")
		}
	}
	return &message, nil
}

func (ch MessageRepoDB) FindMessageByLecturerEmail(lecturerEmail string) ([]Message, *AppError) {
	var messages []Message
	rows, err := ch.db.Query("SELECT * FROM message WHERE lecturerEmail=?", lecturerEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFoundError("Message not found")
		} else {
			log.Println("Error scanning rows ById" + err.Error())
			return nil, UnexpectedError("Unexpected db error")
		}
	}
	defer rows.Close()
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.MessageID, &message.LecturerEmail, &message.ToEmail, &message.Content)
		if err != nil {
			log.Println("Error scanning rows ById" + err.Error())
			return nil, UnexpectedError("Unexpected db error")
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func NewMessageRepositoryDB() MessageRepoDB {

	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/message")

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return MessageRepoDB{db}
}
