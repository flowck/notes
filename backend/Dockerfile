FROM golang

WORKDIR /usr/app

COPY . .

RUN go build -o service ./cmd/service

CMD ./service