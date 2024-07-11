FROM golang:1.22.3

COPY ./app /app
WORKDIR /app
EXPOSE 8080

CMD go mod download && go mod verify && go run .