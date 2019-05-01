// The main program for running the master server

package main

import (
	"github.com/johnathanclong/Goofy-Goblin/pkg/master/generator"
	"github.com/johnathanclong/Goofy-Goblin/pkg/master/modules"
)

func main() {
	generator.GenerateAgent(true, true, false, []string{"exploit/linux/pty_shell_client"})
	modules.ParseInfo("exploit/linux/pty_shell_client")
}
