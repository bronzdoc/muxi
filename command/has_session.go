package command

type hasSessionCommand struct {
	TmuxCommand
}

func NewHasSessionCommand(sessionName string) *hasSessionCommand {
	return &hasSessionCommand{
		TmuxCommand: NewTmuxCommand("has-session -t", sessionName),
	}
}
