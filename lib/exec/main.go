package exec

import (
	"os/exec"
	"strings"
)

func Exec(cmdStr string) string {
	cmd := exec.Command(cmdStr)
	stdout, _ := cmd.Output()

	return strings.TrimSpace(string(stdout))
}
