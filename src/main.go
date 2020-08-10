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

func PrintTime(duration time.Duration) {
    quit := make(chan bool)
    go func() {
        _,_,previousSecond := time.Now().Clock()
        for {
            select {
                case <- quit:
                    return
                default:
                    hour,minute,second := time.Now().Clock()
                    time.Sleep(10 * time.Millisecond)
                    if(second != previousSecond) {
                        cleanTerminal()
                        fmt.Printf("%02d:%02d:%02d\n", hour, minute, second)
                        previousSecond = second
                    }
            }
        }
    }()
    time.Sleep(duration)
    quit <- true
}

func main() {
    PrintTime(10 * time.Second)
}
