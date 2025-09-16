package tools

import (
	"log"

	"github.com/IgorAssuncao/symlink-manager/go/internal/config"
	"github.com/IgorAssuncao/symlink-manager/go/pkg/filesystem"
)

func ProcessTools(tools map[string]config.Tool) {
	homeDir := filesystem.GetHomeDir()

	for name, tool := range tools {
		path := homeDir + "/" + tool.Path
		target := homeDir + "/" + tool.Target

		log.Printf("Checking if %s symlink path exists\n", name)
		if filesystem.FileExists(path) {
			log.Printf("%s already exists\n", path)
			log.Println()
			continue
		}

		log.Printf("Creating symlink for: %s\n", name)
		if err := filesystem.CreateSymlink(target, path); err != nil {
			log.Printf("Error creating symlink: %v\n", err)
		}
		log.Println()
	}

}
