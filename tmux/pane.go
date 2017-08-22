package tmux

import (
	"fmt"
	"os"
	"os/exec"

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

	WINDOW_INDEX = 0

	p.shell(p.commands)

	PANE_INDEX += 1
}

func (p *Pane) shell(commands []string) {
	for _, cmd := range commands {
		shell := []string{
			"send-keys",
			"-t",
			p.sessionName,
			cmd,
			"c-m",
		}

		cmd := exec.Command("tmux", shell...)
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
}
