# simple_blockchain
> 一个简单的区块链,供学习研究
## 数据结构
### block结构
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
### blockchain结构
```go
type blockchain struct {
	CurBlock   *block.Block
	Blocks     []*block.Block
	Difficulty uint
	Locker     *sync.Mutex
}

```

### peer结构
```go
type peer struct {
	svrPeers map[string]*tcpnetwork.Connection
	cliPeers map[string]*tcpnetwork.Connection
	peerName string
}
```

```
go get -v github.com/Frank-gh/simple_blockchain
```
