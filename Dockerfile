FROM docker.io/golang:1.22 AS build

COPY ./ /app/

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/circular-dependency-detector

# FROM docker.io/alpine:3.20.2 as test

# COPY --from=build /app/entrypoint.sh /app/entrypoint.sh
# RUN chmod +x /app/entrypoint.sh

# COPY --from=build /app/circular-dependency-detector /app/circular-dependency-detector
# COPY --from=build /app/testing /app/testing
# ENTRYPOINT ["sh", "/app/entrypoint.sh"]

FROM docker.io/alpine:3.20.2

COPY --from=build /app/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

COPY --from=build /app/circular-dependency-detector /app/circular-dependency-detector
COPY --from=build /app/testing /app/testing
ENTRYPOINT ["sh", "/app/entrypoint.sh"]
