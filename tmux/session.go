package tmux

import (
	"fmt"
	"math/rand"

	"github.com/bronzdoc/muxi/command"
)

// Represents a tmux session
type Session struct {
	tmuxObject
	windows []*Window
}

// Creates a new Session
func NewSession(name string) *Session {
	newName := name
	const RANDSOURCE = 100000

	// Generate tmux session name if not given
	if newName == "" {
		randInt := rand.Intn(RANDSOURCE)
		newName = fmt.Sprintf("%d", randInt)
	}

	s := Session{
		tmuxObject: tmuxObject{
			sessionName: newName,
		},
	}

	s.SetTmuxCommand(
		command.NewSessionCommand(newName),
	)

	s.tmuxCommand.AddPostHook(s.createWindows)

	return &s
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

// Creates a new tmux session and its windows
func (s *Session) Create() {
	s.tmuxCommand.Execute()
}

func (s *Session) createWindows() {
	for _, w := range s.windows {
		w.Create()
	}
}
