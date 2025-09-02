package main

import (
	"fmt"
	"os"

	"github.com/IgorAssuncao/symlink-manager/go/internal/config"
	"github.com/IgorAssuncao/symlink-manager/go/pkg/filesystem"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("%v", err)
	}

	c := config.GetConfig("config.yaml")

	for name, tool := range c.Tools {
		path := homeDir + "/" + tool.Path
		target := homeDir + "/" + tool.Target

		fmt.Printf("Checking if %s symlink path exists\n", name)
		if filesystem.FileExists(path) {
			fmt.Printf("%s already exists\n", path)
			fmt.Println()
			continue
		}

		fmt.Printf("Creating symlink for: %s\n", name)
		if err := filesystem.CreateSymlink(target, path); err != nil {
			fmt.Printf("Error creating symlink: %v\n", err)
		}
		fmt.Println()
	}
	// TODO v2: Make it run as a daemon (service).
	// TODO: v3: instead of checking all of them, diff the config file and
	// change only the key that was changed.
}

