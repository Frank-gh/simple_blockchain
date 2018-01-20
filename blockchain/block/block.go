package block

import (
	"strconv"
	"time"
)

type Block struct {
	Index        int64
	PreviousHash string
	Timestamp    int64
	Data         string
	Hash         string
	Nonce        int64
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

func (this *Block) DumpBlock() string {
	ret := "------------------------------------------------------------\n"
	ret += "|Index				|" + strconv.FormatInt(this.Index, 10) + "\n"
	ret += "|PreviousHash		|" + this.PreviousHash + "\n"
	ret += "|Timestamp			|" + strconv.FormatInt(this.Timestamp, 10) + "\n"
	ret += "|Data				|" + this.Data + "\n"
	ret += "|Hash				|" + this.Hash + "\n"
	ret += "|Nonce				|" + strconv.FormatInt(this.Nonce, 10) + "\n"
	return ret
}
