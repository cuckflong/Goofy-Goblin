// Agent will be the malware itself, implement basic structure and functions here

package agent

import (
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID             uuid.UUID
	Platform       string
	Architecture   string
	InitialCheckIn time.Time
	LastCheckIn    time.Time
	IPs            []string
}

func New()
