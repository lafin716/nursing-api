FROM golang:1.21.1-alpine

RUN go get -u
COPY ./../../ ./src
WORKDIR src
RUN go build -o server

EXPOSE 80

CMD ["./server"]