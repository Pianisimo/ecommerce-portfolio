FROM golang:alpine3.16

WORKDIR '/app'

COPY go.mod go.sum main.go ./

RUN go mod download

COPY . .

EXPOSE 8000

ENTRYPOINT ["go", "run", "main.go"]
