package fs

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	gitignore "github.com/sabhiram/go-gitignore"
)

func getGitignore(base string) *gitignore.GitIgnore {
	giPath := filepath.Join(base, ".gitignore")
	dat, err := os.ReadFile(giPath)
	if err != nil {
		fmt.Printf("Error reading .gitignore: %v\n", err)
		return gitignore.CompileIgnoreLines("")
	}

	lines := strings.Split(string(dat), "\n")
	ignoredFiles := append(lines, ".git") // Always ignore `.git` folder
	return gitignore.CompileIgnoreLines(ignoredFiles...)
}

func readFile(filename string) string {
	dat, _ := os.ReadFile(filename)
	// Validate and convert content to UTF-8 if necessary
	content := string(dat)
	if !utf8.ValidString(content) {
		// Convert to UTF-8 by replacing invalid sequences
		content = string(bytes.ToValidUTF8([]byte(content), []byte("?")))
	}

	if len(content) == 0 {
		return ""
	}

	wd, _ := os.Getwd()

	localFilename := strings.ReplaceAll(filename, wd, "")

	return fmt.Sprintf("## filename: %s\n\n%s\n\n\n", localFilename, content)
}

func GetCodebase(base string) string {
	gitignore := getGitignore(base)


	codebase := ""
	_ = filepath.Walk(base, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if gitignore != nil && gitignore.MatchesPath(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if os.Getenv("DEBUG") != "" {
			fmt.Printf("Loaded %s\n", path)
		}
		codebase += readFile(path)
		return nil
	})

	return codebase
}
