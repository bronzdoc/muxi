package tmux_test

import (
	"github.com/bronzdoc/muxi/command"
	. "github.com/bronzdoc/muxi/tmux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session", func() {
	var session *Session
	var mockCommand *command.FakeCommand

	BeforeEach(func() {
		mockCommand = command.NewFakeCommand("options")

		session = NewSession("test-session")
		session.SetTmuxCommand(mockCommand)

	})

	Describe("AddWindow", func() {
		It("should add a new window to a session", func() {
			window := NewWindow("test-window", "default", "")
			session := NewSession("test-session")
			session.AddWindow(window)

			Expect(len(session.Windows())).To(Equal(1))
		})
	})

	Describe("Name", func() {
		It("should return the session name", func() {
			Expect(session.Name()).To(Equal("test-session"))
		})
	})

	Describe("Create", func() {
		It("Create a new tmux session", func() {
			session.Create()

			Expect(mockCommand.ExecuteCalled).To(Equal(true))
		})
	})
})
