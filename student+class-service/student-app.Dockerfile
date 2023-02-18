FROM python:3.8-slim-buster

# Install RabbitMQ
RUN apt-get update && \
    apt-get install -y rabbitmq-server

# Set the working directory to /app
WORKDIR /app

# Copy the requirements file to the working directory
COPY requirements.txt .

# Install the required packages
RUN pip install -r requirements.txt

# Copy the application code into the container
COPY . .

# Expose the required ports
EXPOSE 8000
EXPOSE 5672

# Start the application
CMD ["sh", "-c", "rabbitmq-server start && python app.py"]