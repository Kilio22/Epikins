FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY config ./config
COPY internal ./internal
COPY pkg ./pkg

RUN go build -o ./epikins_api cmd

EXPOSE 80
CMD ["/app/epikins_api"]
