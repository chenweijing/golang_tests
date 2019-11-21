package main

import (
	"fmt"
	"log"
	"net/http"
)

// func SayHello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello,this is a test!!!"))
// }
// func main() {
// 	// http.HandleFunc("/sayhello", SayHello)
// 	http.Handle("/", http.FileServer(http.Dir("./static")))
// 	log.Fatal(http.ListenAndServe(":9001", nil))
// }

func main() {
	fmt.Println("Now Listening on 80")
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}
