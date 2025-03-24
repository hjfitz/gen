package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDiff(depth string) string {
	return git(
		"log",
		"-n",
		fmt.Sprintf("%s", depth),
		"--stat",
		"--patch",
		"--pretty=format:COMMIT: %H%nAUTHOR: %an%nDATE: %ad%nSUBJECT: %s%n%nBODY:%n%b%n---",
	)
}

func GetMainBranch() string {
	bs := git("branch", "--format=%(refname:short)")
	b_list := strings.Split(bs, "\n")
	for _, b := range b_list {
		if b == "master" {
			return "master"
		}
		if b == "main" {
			return "main"
		}
	}
	return ""
}

func GetDiffAgainstMain(branch string) string {
	return git("diff", fmt.Sprintf("%s...HEAD", branch))
}

func GetMostRecentChanges() string {
	return git("diff", "--cached")
}

func git(args ...string) string {
	sh := exec.Command("git", args...)

	stdout, _ := sh.Output()

	return strings.TrimSpace(string(stdout))
}
