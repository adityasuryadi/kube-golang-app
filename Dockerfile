FROM --platform=linux/amd64 golang:alpine

RUN apk update && apk add --no-cache git \
curl

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]
