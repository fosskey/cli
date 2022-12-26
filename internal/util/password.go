package util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func Password(prompt string) string {

	// Save the current state of the terminal to restore later
	termState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	// Restore terminal state before "normal-exit" from this function
	defer func() {
		term.Restore(int(os.Stdin.Fd()), termState)
		fmt.Print("[···]\n") // cleanup the line
	}()

	// Restore terminal state before ^C "forced-exit" from this function
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		term.Restore(int(os.Stdin.Fd()), termState)
		fmt.Print("\n") // cleanup the line
		os.Exit(1)
	}()

	// Print the prompt message
	fmt.Print(prompt)

	pw := []byte{}

	// Keep reading password until non-empty
	for len(pw) == 0 {
		pw, err = term.ReadPassword(int(os.Stderr.Fd()))
		if err != nil {
			panic(err)
		}
	}

	// Convert to string
	return string(pw)
}
