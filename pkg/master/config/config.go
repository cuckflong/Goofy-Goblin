// This package stores the configuration of the agent

package config

// DebugString mode
var debugString = "false"

//Debug mode
var Debug = false

// VerboseString mode
var verboseString = "false"

// Verbose mode
var Verbose = false

// silentString mode
var silentString = "false"

// Silent mode
var Silent = false

func init() {
	if debugString == "true" {
		Debug = true
	}
	if verboseString == "true" {
		Verbose = true
	}
	if silentString == "true" {
		Silent = true
	}
}
