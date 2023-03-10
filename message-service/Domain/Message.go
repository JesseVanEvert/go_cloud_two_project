package Domain

type Message struct {
	MessageID     int    `json:"messageID"`
	LecturerEmail string `json:"lecturerEmail"`
	ToEmail       string `json:"to"`
	Content       string `json:"content"`
}
