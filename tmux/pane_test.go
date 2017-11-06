package tmux_test

import (
	"github.com/bronzdoc/muxi/command"
	. "github.com/bronzdoc/muxi/tmux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pane", func() {
	var pane *Pane
	var mockCommand *command.FakeCommand

	BeforeEach(func() {
		mockCommand = command.NewFakeCommand("options")

		pane = NewPane("test-pane")
		pane.SetTmuxCommand(mockCommand)

	})

	Describe("AddCommand", func() {
		It("should add a new window to a pane", func() {
			pane.AddCommand("ls")

			Expect(len(pane.Commands())).To(Equal(1))
		})
	})

	Describe("Create", func() {
		It("Create a new tmux pane", func() {
			pane.Create()

			Expect(mockCommand.ExecuteCalled).To(Equal(true))
		})

		Context("when no root given", func() {
			It("should not add the -c option to the tmux command", func() {
				paneCommand := NewPane("").GetTmuxCommand()
				Expect(paneCommand.Options()).To(Equal([]string{""}))
			})
		})
	})
})
