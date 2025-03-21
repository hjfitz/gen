package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDiff(dirname, depth string) string {
	return git(
		"log",
		"-n",
		fmt.Sprintf("%s", depth),
		"--stat",
		"--patch",
		"--pretty=format:COMMIT: %H%nAUTHOR: %an%nDATE: %ad%nSUBJECT: %s%n%nBODY:%n%b%n---",
	)
}

func GetMostRecentChanges() string {
	return git("diff", "--cached")
}

func git(args ...string) string {
	sh := exec.Command("git", args...)

	stdout, _ := sh.Output()

	return strings.TrimSpace(string(stdout))
}
