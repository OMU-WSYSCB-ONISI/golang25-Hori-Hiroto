package main

import (
 "fmt"
 "math/rand"
 "net/http"
 "time"
)

func main() {
 fmt.Println("Week 03 課題（Web おみくじ）")

 http.HandleFunc("/webfortune", fortuneHandler)

 port := "8082"
 fmt.Println("Server running at http://localhost:" + port)

 // ListenAndServe のエラーをチェック
 if err := http.ListenAndServe(":"+port, nil); err != nil {
  fmt.Println("サーバー起動エラー:", err)
 }
}

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
 fortunes := []string{"大吉", "中吉", "吉", "末吉", "凶"}

 rand.Seed(time.Now().UnixNano())
 result := fortunes[rand.Intn(len(fortunes))]

 msg := fmt.Sprintf("今の運勢は %s です！", result)

 // ターミナル表示
 fmt.Println("アクセスされました →", msg)

 w.Header().Set("Content-Type", "text/plain; charset=utf-8")
 fmt.Fprintln(w, msg)
}
