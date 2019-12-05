package main

import (
    "flag"
    "fmt"
    "net"
    "os"
    "port-scanner/table"
    "strconv"
    "sync"
)

const (
    maxPort int = 65535
    minPort int = 1
)

var (
    help     bool   // help info
    protocol string // network protocol
    ip       string // target ip address
    port     int    // designated port
)

func init() {
    flag.BoolVar(&help, "h", false, "help information")
    flag.StringVar(&protocol, "p", "tcp", "network protocol,tcp,udp and so on")
    flag.StringVar(&ip, "ip", "", "target ip address")
    flag.IntVar(&port, "port", 0, "target tcp port,if unspecified,default 1-65535")
}

func main() {
    flag.Parse()
    if help {
        flag.Usage()
        os.Exit(0)
    }
    if ip == "" {
        fmt.Println("empty ip address")
        flag.Usage()
        os.Exit(0)
    }
    if port == 0 {
        allScan()
    } else {
        designatedScan()
    }
    fmt.Println("fetch complete")
}

func designatedScan() {
    wg := new(sync.WaitGroup)
    wg.Add(1)
    go Scan(wg, ip, port)
    wg.Wait()
}

// allScan() scan all port
func allScan() {
    wg := new(sync.WaitGroup)
    for p := minPort; p < maxPort; p++ {
        wg.Add(1)
        go Scan(wg, ip, p)
    }
    wg.Wait()
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
