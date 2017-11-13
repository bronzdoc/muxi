package tmux

import (
	"fmt"

	"github.com/bronzdoc/muxi/command"
	"github.com/spf13/viper"
)

var WINDOW_COUNT = 0

// Window Represents a tmux windows
type Window struct {
	tmuxObject
	index  int
	name   string
	panes  []*Pane
	layout string
	root   string
}

// Create a new Window
func NewWindow(name, layout, root string) *Window {
	w := Window{
		index:      WINDOW_COUNT,
		name:       name,
		layout:     layout,
		root:       root,
		tmuxObject: tmuxObject{},
	}

	if viper.GetBool("here") && w.isFirst() {
		w.SetTmuxCommand(command.CurrentWindowCommand())
	} else {
		windowRoot := root
		windowName := name

		if !IsEmpty(root) {
			windowRoot = fmt.Sprintf("-c %s", root)
		}

		if !IsEmpty(name) {
			windowName = fmt.Sprintf("-n %s", name)
		}

		w.SetTmuxCommand(
			command.NewWindowCommand(
				windowName,
				windowRoot,
			),
		)
	}

	w.tmuxCommand.AddPostHook(w.createPanes)
	w.tmuxCommand.AddPostHook(w.selectLayout)

	WINDOW_COUNT += 1

	return &w
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

func (w *Window) selectLayout() {
	command.NewSelectLayoutCommand(w.layout).Execute()
}

func (w *Window) shell(commands []string) {
	for _, cmd := range commands {
		command.NewShellCommand(w.SessionName(), cmd).Execute()
	}
}

func (w *Window) isFirst() bool {
	return w.index == 0
}
