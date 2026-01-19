FROM golang:1.21-alpine
WORKDIR /app
# Forzamos a Go a que use módulos
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
