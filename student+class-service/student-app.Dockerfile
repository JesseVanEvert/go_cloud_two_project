# Base image
FROM python:3.8-slim-buster

# Set the working directory
WORKDIR /

# Copy the requirements file to the working directory
COPY requirements.txt .

# Install the dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the application code to the working directory
COPY . .
# Set the environment variables for RabbitMQ
ENV RABBITMQ_HOST=localhost
ENV RABBITMQ_PORT=5672
ENV RABBITMQ_USERNAME=guest
ENV RABBITMQ_PASSWORD=guest

# Expose the port for the application
EXPOSE 8000

# Start the application and connect to RabbitMQ
CMD python app.py
#FROM python:3.9-slim-buster
#
#RUN apt-get update && apt-get install -y curl gnupg
#RUN curl -s https://dl.bintray.com/rabbitmq/Keys/rabbitmq-release-signing-key.asc | apt-key add -
#RUN echo "deb https://dl.bintray.com/rabbitmq/debian buster main" > /etc/apt/sources.list.d/bintray.rabbitmq.list
#RUN apt-get update && apt-get install -y rabbitmq-server
#
#COPY requirements.txt /app/
#WORKDIR /app
#RUN pip install --no-cache-dir -r requirements.txt
#
#COPY . /app/
#CMD ["python", "app.py"]