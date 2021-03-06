package ip

import (
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

//	http://www.ipip.net/download.html

var buffer []byte
var ip_offset uint32
var ip_index []byte

func Load(ipBinaryFilePath string) bool {
	var err error
	if buffer, err = ioutil.ReadFile(ipBinaryFilePath); err == nil {
		return true
	}
	return false
}

func Find(ip string) (re struct {
	country  string
	province string
	city     string
}) {

	ipdot := strings.Split(ip, ".")
	pre_ip, _ := strconv.Atoi(ipdot[0])
	tmp_offset := pre_ip * 4

	ip_offset = bytesBigEndianToUint32(buffer[:4])
	ip_index = buffer[4:ip_offset]
	nip := ip2long(ip)

	start_len := bytesLittleEndianToUint32(ip_index[tmp_offset : tmp_offset+4])

	var index_offset uint32 = 0
	var index_length uint32 = 0

	for start := start_len*8 + 1024; start < ip_offset-1028; start += 8 {
		if bytesBigEndianToUint32(ip_index[start:start+4]) >= nip {
			index_length = 0xFF & uint32(ip_index[start+7])
			tmp := ip_index[start+4 : start+7]
			tmp = append(tmp, 0x0)
			index_offset = bytesLittleEndianToUint32(tmp)
			break
		}
	}
	if index_offset != 0 {
		pos := index_offset + ip_offset - 1024
		var tab = strings.Split(string(ip_index[pos-4:pos+index_length-4]), "\t")
		if len(tab) > 2 {
			re.country = tab[0]
			re.province = tab[1]
			re.city = tab[2]
		}
	}
	return re
}

//binary.BigEndian.Uint32(b)
func bytesBigEndianToUint32(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func bytesLittleEndianToUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return bytesBigEndianToUint32(ip)
}
