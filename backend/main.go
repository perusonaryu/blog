package main

import (
	"blog/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

/**
 * メインアプリケーション
 *
 * 簡単なHTTPサーバーを起動し、ルートエンドポイントにアクセスするとメッセージを表示します
 * データベース接続を確認するエンドポイントも追加されています
 *
 * @version 1.0.0
 */
func main() {
	// データベース接続
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	defer db.Close()

	// ルートハンドラーの設定
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/db-check", handleDBCheck(db))

	// サーバー起動ポート
	port := ":8080"
	fmt.Printf("サーバーが%sポートで起動しました\n", port)

	// HTTPサーバーを起動
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}

/**
 * ルートパスのリクエストを処理するハンドラー関数
 *
 * @param {http.ResponseWriter} w - HTTPレスポンスを書き込むためのインターフェース
 * @param {*http.Request} r - HTTPリクエスト情報
 */
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Docker環境へようこそ！")
}

/**
 * データベース接続を確認するハンドラー関数
 *
 * @param {*sql.DB} db - データベース接続オブジェクト
 * @return {http.HandlerFunc} - HTTPハンドラー関数
 */
func handleDBCheck(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("データベース接続テスト")
		// データベース接続テスト
		if err := db.Ping(); err != nil {
			http.Error(w, fmt.Sprintf("データベース接続エラー: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "データベース接続成功！")
	}
}
