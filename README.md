# simple_blockchain
> 一个简单的区块链,供学习研究
## 区块定义
```go
type Block struct {
	Index        int64
	PreviousHash string
	Timestamp    int64
	Data         string
	Hash         string
	Nonce        int64
}
```
