package main

import (
	"fmt"
	"net/http"
	"time"
)

var t time.Time

func goRoutines(w http.ResponseWriter, r *http.Request) {
	t = time.Now()
	for i := 0; i < 1000000; i++ {
		go test()
	}
	_, err := w.Write([]byte("WORKED"))
	if err != nil {
		fmt.Println(err)
	}
}
func test() {
	l := make([]string, 10)
	_ = l
	time.Sleep(3 * time.Second)
	c <- 2
}
