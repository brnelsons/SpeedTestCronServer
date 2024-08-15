FROM golang:1.23.0-alpine as go-builder

WORKDIR /app
COPY . /app/
RUN go build ./src -o /dist/speedtest-server


FROM rockylinux:latest

WORKDIR /app
COPY resources /app/resources
COPY --from=go-builder /dist/speedtest-server /app/speedtest-server
RUN chmod +x /app/speedtest-server

EXPOSE 80
CMD ["speedtest-server"]