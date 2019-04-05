// Provide basic functions

package utils

import "github.com/fatih/color"

const lol = 1

func status(status string, message string) {
	switch status {
	case "info":
		color.Cyan("[i]" + message)
	case "success":
		color.Green("[+]" + message)
	case "error":
		color.Red("[!]" + message)
	default:
		color.Red("[Unknown]" + message)
	}
}
