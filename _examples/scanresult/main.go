package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/palantir/tenablesc-client/tenablesc"
)

func main() {
	client := tenablesc.NewClient(
		os.Getenv("TENABLE_URL"), // Tenable SC host. ensure the url has /rest in their path.
	).SetAPIKey(
		os.Getenv("TENABLE_ACCESS_KEY"), // Tenable SC access key.
		os.Getenv("TENABLE_SECRET_KEY"), // Tenable SC secret key.
	)

	now := time.Now()
	dayBefore := now.Add(-1 * 24 * time.Hour)
	scans, err := client.GetAllScanResultsByTime(dayBefore, now)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listing scan results for the past 24 hours")
	for _, scan := range scans {
		log.Println(scan.Name, scan.ID)
	}

	firstScanResult := scans[0]
	log.Printf("get scan result for ID: %s\n", firstScanResult.ID)
	result, err := client.GetScanResult(string(firstScanResult.ID))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("scan result name returned:", result.Name)

	log.Printf("downloading scan results for ID: %s\n", firstScanResult.ID)
	outFileName := fmt.Sprintf("scan-results-%s.nessus", firstScanResult.ID)
	downloadedData, err := client.DownloadScanResult(string(firstScanResult.ID))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("writing downloaded data to file: %s", outFileName)
	err = os.WriteFile(outFileName, downloadedData, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}

}
