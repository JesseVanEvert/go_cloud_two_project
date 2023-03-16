from datetime import datetime,time
from enum import Enum

from marshmallow_sqlalchemy import fields

from config import db, ma, app

class Operation(Enum):
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
    
class Student(db.Model):
    __tablename__ = "student"
    id = db.Column(db.Integer, primary_key=True)
    lname = db.Column(db.String(32))
    fname = db.Column(db.String(32))
    email = db.Column(db.String(32))
    deleted_at = db.Column(db.String(250),nullable = True)
    classroom_id = db.Column(db.Integer, db.ForeignKey("classroom.id"), nullable = False)

class StudentSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Student
        load_instance = True
        sqla_session = db.session
        include_fk = True
    

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
    )
    
class ClassSchema(ma.SQLAlchemyAutoSchema):
    class Meta:
        model = Classroom
        load_instance = True
        sqla_session = db.session
        include_relationships = True

    student = fields.Nested(StudentSchema,many=True)



classroom_schema = ClassSchema() #1 object
class_schema = ClassSchema(many=True) #many objects
student_schema = StudentSchema()
people_schema = StudentSchema(many=True)
with app.app_context():
    db.create_all()

