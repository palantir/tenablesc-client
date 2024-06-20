package main

import (
	"log"
	"os"

	"github.com/palantir/tenablesc-client/tenablesc"
)

func main() {
	client := tenablesc.NewClient(
		os.Getenv("TENABLE_URL"), // Tenable SC host. ensure the url has /rest in their path.
	).SetAPIKey(
		os.Getenv("TENABLE_ACCESS_KEY"), // Tenable SC access key.
		os.Getenv("TENABLE_SECRET_KEY"), // Tenable SC secret key.
	)

	queries, err := client.GetAllQueries()
	if err != nil {
		log.Fatal(err)
	}

	for i := range queries {
		log.Println(queries[i].Name)
	}

	query, err := client.GetQuery("0")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(query.Name)

}
