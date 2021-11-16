package ft

import "github.com/cespare/xxhash"

func hash(text []byte, salt []byte) uint64 {
	b := make([]byte, len(text)+len(salt))
	b = append(b, text...)
	b = append(b, salt...)
	return xxhash.Sum64(b)
}
