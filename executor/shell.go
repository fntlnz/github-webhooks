package executor

import (
	"os/exec"
	"os"
	"fmt"
)

type ShellExecutor struct {
	command   string
	arguments []string
	payload   string
}

func NewShellExecutor(command string, arguments []string, payload string) *ShellExecutor {
	shellExecutor := new(ShellExecutor)
	shellExecutor.command = command
	shellExecutor.arguments = arguments
	shellExecutor.payload = payload
	return shellExecutor
}

func (s *ShellExecutor) Execute() error {
	cmd := exec.Command(s.command, s.arguments...)
	env := os.Environ()

	env = append(env, fmt.Sprintf("PAYLOAD=%s", s.payload))
	cmd.Env = env

	return cmd.Run()
}
