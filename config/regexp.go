package config

import "regexp"

var (
	// used to remove comments from json
	regexComments = regexp.MustCompile(`\/\/([\w\s\'].*)`)

	// used to remove special chars from paths
	SafeVarName = regexp.MustCompile(`[^a-zA-Z0-9]`)
)
