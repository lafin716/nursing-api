FROM golang:1.21

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN apt-get install -y tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go 

EXPOSE 8080
CMD ["./main"]