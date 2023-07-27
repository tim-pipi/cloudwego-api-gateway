FROM golang:1.20.5
WORKDIR /app
COPY . .
RUN go build -o rpc-server .
CMD ["./rpc-server"]
