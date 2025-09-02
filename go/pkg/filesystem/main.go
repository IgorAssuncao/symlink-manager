package filesystem

import (
	"fmt"
	"os"
)

func CreateSymlink(target string, symlinkPath string) error {
	return os.Symlink(target, symlinkPath)
}

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err!= nil {
		fmt.Printf("Error from FileExists: %v\n", err)
		return false
	} 
	return true
}
