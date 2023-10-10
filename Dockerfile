FROM golang:1.21

ARG ADJACENCY_LIST_PATH

WORKDIR /app

COPY $ADJACENCY_LIST_PATH ./
COPY go.mod ./
COPY ./pkg ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /circular-dependency-detector

CMD ["/circular-dependency-detector"]
