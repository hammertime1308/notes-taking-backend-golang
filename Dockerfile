FROM golang:1.20

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o notes-server
RUN mkdir -p /output/config/
RUN cp -r ./config /output
RUN cp notes-server /output

EXPOSE 8080

CMD ["./notes-server"]
