package arr

import (
	"fmt"
	"net/http"
	"time"
)

func GoRoutinesArr(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 1000000; i++ {
		go TestArr()
	}
	_, err := w.Write([]byte("WORKED"))
	if err != nil {
		fmt.Println(err)
	}
}
func TestArr() {
	c := make([]string, 5000)
	_ = len(c)
	time.Sleep(3 * time.Second)
}
