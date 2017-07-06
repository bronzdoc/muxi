package tmux

import (
	"fmt"
	"os"
	"os/exec"
)

var WINDOW_INDEX = 0

type Window struct {
	sessionName string
	name        string
	panes       []*Pane
	command     []map[string]interface{}
}

func NewWindow(name string) *Window {
	WINDOW_INDEX += 1

	return &Window{
		name: name,
	}
}

func (w *Window) Setup(sessionName string) {
	w.sessionName = sessionName
	w.command = []map[string]interface{}{
		{
			"cmd": BASECOMMAND,
			"args": []string{
				"new-window",
			},
		},
	}
}

// Get window panes
func (w *Window) Panes() []*Pane {
	return w.panes
}

// Add a new window pane
func (w *Window) AddPane(pane *Pane) {
	pane.Setup(w.SessionName())
	w.panes = append(w.panes, pane)
}

// Get window name
func (w *Window) Name() string {
	return w.name
}

// Creates a new tmux session
func (w *Window) Create() {
	for _, c := range w.command {
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

	// Create session windows
	for _, p := range w.panes {
		p.Create()
	}
}

func (w *Window) SessionName() string {
	return w.sessionName
}
