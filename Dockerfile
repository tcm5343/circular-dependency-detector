FROM golang:1.21

# ARG ADJACENCY_LIST_PATH
# RUN ls
# RUN echo ${ADJACENCY_LIST_PATH}

WORKDIR /app

COPY ./list ./list
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY ./pkg ./pkg
COPY ./visualizer ./visualizer

RUN CGO_ENABLED=0 GOOS=linux go build -o ./circular-dependency-detector

CMD ["./circular-dependency-detector"]
