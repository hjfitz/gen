package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"io/fs"
	"strings"

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
	content := string(dat)
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
				//fmt.Printf("Skipping directory: %q\n", path)
				return filepath.SkipDir
			}
			//fmt.Printf("Skipping file: %q\n", path)
			return nil
		}

		codebase += readFile(path)
		return nil
	})

	return codebase
}
