package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

const defaultPort = "8102"
const bar = "-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-"

func dropErr(b []byte, err error) (string) {
	return string(b)
}


func factory() func(http.ResponseWriter, *http.Request) {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	return func(w http.ResponseWriter, r *http.Request) {
		s := bar + "\n"
		s += time.Now().In(loc).Format("2006-01-02 15:04:05 MST") + "\n"
		s += bar + "\n"
		s += dropErr(httputil.DumpRequest(r, true)) + "\n" + strings.Replace(bar, "+", "-", -1) + "\n"
		fmt.Print(s)
		fmt.Fprint(w, s)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = defaultPort }

	http.HandleFunc("/", factory())
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
