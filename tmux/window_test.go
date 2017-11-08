package tmux_test

import (
	"github.com/bronzdoc/muxi/command"
	. "github.com/bronzdoc/muxi/tmux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window", func() {
	var window *Window
	var mockCommand *command.FakeCommand

	BeforeEach(func() {
		mockCommand = command.NewFakeCommand("options")

		window = NewWindow("test-window", "default", "")
		window.SetTmuxCommand(mockCommand)
	})

	Describe("AddPane", func() {
		It("should add a new pane to a window", func() {
			pane := NewPane("/tmp")
			window.AddPane(pane)

			Expect(len(window.Panes())).To(Equal(1))
		})
	})

	Describe("Name", func() {
		It("should return the window name", func() {
			Expect(window.Name()).To(Equal("test-window"))
		})
	})

	Describe("Create", func() {
		It("create a new tmux window", func() {
			window.Create()
			Expect(mockCommand.ExecuteCalled).To(Equal(true))
		})

		Context("when no name given", func() {
			It("should not add the -n option to the tmux command", func() {
				windowCommand := NewWindow("", "default", "/tmp").GetTmuxCommand()
				Expect(windowCommand.Options()).To(Equal([]string{"", "-c /tmp"}))
			})
		})

		Context("when no root given", func() {
			It("should not add the -c option to the tmux command", func() {
				windowCommand := NewWindow("my-new-window", "default", "").GetTmuxCommand()
				Expect(windowCommand.Options()).To(Equal([]string{"-n my-new-window", ""}))
			})
		})
	})
})
