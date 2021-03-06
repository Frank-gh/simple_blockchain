package blockchain

import (
	"errors"
	"sync"
	"time"

	"github.com/Frank-gh/simple_blockchain/blockchain/block"
	"github.com/Frank-gh/simple_blockchain/crypto"
)

var BlockChain *blockchain

// 区块链结构
type blockchain struct {
	CurBlock   *block.Block   //当前最新的区块
	Blocks     []*block.Block //区块数组，用于存储连续的区块
	Difficulty uint           //难度系数
	Locker     *sync.Mutex    //锁
}

func NewBlockChain() *blockchain {
	return &blockchain{
		CurBlock: block.NewInitBlock(),
		//Blocks:     make(map[int64]*block.Block),
		Difficulty: 5,
		Locker:     new(sync.Mutex),
	}
}

func CalculateHashForBlock(blk block.Block) string {
	return crypto.CalcHash(blk.Index, blk.Timestamp, blk.Nonce, blk.PreviousHash, blk.Data)
}

func init() {
	BlockChain = NewBlockChain()
	BlockChain.Blocks = append(BlockChain.Blocks, BlockChain.CurBlock)
}

func (this *blockchain) GenerateNextBlock(data string) *block.Block {
	previousHash := this.CurBlock.Hash
	nextIndex := this.CurBlock.Index + 1
	nextTimestamp := time.Now().UTC().UnixNano()
	var nonce int64 = 0
	var nextHash string = ""
	for {
		if this.IsValidHashDifficulty(nextHash) {
			break
		}

		nonce++
		nextHash = crypto.CalcHash(nextIndex, nextTimestamp, nonce, previousHash, data)
	}
	return block.NewBlock(nextIndex, nextTimestamp, nonce, previousHash, data, nextHash)
}

func (this *blockchain) IsValidHashDifficulty(hash string) bool {
	var i uint = 0
	for _, ch := range hash {
		if ch != '0' {
			break
		}
		i++
	}
	return i == this.Difficulty
}

func (this *blockchain) IsValidNewBlock(newBlock, preBlock *block.Block) error {
	blockHash := CalculateHashForBlock(*newBlock)

	if preBlock.Index+1 != newBlock.Index {
		return errors.New("❌  new block has invalid index")
	} else if preBlock.Hash != newBlock.PreviousHash {
		return errors.New("❌  new block has invalid previous hash")
	} else if blockHash != newBlock.Hash {
		return errors.New("❌  invalid hash: " + blockHash + " " + newBlock.Hash)
	} else if !this.IsValidHashDifficulty(blockHash) {
		return errors.New("❌  invalid hash does not meet difficulty requirements")
	}
	return nil
}

func (this *blockchain) AddBlock(newBlock *block.Block) error {
	this.Locker.Lock()
	defer func() { this.Locker.Unlock() }()

	// 如果新快小于当前链的长度，则用新快覆盖当前链上的块
	if newBlock.Index <= this.Index() {
		this.CurBlock = this.Blocks[newBlock.Index-1]
		this.Blocks = this.Blocks[:newBlock.Index]
	}
	if err := this.IsValidNewBlock(newBlock, this.CurBlock); err != nil {
		return err
	}
	this.CurBlock = newBlock
	this.Blocks = append(this.Blocks, newBlock)

	return nil
}

func (this *blockchain) Index() int64 {
	return int64(len(this.Blocks) - 1)
}

func (this *blockchain) DumpBlockchain() string {
	var ret string
	for _, blk := range this.Blocks {
		ret += blk.DumpBlock() + "\n"
	}
	return ret
}
