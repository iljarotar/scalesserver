FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY server/ server/

RUN CGO_ENABLED=0 GOOS=linux go build -o /scalesserver

EXPOSE 8080

CMD ["/scalesserver"]
