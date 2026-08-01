FROM golang:1.26.5

WORKDIR /app

RUN apt-get update && apt-get install -y make

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest

ENV PATH="/root/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["make", "run"]