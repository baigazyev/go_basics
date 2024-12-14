
FROM golang:1.19


ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o main ./cmd/app/main.go


EXPOSE 8080


CMD ["./main"]
