// Provide basic functions

package utils

import (
	"github.com/fatih/color"
	"github.com/johnathanclong/Goofy-Goblin/pkg/master/config"
)

// Constants
const (
	Info    = 1 // Info information message
	Success = 2 // Success success message
	Error   = 3 // Error error message
	Debug   = 4 // Debug debug message
	Verbose = 5 // Verbose verbose message
)

// Status displays messages with colours according to the status
func Status(status int, message string) {
	if config.Silent {
		return
	}
	switch status {
	case Info:
		color.Cyan("[i] " + message)
	case Success:
		color.Green("[+] " + message)
	case Error:
		color.Red("[!] " + message)
	case Debug:
		color.Red("[Debug] " + message)
	case Verbose:
		color.Yellow("[-] " + message)
	default:
		color.Red("[Unknown] " + message)
	}
}
