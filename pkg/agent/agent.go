// Agent will be the malware itself, implement basic structure and functions here

package agent

import (
	"os/user"
	"runtime"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Agent struct {
	ID             uuid.UUID
	Platform       string
	Architecture   string
	UserName       string
	UserUid        string
	UserGid        string
	Pid            int
	InitialCheckIn time.Time
	LastCheckIn    time.Time
	IPs            []string
	Debug          bool
	Verbose        bool
}

func New(debug bool, verbose bool) Agent {
	a := Agent{}
	a.ID, _ = uuid.NewV4()
	a.Platform = runtime.GOOS
	a.Architecture = runtime.GOARCH
	a.Debug = debug

	u, err := user.Current()
	if err != nil {
		if a.Debug {

		}
	} else {
		a.UserName = u.Username
		a.UserUid = u.Uid
		a.UserGid = u.Gid
	}
	return a
}
