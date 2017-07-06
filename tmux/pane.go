package tmux

import (
	"fmt"
	"os"
	"os/exec"
)

type Pane struct {
	command []map[string]interface{}
	shell   []string
}

func NewPane() *Pane {
	return &Pane{}
}

func (p *Pane) Setup(sessionName string) {
	p.command = []map[string]interface{}{
		{
			"cmd": BASECOMMAND,
			"args": []string{
				"send-keys",
				fmt.Sprintf("%s split-window -t %s", BASECOMMAND, sessionName),
				"c-m",
			},
		},
	}
}

// Creates a new tmux pane
func (p *Pane) Create() {
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
}
