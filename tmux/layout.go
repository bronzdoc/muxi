package tmux

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
			tmuxPane := NewPane()

			for _, pane := range panes.([]interface{}) {
				if commands, ok := pane.(map[interface{}]interface{})["commands"]; ok {
					for _, command := range commands.([]interface{}) {
						tmuxPane.AddCommand(command.(string))
					}
				} else {
					command := panes.([]interface{})[0]
					tmuxPane.AddCommand(command.(string))
				}
			}

			tmuxWindow.AddPane(tmuxPane)
		}

		l.tmuxSession.AddWindow(tmuxWindow)
	}

	return nil
}

// Creates a new tmux layout
func (l *Layout) Create() error {
	err := l.parse()
	if err != nil {
		return err
	}

	l.tmuxSession.Create()

	return nil
}
