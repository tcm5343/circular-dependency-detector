FROM golang:1.21

ARG test
RUN echo ${test}

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


ENTRYPOINT ["sh", "./entrypoint.sh"]
