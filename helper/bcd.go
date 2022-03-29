package helper

func Bcd2Str(bytes []byte) string {

	temp := make([]byte, len(bytes)*2, len(bytes)*2)

	for key, b := range bytes {
		val := ((b & 0xf0) >> 4) & 0x0f
		if val > 9 {
			temp[key*2] = val + 'A' - 10
		} else {
			temp[key*2] = val + '0'
		}

		val = b & 0x0f

		if val > 9 {
			temp[key*2+1] = val + 'A' - 10
		} else {
			temp[key*2+1] = val + '0'
		}
	}

	return string(temp)
}
func Bcd2Bytes(bcd []byte) []byte {
	bcdLen := len(bcd)
	temp := make([]byte, bcdLen/2, bcdLen/2)
	j := 0
	for i := 0; i < (bcdLen+1)/2; i++ {
		temp[i] = byte2bcd(bcd[j])
		j++
		if j >= bcdLen {
			temp[i] = 0x00 + (temp[i] << 4)
		} else {
			temp[i] = byte2bcd(bcd[j]) + +(temp[i] << 4)
			j++
		}
	}
	return temp
}
func byte2bcd(b uint8) uint8 {
	if b >= '0' && b <= '9' {
		b = b - '0'
	} else if b >= 'A' && b <= 'F' {
		b = b - 'A' + 10
	} else if b >= 'a' && b <= 'f' {
		b = b - 'a' + 10
	} else {
		b = b - 48
	}
	return b
}
