// +build test_module

package main

import (
	"fmt"

	"github.com/johnathanclong/Goofy-Goblin/pkg/agent"
	"github.com/johnathanclong/Goofy-Goblin/pkg/core"
)

func init() {
	_init(core.Function{
		Code:   "TEST_MODULE",
		Period: 0,
		Mode:   core.CoreCall,
		Active: true,
		Func:   TrashTalk,
	})
}

func TrashTalk(a agent.Agent, parameters []string) {
	fmt.Println("Agent %s is shit", a.Username)
	fmt.Println("Parameter %s is useless", parameters[0])
}
