FROM golang:1.13

LABEL maintainer="Kareem Khaled <kareem.mohllal@gmail.com>"
LABEL version="1.0"
LABEL project="go-challenge"

WORKDIR /app

# exploit the Docker layer caching by installing dependencies 
# which will be cached if the go.mod and go.sum files are not changed
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main.go .

HEALTHCHECK CMD curl --fail http://localhost:5000/api/ping || exit 1

EXPOSE 5000

CMD ["./main.go"]