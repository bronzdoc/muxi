package tmux

import (
	"github.com/bronzdoc/muxi/tmux/command"
)

var WINDOW_INDEX = 0

type Window struct {
	sessionName string
	name        string
	panes       []*Pane
	tmuxCommand *command.NewWindow
}

func NewWindow(name string) *Window {
	WINDOW_INDEX += 1

	return &Window{
		name:        name,
		tmuxCommand: command.NewWindowCommand(),
	}
}

func (w *Window) Setup(sessionName string) {
	w.sessionName = sessionName
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

// Creates a new tmux window and its child panes
func (w *Window) Create() {
	w.tmuxCommand.Execute()
	w.createPanes()
}

func (w *Window) SessionName() string {
	return w.sessionName
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
		command.NewShellCommand(w.sessionName, cmd).Execute()
	}
}
