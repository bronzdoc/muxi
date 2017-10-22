package command

type NewWindow struct {
	TmuxCommand
}

func NewWindowCommand(options ...string) *NewWindow {
	return &NewWindow{
		TmuxCommand: NewTmuxCommand("new-window", options...),
	}
}
