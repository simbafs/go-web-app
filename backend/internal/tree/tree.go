package tree

import (
	"fmt"
	"io/fs"
	"strings"
)

// Tree act like tree commend in linux, reture the output
func Tree(FS fs.FS) (string, error) {
	var treeOutput strings.Builder

	// WalkFunc to traverse the directory structure
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate depth of current directory/file
		depth := strings.Count(strings.TrimPrefix(path, "/"), "/")

		// Indentation based on depth
		indentation := strings.Repeat("│   ", depth)

		// Append directory/file name with proper indentation
		if d.IsDir() {
			treeOutput.WriteString(fmt.Sprintf("%s├── %s/\n", indentation, d.Name()))
		} else {
			treeOutput.WriteString(fmt.Sprintf("%s├── %s\n", indentation, d.Name()))
		}

		if d.Name() == "_next" && d.IsDir() {
			treeOutput.WriteString(fmt.Sprintf("%s│   ├── ...\n", indentation))
			return fs.SkipDir
		}

		return nil
	}

	// Traverse the directory structure starting from the root
	if err := fs.WalkDir(FS, ".", walkFunc); err != nil {
		return "", err
	}

	return treeOutput.String(), nil
}
