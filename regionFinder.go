package main

import (
	"io"
	"log"
	"net/http"
	"regexp"
)

// This code is borrowed from https://github.com/GoogleCloudPlatform/cloud-run-hello/blob/master/hello.go#L118

func getRegion() string {
	client := &http.Client{}

	// Get region from metadata server
	region := ""
	req, _ := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/region", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			responseBody, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			region = regexp.MustCompile(`projects/[^/]*/regions/`).ReplaceAllString(string(responseBody), "")
		}
	}
	if region == "" {
		// Fallback: get "zone" from metadata server (running on VM e.g. Cloud Run for Anthos)
		req, _ = http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/zone", nil)
		req.Header.Set("Metadata-Flavor", "Google")
		res, err = client.Do(req)
		if err == nil {
			defer res.Body.Close()
			if res.StatusCode == 200 {
				responseBody, err := io.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}
				region = regexp.MustCompile(`projects/[^/]*/zones/`).ReplaceAllString(string(responseBody), "")
			}
		}
	}

	return region
}
