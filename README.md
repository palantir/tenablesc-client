<p align=right>
<a href="https://pkg.go.dev/github.com/palantir/tenablesc-client/tenable"><img src="https://pkg.go.dev/badge/github.com/palantir/tenablesc-client/tenable.svg" alt="Go Reference"></a>
<a href="https://autorelease.general.dmz.palantir.tech/palantir/tenablesc-client"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt=Autorelease></a>
</p>

# Tenable.SC Client

## Overview

This is a golang client for interacting with the Tenable.SC API. 

Use cases include automating asset creation, metric gathering, and general configuration management.

Not all endpoints are implemented, pull requests are welcome!

## Usage Example

```go
package main

import (
	"fmt"
	"os"
	"github.palantir.build/arch/tenablesc-client/tenablesc"
)

func main() {

	client := tenablesc.NewDefaultAPIKeyClient(
		// SC_URL should be the full URL to the API base;
		// Typically this is https://FQDN/rest
		os.Getenv("SC_URL"), 
		// Access and Secret keys are generated from the Users
		// UI in Tenable.SC.
		os.Getenv("SC_ACCESS_KEY"),
		os.Getenv("SC_SECRET_KEY"),
	)

	_, err := client.GetCurrentUser()
	if err != nil {
		fmt.Errorf("unable to authenticate to SC: %w", err)
		os.Exit(1)
	}

	var analysisResult []tenablesc.VulnSumIPResult

	// Composing the query structs is a combination of reading the docs
	// and using browser Developer Tools to identify the right fields by
	// building the queries in the UI. 
	err := client.Analyze(&tenablesc.Analysis{
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
							// if this weren't an example, I'd recommend looking up your
							// repo ID first. your accessible repos may vary.
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
		fmt.Errorf("couldn't get list of vulnerabilities: %w", err)
		os.Exit(1)
    }
	
	fmt.Printf("%+$v", analysisResult)
	
}

```





### References

- [Tenable.SC Vendor Product page](https://www.tenable.com/products/tenable-sc)
- [Tenable.SC API docs](https://docs.tenable.com/tenablesc/api/index.htm)
