// Agent will be the malware itself, implement basic structure and functions here

package agent

import (
	"os/user"
	"runtime"
	"time"

	"github.com/johnathanclong/Goofy-Goblin/pkg/utils"
	uuid "github.com/satori/go.uuid"
)

// Agent contains all the information the agent
type Agent struct {
	ID             uuid.UUID
	Platform       string
	Architecture   string
	UserName       string
	UserUID        string
	UserGID        string
	Pid            int
	InitialCheckIn time.Time
	LastCheckIn    time.Time
	IPs            []string
	Debug          bool
	Verbose        bool
}

// New creates a new agent
func New(debug bool, verbose bool) Agent {
	a := Agent{}
	a.ID = uuid.NewV4()
	a.Platform = runtime.GOOS
	a.Architecture = runtime.GOARCH
	a.Debug = debug

	u, err := user.Current()
	if err != nil {
		if a.Debug {
			utils.Status("error", "Failed getting current user")
		}
	} else {
		a.UserName = u.Username
		a.UserUID = u.Uid
		a.UserGID = u.Gid
	}
	return a
}

// CheckIn tries to check in with the master server
func CheckIn(a Agent) {
	// Todo make sure there is a connection with the master server
	a.LastCheckIn = time.Now()
}

// InitCheckIn do the initial checkin with the master server
func InitCheckIn(a Agent) {
	// Todo initial checkin with the master server
	a.InitialCheckIn = time.Now()
}
