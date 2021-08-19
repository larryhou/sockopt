// +build aix darwin dragonfly freebsd netbsd openbsd solaris windows

package sockopt

import (
    "fmt"
    "syscall"
)

func SetNoDelay(fd int, enabled int) error {
    fmt.Printf("%d:SetNoDelay enabled=%d\n", fd, enabled)
    return syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_NODELAY, enabled)
}

func SetQuickAck(fd int, enabled int) error {return nil}
func SetCork(fd int, enabled int) error {return nil}
