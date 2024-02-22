FROM golang:1.21.1-alpine

WORKDIR /app

COPY ./../../go.mod ./
RUN go mod download

COPY ./../../ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./../../

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/main .
EXPOSE 80
CMD ["./main"]