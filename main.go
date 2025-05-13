package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println(time.Now().Format(time.DateTime), r.URL.Path, " - ", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("<b>hello world!</b>"))
		if err != nil {
			println("writer err:", err.Error())
		}
	}))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
