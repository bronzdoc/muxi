package tmux

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Represents a tmux layout
type Layout struct {
	fileName    string
	content     map[string][]interface{}
	tmuxSession *Session
}

// Creates a new layout from a yaml file
func NewLayout(fileName string) *Layout {
	return &Layout{
		fileName:    fileName,
		tmuxSession: NewSession(""),
	}
}

// Creates a new tmux layout
func (l *Layout) Create() error {
	err := l.parse()
	if err != nil {
		return fmt.Errorf("Parse error: %v", err)
	}

	l.tmuxSession.Create()

	return nil
}

func (l *Layout) parse() error {
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
		windowName := window.(map[interface{}]interface{})["name"]
		tmuxWindow := NewWindow(windowName.(string))

		if panes, ok := window.(map[interface{}]interface{})["panes"]; ok {
			for _, paneCommand := range panes.([]interface{}) {
				switch pcType := paneCommand.(type) {
				default:
					return fmt.Errorf("Invalid pane command: %v", pcType)
				case map[interface{}]interface{}: // Multiple commands for a pane
					if commands, ok := paneCommand.(map[interface{}]interface{})["commands"]; ok {
						tmuxPane := NewPane()

						for _, command := range commands.([]interface{}) {
							tmuxPane.AddCommand(command.(string))
						}

						tmuxWindow.AddPane(tmuxPane)
					}
				case string: // A single command for each pane
					tmuxPane := NewPane()
					tmuxPane.AddCommand(paneCommand.(string))
					tmuxWindow.AddPane(tmuxPane)
				}
			}
		}

		l.tmuxSession.AddWindow(tmuxWindow)
	}

	return nil
}
