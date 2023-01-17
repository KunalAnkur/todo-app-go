FROM golang:1.12.0-alpine3.9
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod downlaod
RUN go build -o main .
CMD ["/app/main"]