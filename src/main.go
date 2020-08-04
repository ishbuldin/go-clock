package main

import (
    "fmt"
    "time"
    "runtime"
    "os"
    "os/exec"
    "log"
)

func cleanTerminal() {
    cmd := exec.Command("clear")
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cls")
    }
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

func printTime() {
    _,_,previousSecond := time.Now().Clock()
    for {
        hour,minute,second := time.Now().Clock()
        time.Sleep(10 * time.Millisecond)
        if(second != previousSecond) {
            cleanTerminal()
            fmt.Printf("%d:%d:%02d \n", hour, minute, second)
            previousSecond = second
        }
    }
}

func main() {
    go printTime()
    var input string
    fmt.Scanln(&input)
}
