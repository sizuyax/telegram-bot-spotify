FROM golang:latest

WORKDIR .

COPY /go.mod /go.sum ./
RUN go mod download

COPY / ./
RUN go build -o main cmd/main.go

EXPOSE 8081

CMD ["sh", "-c", "sleep 2 && ./main"]