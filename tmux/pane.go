package tmux

import (
	"github.com/bronzdoc/muxi/tmux/command"
)

var PANE_INDEX = 0

type Pane struct {
	index       int
	sessionName string
	tmuxCommand *command.NewPane
	commands    []string
}

func NewPane() *Pane {
	return &Pane{
		tmuxCommand: command.NewPaneCommand(),
	}
}

func (p *Pane) Setup(sessionName string) {
	p.sessionName = sessionName
}

// Adds a new command to execute in pane
func (p *Pane) AddCommand(cmd string) {
	p.commands = append(p.commands, cmd)
}

// Creates a new tmux pane
func (p *Pane) Create() {
	p.index = PANE_INDEX

	p.tmuxCommand.Execute()

	p.shell(p.commands)

	PANE_INDEX += 1
}

func (p *Pane) shell(commands []string) {
	for _, cmd := range commands {
		command.NewShellCommand(p.sessionName, cmd).Execute()
	}
}
