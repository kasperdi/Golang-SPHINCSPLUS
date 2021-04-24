package util

import "encoding/binary"
/* import "bytes" */
import "math"
/* import "fmt" */

// For x and y non-negative intergers, toByte(x,y) returns the y-byte bytearray 
// containing the binary representation of x in big-endian byte-order.
func ToByte(x uint32, y uint) []byte {
	buffer := make([]byte, y)
	binary.BigEndian.PutUint32(buffer, x)
	return buffer
}


func ToByte2(in int, outlen int) []byte {
	out := make([]byte, outlen)
	for i := outlen - 1; i >= 0; i-- {
        out[i] = byte(in & 0xff);
        in = in >> 8;
    }
	return out
}

func BytesToUint64(in []byte) uint64 {
	res := uint64(0)

	for i := 0; i < len(in); i++ {
		res = res | (uint64(in[i]) << (8*(len(in) - 1 - i)))
	}
	return res;
}

func BytesToUint32(in []byte) uint32 {
	res := uint32(0)

	for i := 0; i < len(in); i++ {
		res = res | (uint32(in[i]) << (8*(len(in) - 1 - i)))
	}
	return res;
}


func Base_w(X []byte, w int, out_len int) []int {
	in := 0
	out := 0
	total := 0 // Use uint here instead?
	bits := 0
	basew := make([]int, out_len)
	for consumed := 0; consumed < out_len; consumed++ {
		if(bits == 0) {
			
			total = int(X[in])			// Use Uint32 here instead?????
			in++
			bits += 8
		}
		bits -= int(math.Log2(float64(w)))
		basew[out] = (total >> bits) & (w - 1)
		out++
	}
	return basew
}

