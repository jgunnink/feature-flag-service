package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type FeatureFlags struct {
	MyFlag bool `json:"my_flag"`
}

// FeatureFlags is a function that returns a JSON payload with fields from the FeatureFlags struct
func featureFlagsHandler(w http.ResponseWriter, r *http.Request) {
	region := getRegion()

	flags := FeatureFlags{}

	if strings.HasPrefix(region, "us") {
		flags.MyFlag = true
	}

	if strings.HasPrefix(region, "eu") {
		flags.MyFlag = false
	}

	if strings.HasPrefix(region, "us") {
		flags.MyFlag = false
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(flags)
}
