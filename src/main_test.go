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
//     "bou.ke/monkey"
//     "reflect"
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

//     type clock struct {
//         hour int
//         min int
//         sec int
//     }
//     clocks := []clock {
//         { 10, 10, 10 },
//         { 10, 10, 11 },
//         { 10, 10, 12 },
//         { 10, 10, 13 },
//         { 10, 10, 14 },
//         { 10, 10, 15 },
//         { 10, 10, 16 },
//         { 10, 10, 17 },
//         { 10, 10, 18 },
//         { 10, 10, 19 },
//         { 10, 10, 20 },
//         { 10, 10, 21 },
//         { 10, 10, 22 },
//         { 10, 10, 23 },
//         { 10, 10, 24 },
//         { 10, 10, 25 },
//         { 10, 10, 26 },
//         { 10, 10, 27 },
//         { 10, 10, 28 },
//         { 10, 10, 29 },
//         { 10, 10, 30 },
//     }
//     i := 0
//     var t1 time.Time
//     monkey.PatchInstanceMethod(reflect.TypeOf(t1), "Clock", func(t1 time.Time) (hour, min, sec int) {
//         hour = clocks[i].hour
//         min = clocks[i].min
//         sec = clocks[i].sec
//         time.Sleep(1000 * time.Millisecond)
//         i++
//         return
//     } )

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
