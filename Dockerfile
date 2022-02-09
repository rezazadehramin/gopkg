FROM golang:alpine

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /srv

COPY go.* .

RUN go mod download

COPY . .

CMD [ "go", "test", "-v", "./..." ]