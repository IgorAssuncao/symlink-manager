package filesystem

import "os"

func createSymlink(symlinkPath string, target string) error {
	return os.Symlink(symlinkPath, target)
}
