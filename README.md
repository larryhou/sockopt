
```bash
# maxos
sudo dtrace -s setsockopt.d -c './client -addr 9.134.79.3'
```

```bash
# linux
strace -e trace=setsockopt -f ./server -quick-ack=true -no-delay=true
```