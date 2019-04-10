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
	fmt.Printf("Agent %s is shit\n", a.UserName)
	fmt.Printf("Parameter %s is useless\n", parameters[0])
}
