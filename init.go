package bitcointx

import (
	"fmt"
	"strings"
)

// HashCode 32 byte code
type HashCode [32]byte

func (hash *HashCode) compare(anotherHash *HashCode) int {
	hexStr := fmt.Sprintf("%x", hash)
	anotherHexStr := fmt.Sprintf("%x", anotherHash)
	return strings.Compare(hexStr, anotherHexStr)
}
