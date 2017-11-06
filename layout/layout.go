package layout

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"

	"github.com/bronzdoc/muxi/tmux"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ExecCommand allows to Have higher control of exec.Command,
// and will allow us to mock it easier in tests...
var ExecCommand = exec.Command

// Represents a muxi layout
type Layout struct {
	fileName    string
	content     map[string][]interface{}
	TmuxSession *tmux.Session
}

// Creates a new muxi layout from a yaml file
func New(fileName string) *Layout {
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

// Gets a muxi layout content
func (l *Layout) RawContent() ([]byte, error) {
	return getLayoutContent(l.fileName)
}

func Edit(layoutName string) error {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		return fmt.Errorf(`$EDITOR is empty, could not edit "%s" layout`, layoutName)
	}

	layoutPath, err := getLayoutPath(layoutName)
	if err != nil {
		return fmt.Errorf("layout not found: %s", err)
	}

	cmd := ExecCommand(editor, layoutPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// List muxi layouts
func List() (list []string) {
	layoutsPath := viper.GetString("layoutsPath")
	files, _ := ioutil.ReadDir(layoutsPath)
	for _, f := range files {
		// List only files with yaml or yml extension
		hasValidExtension, _ := regexp.MatchString("(.yml|.yaml)", f.Name())
		if hasValidExtension {
			list = append(list, f.Name())
		}
	}

	return
}

// Parses a muxi Layout
func (l *Layout) Parse() error {

	yamlFileContent, err := getLayoutContent(l.fileName)
	if err != nil {
		return fmt.Errorf("parse: %s\n", err)
	}

	err = yaml.Unmarshal(yamlFileContent, &l.content)

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
	switch context.(type) {
	case map[interface{}]interface{}:
		if contextField, ok := context.(map[interface{}]interface{})[fieldName]; ok {
			switch contextField.(type) {
			case []interface{}:
				return contextField.([]interface{})
			}
		}
	}

	return make([]interface{}, 0)
}

func getWindowStringField(context interface{}, field string) string {
	switch context.(type) {
	case map[interface{}]interface{}:
		if name, ok := context.(map[interface{}]interface{})[field]; ok {
			switch name.(type) {
			case string:
				return name.(string)
			}
		}
	}

	return ""
}

func layoutExists(layoutPath string) bool {
	_, err := os.Stat(layoutPath)
	return !os.IsNotExist(err) // negate the bool so the functions makes sense...
}

func getLayoutPath(layoutName string) (string, error) {
	yamlFile := getLayoutWithExtension(layoutName, "yaml")
	ymlFile := getLayoutWithExtension(layoutName, "yml")

	if layoutExists(yamlFile) {
		return yamlFile, nil
	} else if layoutExists(ymlFile) {
		return ymlFile, nil
	}

	return "", fmt.Errorf(`layout "%s" doesn't exists`, layoutName)
}

func getLayoutWithExtension(layoutName, extension string) string {
	return fmt.Sprintf(
		"%s/%s",
		viper.GetString("layoutsPath"),
		fmt.Sprintf("%s.%s", layoutName, extension),
	)
}

func getLayoutContent(layoutName string) ([]byte, error) {
	layoutPath, err := getLayoutPath(layoutName)
	if err != nil {
		return []byte{}, fmt.Errorf("layout not found: %s", err)
	}

	yamlFileContent, err := ioutil.ReadFile(layoutPath)
	if err != nil {
		return []byte{}, fmt.Errorf("can't read layout: %s", err)
	}

	return yamlFileContent, nil
}
