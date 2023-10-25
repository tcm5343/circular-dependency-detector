FROM golang:1.21

ARG ADJACENCY_LIST_PATH
RUN ls
RUN echo ${ADJACENCY_LIST_PATH}

WORKDIR /app

COPY adjacency_list.txt ./
COPY go.mod ./
COPY main.go ./
COPY ./pkg ./
RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -o /circular-dependency-detector

CMD ["/circular-dependency-detector"]
