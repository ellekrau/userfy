FROM golang:1.18-alpine

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go install -tags 'mysql,postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    && go mod download && go mod verify

COPY . .
RUN cd src && go build -v

EXPOSE $PORT

ENTRYPOINT ["tail","-f","/dev/null"]