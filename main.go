package main

import (
	"fmt"
	"go-bench/arr"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"time"
)

var i = 0
var c chan int

func main() {
	r := http.NewServeMux()
	c = make(chan int)

	r.HandleFunc("/test2", goRoutines)
	r.HandleFunc("/test", arr.GoRoutinesArr)
	r.HandleFunc("/force-collect", forceCollect)
	r.HandleFunc("/cpu", cpuLoad)
	r.HandleFunc("/upload-file", fileUpload)
	r.HandleFunc("/download-file", fileDownload)

	fmt.Println("Listening on :80")
	go func() {
		count := 0
		for {
			<-c
			count++
			fmt.Println(count)
			if count == 1000000 {
				fmt.Println(time.Since(t))
			}
		}
	}()
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err)
	}
}

func fileDownload(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("E:\\Backup\\C9\\Client-patch.zip")
	if err != nil {
		log.Fatal(err)
	}
	line := make([]byte, 1_000_000)
	for {
		n, err := f.Read(line)
		if err == nil && n > 0 {
			_, err := w.Write(line)
			if err != nil {
				fmt.Println("Can't write anymore at fileDownload:", err.Error())
				return
			}
			println("SENT LINE")
			//time.Sleep(5 * time.Second)
		}
		if err != nil {
			if err == io.EOF {
				println("SENT FILE")
				return
			}
			log.Fatal(err)
		}
	}
}

func fileUpload(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte(r.Header.Get("Content-Type")))
	//return
	for {
		line := make([]byte, 1_000_000_000)
		n, err := r.Body.Read(line)
		if n > 0 {
			if _, err := w.Write(line); err != nil {
				c := make(chan int)
				<-c
				log.Fatal(err)
			}
		}
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
func forceCollect(w http.ResponseWriter, r *http.Request) {
	debug.FreeOSMemory()
}

func cpuLoad(w http.ResponseWriter, r *http.Request) {
	i++
	ans := 15
	for j := 6; j < 160000000; j++ {
		ans += ans*((j*(5-j)/j*2)*j) - (j / (5 + j) * (j * 14))
	}
	_, err := w.Write([]byte(strconv.Itoa(ans)))
	if err != nil {
		fmt.Println(err)
	}
}
