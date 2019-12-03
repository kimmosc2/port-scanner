package main

import (
    "fmt"
    "net"
    "port-scanner/table"
    "strconv"
    "sync"
)

const (
    maxPort int = 65535
    minPort int = 1
)


func main() {
    wg := new(sync.WaitGroup)
    for p := minPort; p < maxPort; p++ {
        wg.Add(1)
        go Scan(wg, "127.0.0.1", p)
    }
    wg.Wait()
    fmt.Println("fetch complete")
}

func Scan(wg *sync.WaitGroup, targetAddress string, port int) {
    defer wg.Done()
    _, err := net.Dial("tcp", targetAddress+":"+strconv.Itoa(port))
    if err == nil {
        result := table.GetPossibility(port)
        if result != "" {
            fmt.Printf("live port:[%5d]\t\tpossible: [%s]\n", port, result)
        } else {
            fmt.Printf("live port:[%5d]\n", port)
        }
    }
}
