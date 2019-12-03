package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "port-scanner/table"
    "strconv"
    "sync"
)

const (
    maxPort int = 65535
    minPort int = 1
)

var (
    protocol string // network protocol
    ip       string // target ip address
    port     int    // designated port
)

func init() {
    flag.StringVar(&protocol, "p", "tcp", "network protocol,tcp,udp and so on")
    flag.StringVar(&ip, "ip", "", "target ip address")
    flag.IntVar(&port, "port", 0, "target tcp port,if unspecified,default 1-65535")
}

func parseFlag() {
    flag.Parse()
    if ip == "" {
        log.Fatal("empty address")
    }
    // TODO protocol and port unused.
}

func main() {
    parseFlag()
    wg := new(sync.WaitGroup)
    for p := minPort; p < maxPort; p++ {
        wg.Add(1)
        go Scan(wg, ip, p)
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
