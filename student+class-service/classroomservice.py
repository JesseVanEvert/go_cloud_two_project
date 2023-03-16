# classesservice.py

from datetime import datetime
from flask import abort, make_response
from models import Classroom, student_schema, classroom_schema, class_schema, Student, Operation, ClassRoomQueueMessage
from rabbitmq_connection import RabbitMQConnection
import json

from config import db

def read_all():
    classroom = Classroom.query.all()
    return class_schema.dump(classroom)


def create(classroom):
    classname = classroom.get("classroom")
    existing_classroom = Classroom.query.filter(Classroom.classname == classname).one_or_none()

    if existing_classroom is None:
        new_classroom = classroom_schema.load(classroom, session=db.session)
        #new_classroom = ClassSchema.load(classroom, session=db.session)
        db.session.add(new_classroom)
        db.session.commit()
        connection = RabbitMQConnection('rabbitmq-1')

        # Convert class_data to JSON string
        class_data_serializable = classroom_schema.dump(new_classroom)
        #class_data_json = json.dumps(class_data_serializable)

        class_room_message = ClassRoomQueueMessage(Operation.CREATE, class_data_serializable)
        class_room_message_json = json.dumps(class_room_message)

        # Publish a message to the 'Classes' queue
        connection.publish_message('Classes', class_room_message_json)
        
        return classroom_schema.dump(new_classroom), 201
    else:
        abort(406, f"Classroom with Classname {classname} already exists")

def read_one(classname):
    classroom = Classroom.query.filter(Classroom.classname == classname).one_or_none()
    
    if classroom is not None:
        return classroom_schema.dump(classroom)
    else:
        abort(404, f"Classroom with name {classname} not found")


def update(classname, classroom):
    existing_classroom = Classroom.query.filter(Classroom.classname == classname).one_or_none()

    if existing_classroom:
        update_classroom = classroom_schema.load(classroom, session=db.session)
        existing_classroom.classname = update_classroom.classname
        db.session.merge(existing_classroom)
        db.session.commit()
        return classroom_schema.dump(existing_classroom), 201
    else:
        abort(404, f"Classroom with name {classname} not found")


def delete(classname):
    existing_classroom = Classroom.query.filter(Classroom.classname == classname).one_or_none()

    if existing_classroom:
        db.session.delete(existing_classroom)
        db.session.commit()
        return make_response(f"{classname} successfully deleted", 200)
    else:
        abort(404, f"Classroom with last name {classname} not found")
