package ip

import (
	"encoding/binary"
	"net"
)

func Uint32(s string) uint32 {
	if ip := net.ParseIP(s); ip != nil {
		return binary.BigEndian.Uint32(ip.To4())
	}
	return 0
}

func String(u32 uint32) string {
	ab := make([]byte, 4)
	binary.BigEndian.PutUint32(ab, u32)
	return net.IP(ab).String()
}
