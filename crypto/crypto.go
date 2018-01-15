package crypto

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

// 计算hash值
func CalcHash(index, timestamp, nonce int64, previousHash, data string) string {
	str := strconv.FormatInt(index, 10)
	str += previousHash
	str += strconv.FormatInt(timestamp, 10)
	str += data
	str += strconv.FormatInt(nonce, 10)

	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
