FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tydy
RUN go build -o main .

CMD ["postgres:6000", "--", "./main"]