// The main program for running the master server

package main

import (
	"github.com/johnathanclong/Goofy-Goblin/pkg/generator"
)

func main() {
	generator.GenerateAgent(true, true, true, []string{"exploit/linux/pty-shell-client.go"})
}
