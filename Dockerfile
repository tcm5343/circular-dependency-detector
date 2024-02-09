FROM golang:1.21

ARG ADJACENCY_LIST_PATH
RUN echo ${ADJACENCY_LIST_PATH}

WORKDIR /app

COPY ./list ./list
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY ./pkg ./pkg
COPY ./visualizer ./visualizer

COPY entrypoint.sh ./
RUN chmod +x ./entrypoint.sh

RUN CGO_ENABLED=0 GOOS=linux go build -o ./circular-dependency-detector


ENTRYPOINT ["bash", "./entrypoint.sh"]
