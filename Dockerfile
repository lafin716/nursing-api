FROM golang:1.21.1-alpine

WORKDIR /app

COPY ./go.mod ./
RUN go mod download

COPY . .
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go generate ./pkg/ent
RUN wire ./internal
RUN swag init -g cmd/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/main .
EXPOSE 80
CMD ["./main"]