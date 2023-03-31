package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	fileUpdated := make(chan string)

	go WatchEvents(fileUpdated)

	for f := range fileUpdated {
		fmt.Println("File updated:", f)
	}
}

func WatchEvents(updateF chan<- string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	path, _ := os.Getwd()
	filep := fmt.Sprintf("%s/text/text.txt", path)
	err = watcher.Add(filep)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("Error in file path:", err.Error())
	}

	for event := range watcher.Events {
		if event.Has(fsnotify.Write) {
			updateF <- event.Name
		}
	}
}
