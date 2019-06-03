// Agent will be the malware itself, implement basic structure and functions here

package agent

import (
	"fmt"
	"os/user"
	"runtime"
	"sync"
	"time"

	"github.com/johnathanclong/Goofy-Goblin/pkg/agent/config"

	"github.com/johnathanclong/Goofy-Goblin/pkg/agent/utils"
	uuid "github.com/satori/go.uuid"
)

var mux sync.Mutex

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
}

// New creates a new agent
func New() Agent {
	a := Agent{}
	a.ID = uuid.NewV4()
	a.Platform = runtime.GOOS
	a.Architecture = runtime.GOARCH

	u, err := user.Current()
	if err != nil {
		if config.Debug {
			utils.Status(utils.Error, "Failed getting current user")
		}
	} else {
		a.UserName = u.Username
		a.UserUID = u.Uid
		a.UserGID = u.Gid
	}
	InitCheckIn(a)
	utils.Status(utils.Info, fmt.Sprintf("Username: %s", a.UserName))
	utils.Status(utils.Info, fmt.Sprintf("Platform: %s", a.Platform))
	utils.Status(utils.Info, fmt.Sprintf("Architecture: %s", a.Architecture))
	utils.Status(utils.Info, fmt.Sprintf("User UID: %s", a.UserUID))
	utils.Status(utils.Info, fmt.Sprintf("User GID: %s", a.UserGID))
	return a
}

// Heartbeat tries to check in with the master server
func Heartbeat(a Agent) {
	// Todo make sure there is a connection with the master server
	var now = time.Now()
	if config.Verbose {
		utils.Status(utils.Verbose, fmt.Sprintf("Sending heartbeat at %s", now.String()))
	}
	mux.Lock()
	defer mux.Unlock()
	a.LastCheckIn = now
}

// InitCheckIn do the initial checkin with the master server
func InitCheckIn(a Agent) {
	// Todo initial checkin with the master server
	var now = time.Now()
	if config.Verbose {
		utils.Status(utils.Verbose, fmt.Sprintf("Initial checkin at %s", now.String()))
	}
	mux.Lock()
	defer mux.Unlock()
	a.InitialCheckIn = time.Now()
}
