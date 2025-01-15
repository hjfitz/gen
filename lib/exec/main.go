package exec

import (
	"fmt"
	"os/exec"
)

func Exec(cmdStr string, args []string) string {
	cmd := exec.Command(cmdStr)
	stdout, err := cmd.Output()

	fmt.Println(err)
	fmt.Println(stdout)

	return string(stdout)
}
