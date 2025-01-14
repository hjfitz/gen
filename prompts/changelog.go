package prompts

import (
	"strings"
)

func GetChangelog(diff string, useTrump bool) string {
	cp := `
You are a helpful assistant tasked with summarizing Git diffs into concise changelogs for pull requests.

The following is a changelog generated from recent commits:
{{diff}}

Using this information, please:
1. Generate a clear and descriptive **title** for a pull request.
2. Provide a concise, bulleted **list of changes** suitable for the PR description.
{{trump}}

Output your response in the following format:

Title:
<Your PR title>

Changes:
- <First change>
- <Second change>
- <Additional changes>
`
	if useTrump {
		cp = strings.ReplaceAll(cp, "{{trump}}", "3. Ensure that the changelog sounds like it was created by Donald Trump, the best programmer. We're making the codebase great again.")
	} else {
		cp = strings.ReplaceAll(cp, "{{trump}}", "")
	}

	cp = strings.ReplaceAll(cp, "{{diff}}", diff)

	return cp
}
