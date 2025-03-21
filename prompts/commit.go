package prompts

import (
	"strings"
)

func GetCommit(diff string) string {
	cp := `
You are an expert software engineer writing semantic Git commit messages. Based on the diff provided, generate a single concise commit message that follows Conventional Commits standards.

The commit message must:
    * Use a valid semantic type (feat, fix, refactor, chore, test, docs, style, etc.). **Ensure that this is lowercased**.
    * Include a short, clear, and complete summary of what changed
    * Cover the full scope of changes in the diff
    * Use imperative mood (e.g., "add", "fix", "remove")
    * Be on one line (subject line only, no body)

Here is the diff:

{{diff}}

Respond with only the generated commit message. No explanations, no extra text.
`

	cp = strings.ReplaceAll(cp, "{{diff}}", diff)

	return cp
}
