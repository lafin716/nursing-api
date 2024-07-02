FROM golang:1.21

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go 

EXPOSE 8080
CMD ["./main"]