FROM golang:1.21

# ARG ADJACENCY_LIST_PATH
# RUN ls
# RUN echo ${ADJACENCY_LIST_PATH}

WORKDIR /app

# RUN chmod +x run.sh
COPY ./list ./list
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
COPY run.sh ./
COPY ./pkg ./pkg
COPY ./visualizer ./visualizer

RUN chmod +x ./run.sh
RUN CGO_ENABLED=0 GOOS=linux go build -o ./circular-dependency-detector
RUN ls

ENTRYPOINT ["./run.sh"]
