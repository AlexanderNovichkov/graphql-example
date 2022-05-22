FROM golang:1.17-alpine

WORKDIR /server

COPY .. .

RUN go mod download
RUN go build -o /app cmd/server/main.go

EXPOSE 8080

CMD [ "/app" ]