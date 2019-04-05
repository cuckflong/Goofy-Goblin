// Provide basic functions

package utils

import "github.com/fatih/color"

// Status displays messages with colours according to the status
func Status(status string, message string) {
	switch status {
	case "info":
		color.Cyan("[i] " + message)
	case "success":
		color.Green("[+] " + message)
	case "error":
		color.Red("[!] " + message)
	default:
		color.Red("[Unknown] " + message)
	}
}
