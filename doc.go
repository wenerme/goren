package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	release  = "v1.0.0"
	gitsha   = "no gitsha provided"
	compiled = "0"
	version  = ""
)

// getVersion returns the prog version
func getVersion() string {
	if version == "" {
		tm, err := strconv.ParseInt(compiled, 10, 64)
		if err != nil {
			return "unable to parse compiled time"
		}
		version = fmt.Sprintf("%s (git+sha: %s, built: %s)", release, gitsha, time.Unix(tm, 0).Format("02-01-2006"))
	}

	return version
}

type GlobalConfig struct {
}
