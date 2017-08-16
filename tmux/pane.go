package tmux

import (
	"fmt"
	"os"
	"os/exec"
)

var PANE_INDEX = 0

type Pane struct {
	index       int
	command     []map[string]interface{}
	sessionName string
	commands    []string
}

func NewPane() *Pane {
	return &Pane{}
}

func (p *Pane) Setup(sessionName string) {
	p.sessionName = sessionName
	p.command = []map[string]interface{}{
		{
			"cmd": BASECOMMAND,
			"args": []string{
				"split-window",
			},
		},
	}
}

// Adds a new command to execute in pane
func (p *Pane) AddCommand(cmd string) {
	p.commands = append(p.commands, cmd)
}

// Creates a new tmux pane
func (p *Pane) Create() {
	p.index = PANE_INDEX

	for _, c := range p.command {
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
