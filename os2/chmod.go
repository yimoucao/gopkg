package os2

import (
	"io/fs"
	"os"
	"path/filepath"
)

func ChmodRecursively(root string, mode fs.FileMode) error {
	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			err = os.Chmod(path, mode)
			if err != nil {
				return err
			}
			return nil
		})
}
