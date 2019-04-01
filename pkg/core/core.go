// This package contains the components for the agent's core

package core

// Constants
const (
	coreStartUp  = 1 // coreStartup Function to run on start up
	corePeriod   = 2 // corePeriod Function to run on period
	coreCall     = 3 // coreCall Function to run on call
	coreInterupt = 4 // coreInterupt Function to run in signal interupt
)

// Function contains all the information required for the core to work with
type Function struct {
	Name        string
	Description string
	Period      int
	Mode        int
	Func        func([]string)
	ParamsInfo  []Parameter
	Parameters  []string
}

// Call will call the given function with the parameter array
func (f Function) Call() {
	f.Func(f.Parameters)
}

// Parameter contains the information of a parameter
type Parameter struct {
	Name        string
	Type        string
	Description string
}
