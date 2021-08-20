package main

import (
    "bytes"
    "encoding/binary"
    "flag"
    "fmt"
    "github.com/larryhou/sockopt"
    "io"
    "log"
    "net"
    "time"
    "unsafe"
)

func main() {
    var quickAck, noDelay bool
    var port int
    flag.BoolVar(&quickAck, "quick-ack", false, "TCP_QUICKACK sockopt")
    flag.BoolVar(&noDelay, "no-delay", false, "TCP_NODELAY sockopt")
    flag.IntVar(&port, "port", 2121, "server listen port")
    flag.Parse()

    if s, err := net.Listen("tcp", fmt.Sprintf(":%d", port)); err != nil {panic(err)} else {
        for {
            if c, err := s.Accept(); err == nil {
                if t, ok := c.(*net.TCPConn); ok {
                    if r, err := t.SyscallConn(); err == nil {
                        r.Control(func(fd uintptr) {
                            log.Printf("%02d %s\n", fd, c.RemoteAddr())
                            sockopt.PrintSockopts(int(fd))
                            if err := sockopt.SetNoDelay(int(fd), int(*(*byte)(unsafe.Pointer(&noDelay)))); err != nil {
                                log.Printf("%d:SetNoDelay err: %v", fd, err)
                            }
                            if err := sockopt.SetQuickAck(int(fd), int(*(*byte)(unsafe.Pointer(&quickAck)))); err != nil {
                                log.Printf("%d:SetQuickAck err: %v", fd, err)
                            }
                            sockopt.PrintSockopts(int(fd))
                        })
                    }
                }

                go func() {
                    defer c.Close()
                    buf := &bytes.Buffer{}
                    num := 0
                    t := time.NewTicker(time.Second)
                    defer t.Stop()
                    for range t.C {
                        buf.Reset()
                        buf.WriteByte(0)
                        buf.WriteByte(0)
                        buf.WriteString(fmt.Sprintf("%4d", num))
                        buf.WriteByte(' ')
                        buf.WriteString(time.Now().String())
                        buf.WriteByte(' ')
                        buf.WriteString("server short message\n")
                        binary.BigEndian.PutUint16(buf.Bytes(), uint16(buf.Len()-2))
                        if _, err := io.Copy(c, buf); err != nil {return}
                        num++
                    }
                }()

                go func() {
                    buf := make([]byte, 1024)
                    for { if _, err := c.Read(buf); err != nil {return} }
                }()
            }
        }
    }
}

