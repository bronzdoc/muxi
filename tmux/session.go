package tmux

import (
	"fmt"
	"math/rand"

	"github.com/bronzdoc/muxi/tmux/command"
)

type Session struct {
	tmuxObject
	windows []*Window
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
		tmuxObject: tmuxObject{
			tmuxCommand: command.NewSessionCommand(newName),
			sessionName: newName,
		},
	}
}

// Adds windows to the session
func (s *Session) AddWindow(window *Window) {
	window.SetSessionName(s.Name())
	s.windows = append(s.windows, window)
}

// Get session windows
func (s *Session) Windows() []*Window {
	return s.windows
}

// Get session name
func (s *Session) Name() string {
	return s.SessionName()
}

// Creates a new tmux session
func (s *Session) Create() {
	s.tmuxCommand.Execute()

	// Create session windows
	for _, w := range s.windows {
		w.Create()
	}
}
