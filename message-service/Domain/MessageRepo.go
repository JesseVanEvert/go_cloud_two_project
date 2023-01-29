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
	FindMessageByLecturerId(id string) (*Message, *AppError)
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
		err = rows.Scan(&message.MessageID, &message.LecturerID, &message.Content)
		if err != nil {
			log.Println("Error scanning rows" + err.Error())
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (ch MessageRepoDB) FindById(ID string) (*Message, *AppError) {
	var message Message
	err := ch.db.QueryRow("SELECT * FROM message WHERE messageID=?", ID).Scan(&message.MessageID, &message.LecturerID, &message.Content)
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

func (ch MessageRepoDB) FindMessageByLecturerId(ID string) (*Message, *AppError) {
	var message Message
	err := ch.db.QueryRow("SELECT * FROM message WHERE lecturerID=?", ID).Scan(&message.MessageID, &message.LecturerID, &message.Content)
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
