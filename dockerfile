FROM golang:1.26.5

WORKDIR /app

RUN apt-get update && apt-get install -y make

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["make", "run"]