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

	scans, err := client.GetAllScans()
	if err != nil {
		log.Fatal(err)
	}

	for _, scan := range scans {
		log.Println(scan.Name)
	}

}
