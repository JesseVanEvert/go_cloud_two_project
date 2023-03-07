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
### Start the application
CMD ["sh", "-c", "rabbitmq-server start && python app.py"]
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