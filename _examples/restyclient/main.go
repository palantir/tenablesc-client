package main

import (
	"crypto/tls"
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

	// Configuring the client's minimum TLS version to be TLS v1.2
	client.RestyClient().SetTLSClientConfig(&tls.Config{
		MinVersion: tls.VersionTLS12,
	})

	_, err := client.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
	}
}
