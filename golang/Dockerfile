FROM golang:1.19.2-alpine
WORKDIR /app
ADD ./golang .
RUN go mod download
RUN go build -o /main
EXPOSE 8090
ENTRYPOINT ["/main"]