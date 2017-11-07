package layout_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bronzdoc/muxi/command"
	. "github.com/bronzdoc/muxi/layout"
	"github.com/bronzdoc/muxi/tmux"
	"github.com/spf13/viper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const MUXI_LAYOUTS_PATH = "/tmp/muxi_test"

var _ = Describe("Layout", func() {

	BeforeEach(func() {
		os.Mkdir(MUXI_LAYOUTS_PATH, 0777)
		viper.Set("layoutsPath", MUXI_LAYOUTS_PATH)
	})

	AfterEach(func() {
		os.RemoveAll(MUXI_LAYOUTS_PATH)
	})

	Describe("Parse", func() {
		It("Should parse the tmux layout definition correctly", func() {

			tmux_template_content := []byte(`---
name: my-session-name
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

			tmux_template_file := fmt.Sprintf("%s/test.yml", MUXI_LAYOUTS_PATH)
			ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
			layout := New("test")
			err := layout.Parse()

			Expect(err).To(BeNil())

			layout_content := layout.Content()

			session_name := layout_content["name"]
			window_keys := layout_content["windows"].([]interface{})[0].(map[interface{}]interface{})

			Expect(session_name).To(Equal("my-session-name"))
			Expect(window_keys["name"]).To(Equal("test"))
			Expect(window_keys["root"]).To(Equal("/tmp"))
			Expect(window_keys["layout"]).To(Equal("tiled"))
			Expect(len(window_keys["panes"].([]interface{}))).To(Equal(4))
			Expect(window_keys["panes"].([]interface{})[0]).To(Equal("ls -liah"))
			Expect(window_keys["panes"].([]interface{})[1]).To(Equal("env"))
			Expect(window_keys["panes"].([]interface{})[2]).To(Equal("echo \"jar jar binks\""))
			Expect(window_keys["panes"].([]interface{})[3]).To(Equal("vim test.yml"))
		})

		Context("when session name is empty", func() {
			It("should add the correct default", func() {
				tmux_template_content := []byte(`---
name:
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

				tmux_template_file := fmt.Sprintf("%s/test_empty_session_name.yml", MUXI_LAYOUTS_PATH)
				ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
				layout := New("test_empty_session_name")
				err := layout.Parse()

				Expect(err).To(BeNil())
			})
		})

		Context("when root is empty", func() {
			It("should add the correct default", func() {
				tmux_template_content := []byte(`---
windows:
  - name: test
    root:
    layout: tiled
    panes:
      - env
 `)

				tmux_template_file := fmt.Sprintf("%s/test_empty_name.yml", MUXI_LAYOUTS_PATH)
				ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
				layout := New("test_empty_name")
				err := layout.Parse()

				Expect(err).To(BeNil())
			})
		})

		Context("when name is empty", func() {
			It("should add the correct default", func() {
				tmux_template_content := []byte(`---
windows:
  - name:
    root: /tmp
    layout: tiled
    panes:
      - env
 `)
				tmux_template_file := fmt.Sprintf("%s/test_empty_name.yml", MUXI_LAYOUTS_PATH)
				ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
				layout := New("test_empty_name")
				err := layout.Parse()

				Expect(err).To(BeNil())
			})
		})

		Context("when layout is empty", func() {
			It("should add the correct default", func() {
				tmux_template_content := []byte(`---
windows:
  - name: test
    root: /tmp
    layout:
    panes:
      - env
 `)
				tmux_template_file := fmt.Sprintf("%s/test_empty_layout.yml", MUXI_LAYOUTS_PATH)
				ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
				layout := New("test_empty_layout")
				err := layout.Parse()

				Expect(err).To(BeNil())
			})
		})

		Context("when panes is empty", func() {
			It("should add the correct default", func() {
				tmux_template_content := []byte(`---
windows:
  - name: test
    root: /tmp
    layout:
    panes:
 `)
				tmux_template_file := fmt.Sprintf("%s/test_empty_panes.yml", MUXI_LAYOUTS_PATH)
				ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
				layout := New("test_empty_panes")
				err := layout.Parse()

				Expect(err).To(BeNil())
			})
		})
	})

	Describe("RawContent", func() {
		It("should show layout raw content", func() {
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

			tmux_template_file := fmt.Sprintf("%s/test.yml", MUXI_LAYOUTS_PATH)
			ioutil.WriteFile(tmux_template_file, tmux_template_content, 0777)
			layout := New("test")

			Expect(layout.RawContent()).To(Equal(tmux_template_content))
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
