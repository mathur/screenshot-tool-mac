package main

import (
    "fmt"
    "github.com/atotto/clipboard"
    "github.com/howeyc/fsnotify"
    "log"
    "os"
    "strings"
    //"time"
)

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan bool)

    go func() {
        for {
            select {
            case ev := <-watcher.Event:
                if ev.IsCreate() {
                    strEvent := fmt.Sprintf("%s", ev)
                    endingIndex := strings.Index(strEvent, "\":")
                    filename := strEvent[1:endingIndex]
                    beginningIndex := strings.LastIndex(filename, "/") + 1

                    if string(filename[beginningIndex]) != "." {
                        log.Println("Detected " + filename)

                        // copy url to clipboard (TODO actually make it the URL not filename)
                        err = clipboard.WriteAll(filename)
                        
                        // delete the file
                        os.Remove(filename)
                    }

                   //time.Sleep(1000 * time.Millisecond)
                }
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch("/Users/rohan/Screenshots/")
    if err != nil {
        log.Fatal(err)
    }

    <-done
    watcher.Close()
}