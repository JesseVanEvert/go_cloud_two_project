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
