FROM golang:lastest

WORKDIR /app

COPY . .

RUN go build -o math

CMD ["./math"]