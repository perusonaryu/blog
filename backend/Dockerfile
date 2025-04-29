FROM golang:1.24-alpine

WORKDIR /backend

COPY . .

RUN go mod download

# ホットリロード用のツールをインストール
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]