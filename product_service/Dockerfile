FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tydy
RUN go build -o main .

CMD ["postgres:7000", "--", "./main"]