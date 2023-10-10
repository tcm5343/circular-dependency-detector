FROM golang:1.21

WORKDIR /app

COPY go.mod ./
COPY ./pkg ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /circular-dependency-detector

CMD ["/circular-dependency-detector"]
