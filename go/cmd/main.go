package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/IgorAssuncao/symlink-manager/go/internal/config"
	"github.com/IgorAssuncao/symlink-manager/go/internal/tools"
	"github.com/IgorAssuncao/symlink-manager/go/pkg/filesystem"
)

var wg sync.WaitGroup

func gracefulShutdown(ticker time.Ticker) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	ticker.Stop()

	log.Println("Gracefully Shutting Down...")
	defer wg.Done()
}

func main() {
	c := config.NewConfig("config.yaml")

	ticker := time.NewTicker(5 * time.Second)

	wg.Go(func() {
		tools.ProcessTools(c.Tools)

		counter := 0
		for t := range ticker.C {
			counter++
			log.Println("Tick at", t)
			if fileModifiedInLastMinute := filesystem.GetFileLastModified("config.yaml").After(time.Now().Add(-1 * time.Minute)); fileModifiedInLastMinute {
				tools.ProcessTools(c.Tools)
			}
			if counter >= 3 {
				ticker.Stop()
				break
			}
		}

		log.Println("Stopped")
	})

	go gracefulShutdown(*ticker)

	wg.Wait()
}

// TODO: Improve logs.

// TODO: v2: Make it run as a daemon (service).
// By having an infinite loop or a go ticker that works as a timer
// checking every minute or a custom time.

// TODO: v2.5: read config file path from cli flag.

// TODO: v3: instead of checking all of them, diff the config file and
// change only the key that was changed.
