package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Week 04: ここに課題のコードを記述してください
	fmt.Println("Week 04 課題")

	// 追加：Webサーバ起動
	startWebServer()
}

// Webサーバ起動関数
func startWebServer() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("15:04")
		userAgent := r.Header.Get("User-Agent")
		msg := fmt.Sprintf("今の時刻は %s で，利用しているブラウザは %s ですね", now, userAgent)
		fmt.Fprintln(w, msg)
	})

	fmt.Println("http://localhost:8081/info にアクセスしてください")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("サーバ起動エラー:", err)
	}
}
