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