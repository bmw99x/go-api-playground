package helpers

import (
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffcontext"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
	"time"
)

func GetFeatureFlagValue(key string) bool {
	// Load in feature flags from file
	// poll for changes every 3 seconds
	err := ffclient.Init(ffclient.Config{
		PollingInterval: 3 * time.Second,
		Retriever: &fileretriever.Retriever{
			Path: "flags.yaml",
		},
	})
	if err != nil {
		panic(err)
	}
	// If we ever wanted to have variations, we could allow
	// this function to pass in some unique ID and use that
	// to fetch the user value
	user := ffcontext.NewEvaluationContext("user-key")
	result, err := ffclient.BoolVariation("database", user, false)
	if err != nil {
		panic(err)
	}
	println("Feature flag value is", result)
	defer ffclient.Close()
	return result
}
