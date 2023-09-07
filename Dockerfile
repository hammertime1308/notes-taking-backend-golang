FROM golang:1.20 as builder

WORKDIR /build
COPY . .
RUN GOOS=linux go build -o notes-server .
COPY notes-server /output/
COPY config /output/



FROM alpine:latest
WORKDIR /app

COPY --from=build /output .
EXPOSE 8080
CMD ["./notes-server"]
