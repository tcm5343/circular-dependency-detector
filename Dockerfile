FROM golang:1.21

# ARG ADJACENCY_LIST_PATH
# RUN ls
# RUN echo ${ADJACENCY_LIST_PATH}

WORKDIR /app

COPY ./list ./list
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY entrypoint.sh ./
COPY ./pkg ./pkg
COPY ./visualizer ./visualizer

RUN chmod +x ./entrypoint.sh
RUN CGO_ENABLED=0 GOOS=linux go build -o ./circular-dependency-detector
RUN ls

ENTRYPOINT ["bash", "./entrypoint.sh"]
