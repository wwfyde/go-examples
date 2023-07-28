package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

// path "/" handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Parse Request header
	header := r.Header
	jsonData, _ := json.MarshalIndent(header, "", "    ")
	fmt.Printf("%s\n", jsonData)
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, "Hello, 世界!\n"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	log.Println(flag.Args())
	if len(flag.Args()) >= 1 {

	}
	log.Println("Starting...")
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
