// This package contains the components for the agent's core

package core

// FunctionList is an array of function loaded into the agent
var FunctionList []Function

// EventChannel is the channel for events
var EventChannel = make(chan Event, 500)

// Constants
const (
	coreStartUp  = 1 // coreStartup Function to run on start up
	corePeriod   = 2 // corePeriod Function to run on period
	coreCall     = 3 // coreCall Function to run on call
	coreInterupt = 4 // coreInterupt Function to run in signal interupt
)

// Function contains all the information required for the core to work with
type Function struct {
	Code   string
	Period int
	Mode   int
	Active bool
	Func   func([]string)
}

// Event is the
type Event struct {
	Code       string
	Parameters []string
}

// Call will call the given function with the parameter array
func Call(event Event) {
	for _, f := range FunctionList {
		if f.Code == event.Code {
			go f.Func(event.Parameters)
			return
		}
	}
}

// EventLoop wait for events and call them one by oen
func EventLoop() {
	for {
		e := <-EventChannel
		Call(e)
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
