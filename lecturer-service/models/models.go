package models

/*class Operation(Enum):
    DELETE = 0,
    CREATE = 1,
    UPDATE = 2

class ClassRoomQueueMessage(object):
    operation = None,
    class_room = None,
    class_room_id = None

    def __init__(self, operation, classroom, class_room_id):
        self.operation = operation
        self.classroom = classroom
        self.class_room_id = class_room_id

		class Classroom(db.Model):
		__tablename__ = "classroom"
		id = db.Column(db.Integer, primary_key=True)
		classname = db.Column(db.String, nullable=False)
		student = db.relationship(
			Student,
			backref="classroom",
			cascade="all, delete, delete-orphan",
			single_parent=True,
			order_by="desc(Student.email)",
		)*/

var Operation  = []string {
	"DELETE",
	"CREATE",
	"UPDATE",
}

type ClassRoomQueueMessage struct {
	Operation   string `json:"operation"`
	ClassRoom string `json:"class_room"`
	ClassRoomId int `json:"class_room_id"`
}

type ClassRoom struct {
	ID        int    
	Classname string 
}

type RequestPayload struct {
	Action  string         `json:"action"`
	Auth    AuthPayload    `json:"auth,omitempty"`
	Log     LogPayload     `json:"log,omitempty"`
	Message MessagePayload `json:"message,omitempty"`
}

type MessagePayload struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Message string   `json:"message"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LecturerPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type ClassLecturerPayload struct {
	ClassId    int `json:"class_id"`
	LecturerId int `json:"lecturer_id"`
}

type IDPayload struct {
	ID int `json:"id"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}