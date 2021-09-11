package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!\n"))
		writer.Write([]byte(Fib(100000)))
	})
	http.ListenAndServe(":8099", nil)
}

func Fib(n int) string {
	res := ""
	for a, b := 0, 1; b < n; a, b = b, a+b {
		res += fmt.Sprintf("%d, ", b)
	}
	return res
}
