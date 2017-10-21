package layout

import (
	"fmt"
	"github.com/bronzdoc/muxi/tmux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Represents a muxi layout
type Layout struct {
	fileName    string
	content     map[string][]interface{}
	TmuxSession *tmux.Session
}

// Creates a new muxi layout from a yaml file
func NewLayout(fileName string) *Layout {
	return &Layout{
		fileName:    fileName,
		TmuxSession: tmux.NewSession(""),
	}
}

// Creates a new tmux layout based on a muxi layout
func (l *Layout) Create() {
	l.TmuxSession.Create()
}

// Gets a muxi layout content
func (l *Layout) Content() map[string][]interface{} {
	return l.content
}

// Parses a muxi Layout
func (l *Layout) Parse() error {
	yamlFile, err := ioutil.ReadFile(l.fileName)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &l.content)

	if err != nil {
		return err
	}

	windows := l.content["windows"]

	for _, window := range windows {
		tmuxWindow := tmux.NewWindow(
			getWindowName(window),
			getWindowLayout(window),
			getWindowRoot(window),
		)

		l.TmuxSession.AddWindow(tmuxWindow)

		for _, paneCommand := range getWindowPanes(window) {
			switch pcType := paneCommand.(type) {
			default:
				return fmt.Errorf("Invalid pane command: %v", pcType)
			case map[interface{}]interface{}: // Multiple commands for a pane
				if commands, ok := paneCommand.(map[interface{}]interface{})["commands"]; ok {
					tmuxPane := tmux.NewPane(getWindowRoot(window))

					for _, command := range commands.([]interface{}) {
						tmuxPane.AddCommand(command.(string))
					}

					tmuxWindow.AddPane(tmuxPane)
				}
			case string: // A single command for each pane
				tmuxPane := tmux.NewPane(getWindowRoot(window))
				tmuxPane.AddCommand(paneCommand.(string))
				tmuxWindow.AddPane(tmuxPane)
			}
		}
	}

	return nil
}

func getWindowPanes(context interface{}) []interface{} {
	return getWindowSliceField(context, "panes")
}

func getWindowName(context interface{}) string {
	return getWindowStringField(context, "name")
}

func getWindowRoot(context interface{}) string {
	return getWindowStringField(context, "root")
}

func getWindowLayout(context interface{}) string {
	return getWindowStringField(context, "layout")
}

func getWindowSliceField(context interface{}, fieldName string) []interface{} {
	field := make([]interface{}, 0)

	switch context.(type) {
	case map[interface{}]interface{}:
		if contextField, ok := context.(map[interface{}]interface{})[fieldName]; ok {
			field = contextField.([]interface{})
		}
	}

	return field
}

func getWindowStringField(context interface{}, field string) string {
	switch context.(type) {
	default:
		return ""

	case map[interface{}]interface{}:
		if name, ok := context.(map[interface{}]interface{})[field]; ok {
			return name.(string)
		} else {
			return ""
		}
	}
}
