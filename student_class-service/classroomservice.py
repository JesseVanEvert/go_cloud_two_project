# classesservice.py

from datetime import datetime
from flask import abort, make_response
from models import Classroom, student_schema, classroom_schema, class_schema, Student, Operation, ClassRoomQueueMessage
##from rabbitmq_connection import RabbitMQConnection
import pika, os
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

        db.session.add(new_classroom)
        db.session.commit()

        url = os.environ.get('CLOUDAMQP_URL', 'amqp://guest:guest@localhost:5672/')
        params = pika.URLParameters(url)
        params.socket_timeout = 5

        connection = pika.BlockingConnection(params) # Connect to CloudAMQP
        channel = connection.channel() # start a channel
        channel.queue_declare(queue='Classes', durable=True) # Declare a queue

        class_room_message = {
            "operation": "CREATE",
            "class_room": new_classroom.classname,
            "class_room_id": new_classroom.id
        }

        class_room_message_json = json.dumps(class_room_message)
        print("Sending message:", class_room_message_json)

        # Publish a message to the 'Classes' queue
        channel.basic_publish('', 'Classes', class_room_message_json)
        
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
    existing_classroom = Classroom.query.filter(Classroom.id == id).one_or_none()

    if existing_classroom:
        db.session.delete(existing_classroom)
        db.session.commit()
        return make_response(f"{id} successfully deleted", 200)
    else:
        abort(404, f"Classroom with id {id} not found")
