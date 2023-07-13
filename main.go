package main

import (
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var ui lorca.UI
	ui, err := lorca.New("https://baidu.com", "", 800, 600, "--disable-sync", "--disable-translate")
	if err != nil {
		log.Fatal(err)
	}
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ui.Done():
	case <-chSignal:
	}
	ui.Close()
}
