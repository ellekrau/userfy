FROM golang:1.18-alpine

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN cd src && go build -v -o /app

COPY .env /app

EXPOSE $PORT

CMD ["./userfy"]