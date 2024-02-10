FROM golang:1.21

# ARG test
# RUN echo ${test}

COPY ./ /app/
RUN chmod +x /app/entrypoint.sh

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/circular-dependency-detector


CMD ["sh", "/app/entrypoint.sh"]
