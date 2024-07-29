FROM golang:1.22.5 AS builder
COPY ./src /app
WORKDIR /app
RUN go mod download
RUN go mod verify
RUN go build -o server .

FROM golang:1.22.5
COPY --from=builder /app/server /app/server
EXPOSE 8080
CMD ["/app/server"]