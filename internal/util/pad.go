package util

func PKCSpad(data []byte) (blocked []byte) {
	blocked_size := (1 + (len(data) / 16)) * 16
	diff := byte(blocked_size - len(data))
	blocked = make([]byte, blocked_size)

	for k, v := range data {
		blocked[k] = v
	}

	for i := len(data); i < blocked_size; i++ {
		blocked[i] = diff
	}

	return
}

func PKCSunpad(blocked []byte) (data []byte) {
	padlen := blocked[len(blocked)-1]
	oglen := len(blocked) - int(padlen)

	data = make([]byte, oglen)
	for i := range data {
		data[i] = blocked[i]
	}

	return
}
