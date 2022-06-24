FROM golang

WORKDIR /usr/app

COPY . .

RUN go build -o notes-service ./cmd/service

CMD ./user-auth