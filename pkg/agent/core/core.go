// This package contains the components for the agent's core

package core

import (
	"github.com/johnathanclong/Goofy-Goblin/pkg/agent/agent"
)

// FunctionList is an array of function loaded into the agent
var FunctionList []Function

// EventChannel is the channel for events
var EventChannel = make(chan Event, 500)

// Constants
const (
	CoreStartUp  = 1 // CoreStartup Function to run on start up
	CorePeriod   = 2 // CorePeriod Function to run on period
	CoreCall     = 3 // CoreCall Function to run on call
	CoreInterupt = 4 // CoreInterupt Function to run in signal interupt
)

// Function contains all the information required for the core to work with
type Function struct {
	Code   string
	Period int
	Mode   int
	Active bool
	Func   func(agent.Agent, []string)
}

// Event contains the information require to perform an event
type Event struct {
	Code       string
	Parameters []string
}

// Call will call the given function with the parameter array
func Call(a agent.Agent, event Event) {
	for _, f := range FunctionList {
		if f.Code == event.Code {
			go f.Func(a, event.Parameters)
			return
		}
	}
}

// EventLoop wait for events and call them one by oen
func EventLoop(a agent.Agent) {
	for {
		e := <-EventChannel
		Call(a, e)
	}
}

// EventEmit emits a new event into the channel
func EventEmit(code string, parameters []string) {
	e := Event{
		Code:       code,
		Parameters: parameters,
	}
	EventChannel <- e
}
