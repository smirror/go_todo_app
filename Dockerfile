# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.21.3-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# ---------------------------------------------------
# deploy
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ---------------------------------------------------
# local dev
FROM golang:1.21.3 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]