import json

from models import Student,student_schema,people_schema,Classroom
from flask import abort, make_response

from config import db
from datetime import datetime,time
from rabbitmq_connection import RabbitMQConnection



def read_all():
    student = Student.query.all()
    return people_schema.dump(student)


def create(student):
    a = datetime(1, 12, 28, 23, 55, 59, 342380)
    lname = student.get("lname")
    existing_student = Student.query.filter(Student.lname == lname).one_or_none()

    if existing_student is None:
        new_student = student_schema.load(student, session=db.session)
        new_student.deleted_at = '--'
        db.session.add(new_student)
        db.session.commit()
        classroom = Classroom.query.filter(Classroom.id == new_student.classroom_id).one_or_none()
        send_classroom(classroom)
        return student_schema.dump(new_student), 201
    else:
        abort(406, f"Person with last name {lname} already exists")


def read_one(student_id):
    student = Student.query.filter(Student.id == student_id).one_or_none()

    if student is not None:
        return student_schema.dump(student)
    else:
        abort(404, f"Person with last name {student_id} not found")


def update(student_id):
    existing_student = Student.query.filter(Student.id == student_id).one_or_none()

    if existing_student:
        #update_student = student_schema.load(student, session=db.session)
        existing_student.deleted_at = datetime.now()
        db.session.merge(existing_student)
        db.session.commit()
        return student_schema.dump(existing_student), 201
    else:
        abort(404, f"Person with last name {student_id} not found")
def send_classroom(class_data):
    connection = RabbitMQConnection('localhost')

    # Publish a message to the 'student_creation' queue
    connection.publish_message('student_creation', class_data)

    # Consume messages from the 'student_creation' queue
    def callback(ch, method, properties, body):
        print("Received message:", body)

    connection.consume_messages('student_creation', callback)

    # Close the connection
    connection.close()



# def undelete(lname):
#     existing_student = Student.query.filter(Student.lname == lname).one_or_none()

#     if existing_student:
#         #update_student = student_schema.load(student, session=db.session)
#         existing_student.deleted_at = None
#         db.session.merge(existing_student)
#         db.session.commit()
#         return student_schema.dump(existing_student), 201
#     else:
#         abort(404, f"Person with last name {lname} not found")


# def delete(lname):
#     existing_student = Person.query.filter(Person.lname == lname).one_or_none()

#     if existing_student:
#         db.session.delete(existing_student)
#         db.session.commit()
#         return make_response(f"{lname} successfully deleted", 200)
#     else:
#         abort(404, f"Person with last name {lname} not found")