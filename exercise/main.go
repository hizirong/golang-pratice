package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World")) //w 只接收byte array 所以要先轉型
}

func main() {
	http.HandleFunc("/", handler) //指派上述寫好的func，負責處理所有請求
	http.ListenAndServe(":8080", nil)
}
