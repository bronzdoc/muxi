package tmux

import (
	"github.com/bronzdoc/muxi/command"
)

type tmuxObject struct {
	sessionName string
	tmuxCommand command.BaseCommand
}

func (t *tmuxObject) SetSessionName(sessionName string) {
	t.sessionName = sessionName
}

func (t *tmuxObject) SessionName() string {
	return t.sessionName
}

func (t *tmuxObject) SetTmuxCommand(cmd command.BaseCommand) {
	t.tmuxCommand = cmd
}
