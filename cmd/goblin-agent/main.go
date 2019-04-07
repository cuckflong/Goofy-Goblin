// The main program for the agent

package main

import (
	"sync"
	"time"

	"github.com/johnathanclong/Goofy-Goblin/pkg/config"

	"github.com/johnathanclong/Goofy-Goblin/pkg/agent"
	"github.com/johnathanclong/Goofy-Goblin/pkg/core"
	"github.com/johnathanclong/Goofy-Goblin/pkg/utils"
)

var mux sync.Mutex
var a agent.Agent

func init() {
	utils.Status(utils.Success, "Initializing agent")
}

func _init(f core.Function) {
	mux.Lock()
	defer mux.Unlock()
	core.FunctionList = append(core.FunctionList, f)
}

func main() {
	if config.Debug {
		utils.Status(utils.Info, "Starting agent in debug mode")
	}
	if config.Verbose {
		utils.Status(utils.Info, "Starting agent in verbose mode")
	}
	a = agent.New()

	go core.EventLoop(a)

	for {
		agent.Heartbeat(a)
		time.Sleep(60 * time.Second)
	}
}
