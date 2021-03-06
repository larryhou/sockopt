package sockopt

import (
    "fmt"
    "syscall"
)

func PrintSockopts(fd int) {
    if v, err := syscall.GetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF); err == nil {fmt.Printf("# %d::RCVBUF=%.1fKB\n", fd, float64(v)/1024)}
    if v, err := syscall.GetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_SNDBUF); err == nil {fmt.Printf("# %d::SNDBUF=%.1fKB\n", fd, float64(v)/1024)}
    if v, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, 0x001); err == nil {fmt.Printf("# %d::TCP_NODELAY=%X\n", fd, v)}
    if v, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, 0x103); err == nil {fmt.Printf("# %d::TCP_SENDMOREACKS=%X\n", fd, v)}
    if v, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, 0x004); err == nil {fmt.Printf("# %d::TCP_NOPUSH=%X\n", fd, v)}
    if v, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, 0x008); err == nil {fmt.Printf("# %d::TCP_NOOPT=%X\n", fd, v)}
    if v, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, 0x00C); err == nil {fmt.Printf("# %d::TCP_QUICKACK=%X\n", fd, v)}
    if v, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, 0x003); err == nil {fmt.Printf("# %d::TCP_CORK=%X\n", fd, v)}
}
