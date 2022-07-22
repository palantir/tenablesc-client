package main

import (
	"fmt"
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

	var analysisResult []tenablesc.VulnSumIPResult

	_, err := client.Analyze(&tenablesc.Analysis{
		Type: "vuln",
		Query: tenablesc.AnalysisQuery{
			Type:       "vuln",
			SourceType: "cumulative",
			Tool:       "sumip",
			Filters: []tenablesc.AnalysisFilter{
				{
					FilterName: "repository",
					Operator:   "=",
					Value: []map[string]string{
						{
							"id": "1",
						},
					},
				},
			},
		},
		SourceType:    "cumulative",
		SortField:     "score",
		SortDirection: "desc",
	},
		&analysisResult,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+#v", analysisResult)

}
