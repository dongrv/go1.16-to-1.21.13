package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"test-go1.21.13/generic"
	"time"
)

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println(time.Now().Format(time.DateTime), r.URL.Path, " - ", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", middleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("<b>hello world!</b>"))
		if err != nil {
			println("writer err:", err.Error())
		}
	}))
	// 注册pprof路由
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

	// 测试pgo优化编译
	// curl -o default.pgo http://localhost:8080/debug/pprof/profile?seconds=30s
	// go build -pgo=auto -o main.exe .
}

func main() {
	generic.Search[int]([]int{1, 2, 3}, 1)
	generic.Search[float32]([]float32{1, 2, 3}, 1)

	fmt.Printf("%v", generic.Min[int](1, 2, 3))
	fmt.Printf("%v", generic.Min[float64](1, 2, 3))

	// 分析泛型编译后的汇编语句:
	// go build -a -o main.exe .
	// go tool objdump -S .\main.exe
}
