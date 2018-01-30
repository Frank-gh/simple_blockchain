package block

import (
	"strconv"
	"time"
)

// 区块的数据结构
type Block struct {
	Index        int64  // 块编号
	PreviousHash string // 前一个快的哈希
	Timestamp    int64  // 时间戳
	Data         string // 存储的数据
	Hash         string // 当前块的哈希
	Nonce        int64  // 随机数
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

// 初始化块（原始快）
func NewInitBlock() *Block {
	return &Block{
		Index:        0,
		PreviousHash: "0000000000000000000000000000000000000000000000000000000000000000",
		Timestamp:    time.Now().UTC().UnixNano(),
		Data:         "Shylock's Simple Blockchain!",
		Hash:         "000006030393fbbf48010544a74edca75ecd36bde25a198a623ce13ea20e29e4",
		Nonce:        7365,
	}
}

// 打印块信息
func (this *Block) DumpBlock() string {
	ret := "------------------------------------------------------------\n"
	ret += "|Index				|" + strconv.FormatInt(this.Index, 10) + "\n"
	ret += "|PreviousHash			|" + this.PreviousHash + "\n"
	ret += "|Timestamp			|" + strconv.FormatInt(this.Timestamp, 10) + "\n"
	ret += "|Data				|" + this.Data + "\n"
	ret += "|Hash				|" + this.Hash + "\n"
	ret += "|Nonce				|" + strconv.FormatInt(this.Nonce, 10) + "\n"
	return ret
}
