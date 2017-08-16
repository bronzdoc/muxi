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
	commands    []string
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

// Creates a new tmux window
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

	if len(w.panes) != 0 {
		firstIndex := 0
		firstPane := w.panes[firstIndex]

		w.shell(firstPane.commands)

		for _, p := range w.panes[(firstIndex + 1):] {
			p.Create()
		}
	}
}

func (w *Window) SessionName() string {
	return w.sessionName
}

func (w *Window) shell(commands []string) {
	for _, cmd := range commands {
		shell := []string{
			"send-keys",
			"-t",
			w.sessionName,
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
