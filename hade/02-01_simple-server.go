package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.String("port", "8000", "the listening port")
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, fmt.Sprintf("Hello, 世界!"))
	if err != nil {
		return
	}
}
func main() {
	// Register pattern-handler to http
	http.HandleFunc("/hello", HelloHandler)
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("Starting to listen on 8000 and server",
		"url", "url",
		"attempt", 3,
		"duration", time.Second,
	)
	logger.Info("url", zap.String("url", fmt.Sprintf("http://localhost:%s/hello", *port)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
