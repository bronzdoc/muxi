package tmux

import (
	"fmt"

	"github.com/bronzdoc/muxi/tmux/command"
)

var WINDOW_INDEX = 0

// Window Represents a tmux windows
type Window struct {
	tmuxObject
	name   string
	panes  []*Pane
	layout string
	root   string
}

// Create a new Window
func NewWindow(name, layout, root string) *Window {
	WINDOW_INDEX += 1

	return &Window{
		name:   name,
		layout: layout,
		root:   root,
		tmuxObject: tmuxObject{
			tmuxCommand: command.NewWindowCommand(
				fmt.Sprintf("-n %s", name),
				fmt.Sprintf("-c %s", root),
			),
		},
	}
}

// Get window panes
func (w *Window) Panes() []*Pane {
	return w.panes
}

// Add a new window pane
func (w *Window) AddPane(pane *Pane) {
	pane.SetSessionName(w.SessionName())
	w.panes = append(w.panes, pane)
}

// Get window name
func (w *Window) Name() string {
	return w.name
}

// Creates a new tmux window and its child panes
func (w *Window) Create() {
	w.tmuxCommand.Execute()
	w.createPanes()
	command.NewSelectLayoutCommand(w.layout).Execute()
}

func (w *Window) createPanes() {
	if len(w.panes) != 0 {
		firstIndex := 0
		firstPane := w.panes[firstIndex]

		w.shell(firstPane.commands)

		for _, p := range w.panes[(firstIndex + 1):] {
			p.Create()
		}
	}
}

func (w *Window) shell(commands []string) {
	for _, cmd := range commands {
		command.NewShellCommand(w.SessionName(), cmd).Execute()
	}
}
