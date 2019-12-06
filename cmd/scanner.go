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
    version  bool   // version
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
    flag.BoolVar(&version,"version",false,"show program version")
}

func main() {
    flag.Parse()
    // help information
    if help {
        flag.Usage()
        os.Exit(1)
    }
    // version information
    if version {
        fmt.Println("port-scanner version:1.0.191205")
        fmt.Println("Author:BuTn<https://github.com/kimmosc2>")
        os.Exit(1)
    }
    // ip not null
    if ip == "" {
        fmt.Println("empty ip address")
        flag.Usage()
        os.Exit(0)
    }

    if port == 0 {
        allScan()
    // single
    } else {
        designatedScan()
    }
    fmt.Println()
    fmt.Println("fetch complete")
}

// single scan
func designatedScan() {
    pool := new(Pool)
    pool.Add(1)
    go Scan(pool, ip, port)
    pool.Wait()
}

// allScan() scan all port,allScan use channel and
// mutex build a pool to limit goroutine number, more
// goroutine will set up more connect, the system can't
// set up socket,this will effect the result.
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
