package git

import (
	"fmt"
	"os/exec"
)

func GetDiff(dirname, depth string) string {
	args := []string{
		"log",
		"-n", fmt.Sprintf("%s", depth),
		"--stat",
		"--patch",
		"--pretty=format:COMMIT: %H%nAUTHOR: %an%nDATE: %ad%nSUBJECT: %s%n%nBODY:%n%b%n---",
	}

	sh := exec.Command("git", args...)
	stdout, _ := sh.Output()

	return string(stdout)
}
