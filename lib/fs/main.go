package fs

import (
	"fmt"
)

/*
func shouldBeIgnored(file string, toIgnore []string) bool {
	return true
}
*/

func GetCodebase(base string, toIgnore []string) string {
	defaultIgnores := []string{
    "__test__",
    "dist",
    "__fixtures__",
    "node_modules",
    "spec.js",
    "spec.ts",
    "__tests__",
	}

	ignoredDirs := append(defaultIgnores, toIgnore...)

	fmt.Println(ignoredDirs)

	return ""
}
