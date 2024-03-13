package assestserver

import "io/fs"

// https://github.com/golang/go/issues/43431#issuecomment-752662261
func CD(embedFS fs.FS, root string) fs.FS {
	newFS, err := fs.Sub(embedFS, root)
	if err != nil {
		panic(err)
	}
	return newFS
}
