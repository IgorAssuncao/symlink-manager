package filesystem

import (
	"log"
	"os"
	"time"
)

func CreateSymlink(target string, symlinkPath string) error {
	return os.Symlink(target, symlinkPath)
}

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err!= nil {
		log.Printf("Error from FileExists: %v\n", err)
		return false
	} 
	return true
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("%v", err)
	}

	return homeDir
}

func GetFileLastModified(configFilePath string) time.Time {
	fileInfo, err:= os.Stat(configFilePath)
	if err != nil {
		log.Printf("%v", err) 
	}
	return fileInfo.ModTime()
}

