FROM golang:1.14

COPY . /app
WORKDIR /app

ENV GO111MODULE=on

RUN apt-get update && apt-get install build-essential -y
RUN cd cons && GOOS=linux go build -o cons
CMD ["./cons/cons"]