// The main program for running the master server

package main

import (
	"github.com/johnathanclong/Goofy-Goblin/pkg/generator"
	"github.com/johnathanclong/Goofy-Goblin/pkg/master"
)

func main() {
	generator.GenerateAgent(true, true, false, []string{"exploit/linux/pty_shell_client"})
	master.ParseInfo("exploit/linux/pty_shell_client")
}
