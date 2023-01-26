package Domain

type Message struct {
	MessageID  int    `json:"messageID"`
	LecturerID int    `json:"lecturerID"`
	Content    string `json:"content"`
}
