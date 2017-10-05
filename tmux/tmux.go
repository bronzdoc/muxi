package tmux

import (
	"github.com/bronzdoc/muxi/command"
)

type AsTmuxObject interface {
	SetSessionName(string)
	GetSessionName() string
}

type tmuxObject struct {
	sessionName string
	tmuxCommand command.TmuxCommand
}

func (t *tmuxObject) SetSessionName(sessionName string) {
	t.sessionName = sessionName
}

func (t *tmuxObject) SessionName() string {
	return t.sessionName
}
