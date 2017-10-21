package layout_test

import (
	"io/ioutil"

	"github.com/bronzdoc/muxi/command"
	. "github.com/bronzdoc/muxi/layout"
	"github.com/bronzdoc/muxi/tmux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Layout", func() {
	Describe("Parse", func() {
		It("Should parse the tmux layout definition correctly", func() {

			tmux_template_content := []byte(`---
windows:
  - name: test
    root: /tmp
    layout: tiled
    panes:
      - ls -liah
      - env
      - echo "jar jar binks"
      - vim test.yml
 `)

			tmux_template_file := "/tmp/test.yml"
			ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
			layout := NewLayout(tmux_template_file)
			err := layout.Parse()

			Expect(err).To(BeNil())

			layout_content := layout.Content()
			window_keys := layout_content["windows"][0].(map[interface{}]interface{})

			Expect(window_keys["name"]).To(Equal("test"))
			Expect(window_keys["root"]).To(Equal("/tmp"))
			Expect(window_keys["layout"]).To(Equal("tiled"))
			Expect(len(window_keys["panes"].([]interface{}))).To(Equal(4))
			Expect(window_keys["panes"].([]interface{})[0]).To(Equal("ls -liah"))
			Expect(window_keys["panes"].([]interface{})[1]).To(Equal("env"))
			Expect(window_keys["panes"].([]interface{})[2]).To(Equal("echo \"jar jar binks\""))
			Expect(window_keys["panes"].([]interface{})[3]).To(Equal("vim test.yml"))
		})
	})

	Describe("Create", func() {
		It("should create a muxi layout", func() {
			mockCommand := command.NewFakeCommand("options")
			session := tmux.Session{}
			session.SetTmuxCommand(mockCommand)

			layout := Layout{TmuxSession: &session}

			Expect(mockCommand.ExecuteCalled).To(Equal(false))

			layout.Create()

			Expect(mockCommand.ExecuteCalled).To(Equal(true))
		})
	})
})
