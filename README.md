# Golang Docker 開発環境

Docker を使用した Golang 開発環境のセットアップです。

## 機能

- Go 1.21（Alpine Linux ベース）
- ホットリロード（air）
- 簡単な HTTP サーバーのサンプル

## 使い方

### 環境の起動

```bash
docker-compose up
```

### 環境の停止

```bash
docker-compose down
```

### アクセス方法

ブラウザで以下の URL にアクセスしてください：

```
http://localhost:8080
```

## ディレクトリ構造

```
.
├── .air.toml      # ホットリロード設定
├── Dockerfile     # Dockerイメージ定義
├── README.md      # このファイル
├── docker-compose.yml # Docker Compose設定
├── go.mod         # Goモジュール定義
└── main.go        # サンプルアプリケーション
```
