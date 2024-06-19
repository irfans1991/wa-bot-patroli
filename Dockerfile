FROM golang:alpine

# RUN apk update && apk upgrade 

# RUN apk add --no-cache gcc g++ sqlite

RUN apk add --no-cache build-base sqlite

ENV CGO_ENABLED=1

WORKDIR /opt/patroli-bot

RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "main.go"]
