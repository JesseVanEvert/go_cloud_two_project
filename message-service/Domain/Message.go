package Domain

type Message struct {
	MessageID     int    `json:"messageID"`
	LecturerEmail string `json:"lecturerEmail"`
	Content       string `json:"content"`
}
