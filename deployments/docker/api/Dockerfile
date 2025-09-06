FROM golang:1.25-alpine3.22
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
CMD ["go", "run", "./cmd/api"]