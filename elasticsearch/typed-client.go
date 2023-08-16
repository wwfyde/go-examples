package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"log"
	"os"
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
	es, err := elasticsearch.NewTypedClient(cfg)

	//// Initialize a client with the default settings.
	//es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// check if client to es successfully
	log.Printf("Client: %s", elasticsearch.Version)

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
	// Creating an index
	es.Indices.Create("my_index").Do(context.TODO())
	// Indexing document
	es.Index("my_index").
		Id("1").
		Request(document).
		Do(context.TODO())

	es.Get("my_index", "id").Do(context.TODO())
	// Searching Document
	res, _ := es.Search().
		Index("my_index").
		Request(&search.Request{
			Query: &types.Query{MatchAll: &types.MatchAllQuery{}},
		}).
		Do(context.TODO())
	// Print Searched document
	fmt.Printf("%s\n", res.Hits.Hits[0].Source_)
}
