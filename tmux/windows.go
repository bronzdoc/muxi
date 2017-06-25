package tmux

import (
	"fmt"
	"os"
	"os/exec"
)

type Window struct {
	session string
	name    string
	panes   []Pane
	command []map[string]interface{}
}

func NewWindow(session, name string, panes []Pane) *Window {
	return &Window{
		name:  name,
		panes: panes,
		command: []map[string]interface{}{
			{
				"cmd": BASECOMMAND,
				"args": []string{
					"send-keys",
					"-t",
					session,
					fmt.Sprintf("%s new-window", BASECOMMAND),
					"c-m",
				},
			},
		},
	}
}

// Get window panes
func (w *Window) Panes() []Pane {
	return w.panes
}

// Get window name
func (w *Window) Name() string {
	return w.name
}

// Creates a new tmux session
func (w *Window) Create() {
	for _, c := range w.command {
		fmt.Println(c["args"])
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
