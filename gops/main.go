package main

import (
	"github.com/google/gops/agent"
	"net/http"
)

func main() {

	err := agent.Listen(agent.Options{})
	_ = err
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err2 := w.Write([]byte(`hello, 世界!`))
		if err2 != nil {
			return
		}
	})
	_ = http.ListenAndServe(":8000", http.DefaultServeMux)
}
