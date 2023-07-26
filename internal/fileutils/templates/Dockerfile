FROM golang:1.20.5
WORKDIR /app
COPY . .
RUN sh ./build.sh
EXPOSE 8080
CMD ["./output/bootstrap.sh"]
