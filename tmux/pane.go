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
	p := Pane{}

	paneRoot := root

	if !IsEmpty(root) {
		paneRoot = fmt.Sprintf("-c %s", root)
	}

	p.SetTmuxCommand(
		command.NewPaneCommand(paneRoot),
	)

	p.tmuxCommand.AddPostHook(p.shell)

	return &p
}

// Adds a new command to execute in pane
func (p *Pane) AddCommand(cmd string) {
	p.commands = append(p.commands, cmd)
}

func (p *Pane) Commands() []string {
	return p.commands
}

// Creates a new tmux pane and execute the pane commands
func (p *Pane) Create() {
	p.index = PANE_INDEX

	p.tmuxCommand.Execute()

	PANE_INDEX += 1
}

func (p *Pane) shell() {
	for _, cmd := range p.commands {
		command.NewShellCommand(p.SessionName(), cmd).Execute()
	}
}
