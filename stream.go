package sockopt

import (
    "bytes"
    "encoding/binary"
    "io"
    "os"
)

type Stream struct {
    Rwp io.ReadWriter
}

func (s *Stream) Name() string {
    if f, ok := s.Rwp.(*os.File); ok {return f.Name()}
    return ""
}

func (s *Stream) ReadString(buf []byte) (string, error) {
    if err := s.Read(buf, 2); err != nil {return "", err}
    n := int(binary.BigEndian.Uint16(buf))
    if n < cap(buf) {
        if err := s.Read(buf, n); err != nil {return "", err}
        return string(buf[:n]), nil
    } else {
        b := &bytes.Buffer{}
        for t := 0; t < n; {
            num := cap(buf)
            if n - t < num {num = n - t}
            if err := s.Read(buf, num); err != nil {return "", err}
            if _, err := b.Write(buf[:num]); err != nil {return "", err}
            t += num
        }
        return b.String(), nil
    }
}

func (s *Stream) WriteString(buf []byte, v string) error {
    n := len(v)
    binary.BigEndian.PutUint16(buf, uint16(n))
    if err := s.Write(buf, 2); err != nil {return err}
    for t := 0; t < n; {
        num := cap(buf)
        if n - t < num {num = n - t}
        copy(buf, v[t:t+num])
        if err := s.Write(buf, num); err != nil {return err}
        t += num
    }
    return nil
}

func (s *Stream) Read(p []byte, n int) error {
    for t := 0; t < n; {
        if i, err := s.Rwp.Read(p[t:n]); err == nil {t+=i} else {
            if err == io.EOF && t+i==n {break}
            return err
        }
    }
    return nil
}

func (s *Stream) Write(p []byte, n int) error {
    for t := 0; t < n; {
        if i, err := s.Rwp.Write(p[t:n]); err != nil {return err} else {t+=i}
    }
    return nil
}

func (s *Stream) Close() error {
    if c, ok := s.Rwp.(io.Closer); ok { return c.Close() }
    return nil
}
