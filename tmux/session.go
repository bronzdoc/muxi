package tmux

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type Session struct {
	name    string
	windows []*Window
	command []map[string]interface{}
}

// New Session
func NewSession(name string) *Session {
	newName := name
	const RANDSOURCE = 100000

	// Generate tmux session name if not given
	if newName == "" {
		randInt := rand.Intn(RANDSOURCE)
		newName = fmt.Sprintf("%d", randInt)
	}

	return &Session{
		name: newName,
		command: []map[string]interface{}{
			{
				"cmd":  BASECOMMAND,
				"args": []string{"rename-session", newName},
			},
			{
				"cmd":  BASECOMMAND,
				"args": []string{"switch-client", "-t", newName},
			},
		},
	}
}

// Adds windows to the session
func (s *Session) AddWindow(window *Window) {
	window.Setup(s.name)
	s.windows = append(s.windows, window)
}

// Get session windows
func (s *Session) Windows() []*Window {
	return s.windows
}

// Get session name
func (s *Session) Name() string {
	return s.name
}

// Creates a new tmux session
func (s *Session) Create() {
	for _, c := range s.command {
		cmd := exec.Command(c["cmd"].(string), c["args"].([]string)...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println(err)
		}
	}

	// Create session windows
	for _, w := range s.windows {
		w.Create()
	}
}
