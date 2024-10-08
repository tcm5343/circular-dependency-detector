FROM docker.io/golang:1.23

COPY ./ /app/
RUN chmod +x /app/entrypoint.sh

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/circular-dependency-detector

ENTRYPOINT ["sh", "/app/entrypoint.sh"]
