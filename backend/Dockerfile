FROM golang:latest

WORKDIR '/app'

COPY go.mod go.sum main.go ./

RUN react-go-server mod download

COPY . .

EXPOSE 8000

ENTRYPOINT ["go", "run", "main.go"]
