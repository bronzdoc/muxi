package tmux

import (
	"fmt"
	"log"
	"math/rand"
	"time"

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

	// Generate tmux session name if not given
	if newName == "" {
		rand.Seed(int64(time.Now().Nanosecond()))
		newName = fmt.Sprintf("%d", rand.Int())
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
	if err := s.tmuxCommand.Execute(); err != nil {
		log.Fatalf("Couldn't create session '%s' %v", s.sessionName, err)
	}
}

func (s *Session) createWindows() {
	for _, w := range s.windows {
		w.Create()
	}
}
