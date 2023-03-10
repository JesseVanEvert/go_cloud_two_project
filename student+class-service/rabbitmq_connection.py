import pika

class RabbitMQConnection:
    def __init__(self, host):
        self.connection = pika.BlockingConnection(pika.ConnectionParameters(host))
        self.channel = self.connection.channel()

    def publish_message(self, queue_name, message):
        self.channel.queue_declare(queue=queue_name)
        self.channel.basic_publish(exchange='', routing_key=queue_name, body=message)

    def consume_messages(self, queue_name, callback):
        self.channel.queue_declare(queue=queue_name)
        self.channel.basic_consume(queue_name, callback)
        self.channel.start_consuming()

    def close(self):
        self.connection.close()
    """ def send_classroom(class_data):
    connection = RabbitMQConnection('localhost')
    
    # Convert class_data to JSON string
    class_data_json = json.dumps(class_data, cls=ClassSchema)

    # Publish a message to the 'student_creation' queue
    connection.publish_message('student_creation', class_data_json)

    # Consume messages from the 'student_creation' queue
    def callback(ch, method, properties, body):
        # Convert JSON string back to Python object
        class_data_received = json.loads(body, cls=ClassSchema)
        print("Received message:", class_data_received)

    connection.consume_messages('student_creation', callback)

    # Close the connection
    connection.close() 
    """