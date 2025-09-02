package main

import (
	"fmt"
	"os"

	"github.com/IgorAssuncao/symlink-manager/go/internal/config"
)

func main() {
	println("Hi!")

	currDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v", err)
	}

	println(currDir)

	c := config.Config{}

	c.GetConfig()

	for _, t := range c.Tools {
		fmt.Printf("%v", t)
	}
	// TODO: loop through tools array checking if symlink or file exists,
	// and then creating the symlink.
	// v2: instead of checking all of them, diff the config file and
	// change only the key that was changed.
}

