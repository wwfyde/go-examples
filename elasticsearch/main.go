package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)
	var (
	//r map[string]any
	//wg sync.WaitGroup
	)

	// Create a Config
	cert, _ := os.ReadFile("elasticsearch/http_ca.crt")
	fmt.Printf("cert: %s\n", cert)
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "EmVusjGroBAE9QtfJoNF",
		CACert:   cert,
	}

	// Client to es with cfg
	es, err := elasticsearch.NewClient(cfg)

	//// Initialize a client with the default settings.
	//es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// check if client to es successfully
	log.Printf("Client: %s", elasticsearch.Version)
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	defer res.Body.Close()
	log.Println(res.String())

	// Indexing a document
	document := struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}{
		Id:    1,
		Name:  "Foo",
		Price: 10,
	}
	data, err := json.Marshal(document)
	// Creating an index
	es.Indices.Create("my_index")
	res, err = es.Index("my_index", bytes.NewReader(data))

	es.Get("my_index", "id")
	// Searching Document
	query := `{ query: {match_all: {} } }`
	res, _ = es.Search(
		es.Search.WithIndex("my_index"),
		es.Search.WithBody(strings.NewReader(query)),
	)
	fmt.Println(res.Status(), res.Header, res.Body)
}
