// This package contains the components for the agent's core

package core

// FunctionList is an array of function loaded into the agent
var FunctionList []Function

// Constants
const (
	coreStartUp  = 1 // coreStartup Function to run on start up
	corePeriod   = 2 // corePeriod Function to run on period
	coreCall     = 3 // coreCall Function to run on call
	coreInterupt = 4 // coreInterupt Function to run in signal interupt
)

// Function contains all the information required for the core to work with
type Function struct {
	Name       string
	Code       string
	Period     int
	Mode       int
	Func       func([]string)
	Parameters []string
}

// Event is the
type Event struct {
	Code       string
	Parameters []string
}

// Call will call the given function with the parameter array
func (f Function) Call() {
	f.Func(f.Parameters)
}
