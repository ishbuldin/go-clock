package main

import (
    "testing"
    "os"
    "bytes"
    "time"
    "io"
    "strings"
    "log"
    "strconv"
)

type test struct {
    seconds int
    expectedSeconds int
    timeFormat string
}

var tests = []test {
    { 0, 0, "15:04:05" },
    { 1, 1, "15:04:05" },
    { 2, 2, "15:04:05" },
    { 3, 3, "15:04:05" },
}

func TestPrintTime(t *testing.T) {

    for i, data := range tests {
        log.Print("Start test: #", strconv.Itoa(i))

        old := os.Stdout
        r, w, _ := os.Pipe()
        os.Stdout = w

        PrintTime(time.Duration(data.seconds) * time.Second)

        c1 := make(chan string, 1)
        go func() {
            var buf bytes.Buffer
            io.Copy(&buf, r)
            c1 <- buf.String()
        }()
        w.Close()
        os.Stdout = old

        expectedSeconds := data.expectedSeconds
        select {
            case msg := <- c1:
                 r.Close()
                 output := strings.Fields(msg)
                 if len(output) == expectedSeconds {
                    log.Print("seconds:", len(output))
                 } else {
                    log.Print("seconds:", len(output))
                    t.Error(
                        "expectedSeconds", expectedSeconds,
                        "gotSeconds", len(output),
                    )
                 }
                 timeFormat := data.timeFormat
                 for _, v := range output {
                    log.Print(v[11:])
                    _, e := time.Parse(timeFormat, v[11:])
                    if e != nil {
                        t.Error(e.Error())
                    }
                 }
        }
    }

}