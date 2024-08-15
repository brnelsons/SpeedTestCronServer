FROM golang:1.22.0-alpine AS go-builder
WORKDIR /build
COPY . /build/
RUN go build -o /dist/server .


FROM rockylinux:9-minimal
WORKDIR /app
COPY static /app/static
COPY --from=go-builder /dist/server /app/speedtest
RUN chmod +x /app/speedtest

EXPOSE 8080
CMD ["/app/speedtest"]