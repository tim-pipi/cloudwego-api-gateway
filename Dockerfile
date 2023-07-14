FROM golang:1.20.5
WORKDIR /app
COPY . .
RUN go build -o cwgo ./cmd/cloudwego
EXPOSE 8080
CMD ["./cwgo"]