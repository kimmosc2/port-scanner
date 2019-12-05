package main

import (
    "flag"
    "fmt"
    "net"
    "os"
    "port-scanner/table"
    "strconv"
    "sync"
    "time"
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

type Pool struct {
    Queue chan struct{}
    sync.WaitGroup
}

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
    pool := new(Pool)
    pool.Add(1)
    go Scan(pool, ip, port)
    pool.Wait()
}

// allScan() scan all port
func allScan() {
    pool := new(Pool)
    pool.Queue = make(chan struct{}, 5000)
    for p := minPort; p < maxPort; p++ {
        pool.Queue <- struct{}{}
        pool.Add(1)
        go Scan(pool, ip, p)
    }
    pool.Wait()
}

func Scan(pool *Pool, targetAddress string, port int) {
    _, err := net.DialTimeout("tcp", targetAddress+":"+strconv.Itoa(port), time.Second*3)
    if err == nil {
        result := table.GetPossibility(port)
        if result != "" {
            fmt.Printf("live port:[%5d]\t\tpossible: [%s]\n", port, result)
        } else {
            fmt.Printf("live port:[%5d]\n", port)
        }
    }
    pool.Done()
    <-pool.Queue
}
