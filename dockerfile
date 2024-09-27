FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
ENV PORT=8080
EXPOSE 8080
CMD ["go","run", "main.go"]