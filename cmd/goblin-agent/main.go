// The main program for the agent

package main

import (
	"sync"

	"github.com/johnathanclong/Goofy-Goblin/pkg/agent"
	"github.com/johnathanclong/Goofy-Goblin/pkg/core"
	"github.com/johnathanclong/Goofy-Goblin/pkg/utils"
)

var mux sync.Mutex

func init() {
	utils.Status("success", "Initializing agent")
}

func _init(f core.Function) {
	mux.Lock()
	defer mux.Unlock()
	core.FunctionList = append(core.FunctionList, f)
}

func main() {
	agent := agent.New()
	go core.EventLoop()
	var b []string
	core.EventEmit("a", b)
	for {

	}
}
