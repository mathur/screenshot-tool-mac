package main

import (
    "bufio"
    "bytes"
    "encoding/base64"
    "fmt"
    "github.com/atotto/clipboard"
    "github.com/howeyc/fsnotify"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
)

const (
    clientID     = "48084fee9cbbc92"
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

                        imgFile, err := os.Open(filename)
                        if err != nil {
                            fmt.Println(err)
                            os.Exit(1)
                        }
                        defer imgFile.Close()
                        fInfo, _ := imgFile.Stat()
                        var size int64 = fInfo.Size()
                        buf := make([]byte, size)
                        fReader := bufio.NewReader(imgFile)
                        fReader.Read(buf)
                        imgBase64Str := base64.StdEncoding.EncodeToString(buf)
                        fmt.Println(imgBase64Str)

                        url := "https://api.imgur.com/3/image"
                        var jsonStr = []byte(`{"image":"` + imgBase64Str + `"}`)
                        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
                        req.Header.Set("Authorization", "Client-ID 48084fee9cbbc92")
                        req.Header.Set("Content-Type", "application/json")

                        client := &http.Client{}
                        resp, err := client.Do(req)
                        if err != nil {
                            panic(err)
                        }
                        defer resp.Body.Close()

                        body, _ := ioutil.ReadAll(resp.Body)
                        fmt.Println("response Body:", string(body))

                        err = clipboard.WriteAll(filename)
                        os.Remove(filename)
                    }
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