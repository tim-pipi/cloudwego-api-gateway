FROM golang:1.20.5
WORKDIR /app
COPY . .
RUN go build -o http
EXPOSE 8080
CMD ["./http"]