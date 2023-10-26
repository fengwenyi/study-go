package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve failed, err: %v\n", err)
		return
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {

	var path = "/Users/fengwenyi/Projects/go-projects/src/fengwenyi.com/study-go/05-http/hello.txt"

	b, _ := os.ReadFile(path)
	_, _ = fmt.Fprintln(w, string(b))

}
