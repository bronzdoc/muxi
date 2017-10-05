package tmux

import (
	"fmt"

	"github.com/bronzdoc/muxi/command"
)

var PANE_INDEX = 0

// Represents a tmux pane
type Pane struct {
	tmuxObject
	index    int
	commands []string
}

// Create a new Pane
func NewPane(root string) *Pane {
	return &Pane{
		tmuxObject: tmuxObject{
			tmuxCommand: command.NewPaneCommand(
				fmt.Sprintf("-c %s", root),
			),
		},
	}
}

// Adds a new command to execute in pane
func (p *Pane) AddCommand(cmd string) {
	p.commands = append(p.commands, cmd)
}

// Creates a new tmux pane and execute the pane commands
func (p *Pane) Create() {
	p.index = PANE_INDEX

	p.tmuxCommand.Execute()

	p.shell(p.commands)

	PANE_INDEX += 1
}

func (p *Pane) shell(commands []string) {
	for _, cmd := range commands {
		command.NewShellCommand(p.SessionName(), cmd).Execute()
	}
}
