FROM golang:1.16-alpine
RUN apk add --no-cache git

WORKDIR /app/sysmon-server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/sysmon-server .

EXPOSE 6060
CMD ["./out/sysmon-server"]
