package block

import (
	"time"
)

type Block struct {
	Index        int64  `json:index`
	PreviousHash string `json:previousHash`
	Timestamp    int64  `json:timestamp`
	Data         string `json:data`
	Hash         string `json:hash`
	Nonce        int64  `json:nonce`
}

// 创建新块
func NewBlock(index, timestamp, nonce int64, previousHash, data, hash string) *Block {
	return &Block{
		Index:        index,
		PreviousHash: previousHash,
		Timestamp:    timestamp,
		Data:         data,
		Hash:         hash,
		Nonce:        nonce,
	}
}

// 初始化块
func NewInitBlock() *Block {
	return &Block{
		Index:        0,
		PreviousHash: "0",
		Timestamp:    time.Now().UTC().UnixNano(),
		Data:         "Shylock's Simple Blockchain!",
		Hash:         "",
		Nonce:        0,
	}
}
