FROM golang:1.19-alpine
RUN mkdir -p /app
WORKDIR /app

COPY light-state-machine /app/light-state-machine

EXPOSE 80

ENTRYPOINT ["/app/light-state-machine"]