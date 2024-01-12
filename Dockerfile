FROM golang:alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o ./out/dist .

EXPOSE 8080

CMD ["./out/dist"]