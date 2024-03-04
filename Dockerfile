FROM golang:latest

WORKDIR .

# Download dependencies
COPY /go.mod /go.sum ./
RUN go mod download

# Build the app
COPY / ./
RUN go build -o main .

EXPOSE 8081

# Sleep for two seconds to give database time to start
CMD ["sh", "-c", "sleep 2 && ./main"]