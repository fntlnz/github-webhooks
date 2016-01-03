package executor

import "os/exec"

type ShellExecutor struct {
	command   string
	arguments []string
}

func NewShellExecutor(command string, arguments []string) *ShellExecutor {
	shellExecutor := new(ShellExecutor)
	shellExecutor.command = command
	shellExecutor.arguments = arguments
	return shellExecutor
}

func (s *ShellExecutor) Execute() error {
	cmd := exec.Command(s.command, s.arguments...)
	return cmd.Run()
}
