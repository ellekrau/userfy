FROM golang:1.18-alpine

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./

COPY . .
RUN cd src && go build -v -o /app

EXPOSE $PORT

CMD ["./userfy"]