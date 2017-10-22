package command

/*
	This code is an utility to mock or stubs commands
*/

type FakeCommand struct {
	TmuxCommand
	ExecuteCalled bool
}

// Creates a new fake command
func NewFakeCommand(options string) *FakeCommand {
	return &FakeCommand{
		TmuxCommand:   NewTmuxCommand("fake-tmux-command", options),
		ExecuteCalled: false,
	}
}

func (c *FakeCommand) Execute() {
	c.ExecuteCalled = true
}
