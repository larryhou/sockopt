package main

import (
    "flag"
    "fmt"
    "github.com/larryhou/sockopt"
    "log"
    "net"
    "syscall"
    "unsafe"
)

func main() {
    var quickAck, noDelay bool
    var addr string
    var port int
    flag.BoolVar(&quickAck, "quick-ack", false, "TCP_QUICKACK sockopt")
    flag.BoolVar(&noDelay, "no-delay", false, "TCP_NODELAY sockopt")
    flag.StringVar(&addr, "addr", "localhost", "server address")
    flag.IntVar(&port, "port", 2121, "server port")
    flag.Parse()

    dialer := &net.Dialer{Control: func(network, address string, c syscall.RawConn) error {
        return c.Control(func(fd uintptr) {
            if err := sockopt.SetNoDelay(int(fd), int(*(*byte)(unsafe.Pointer(&noDelay)))); err != nil {
                log.Printf("%d:SetNoDelay err: %v", fd, err)
            }
            if err := sockopt.SetQuickAck(int(fd), int(*(*byte)(unsafe.Pointer(&quickAck)))); err != nil {
                log.Printf("%d:SetQuickAck err: %v", fd, err)
            }
        })
    }}

    if c, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", addr, port)); err != nil {panic(err)} else {
        if t, ok := c.(*net.TCPConn); ok {
            t.SetNoDelay(false)
        }
        buf := make([]byte, 16<<10)
        stream := &sockopt.Stream{Rwp: c}
        defer stream.Close()
        for {
            c.Write([]byte{'x'})
            c.Write([]byte{'y'})
            c.Write([]byte{'z'})
            if s, err := stream.ReadString(buf); err != nil {panic(err)} else {
                fmt.Printf(">> %s", s)
            }
        }
    }
}
