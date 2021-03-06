package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/Frank-gh/simple_blockchain/blockchain"
	"github.com/Frank-gh/simple_blockchain/blockchain/block"
	"github.com/Frank-gh/tcpnetwork"
	"github.com/golang/glog"
)

var handlers = make(map[string]FuncHandler)

func AddHandler(name string, f FuncHandler) {
	handlers[name] = f
}
func handleInput(data []byte, conn *tcpnetwork.Connection) {
	pkg := &json_pkg{}
	err := json.Unmarshal(data, pkg)
	if err != nil {
		glog.Error(err.Error())
	}
	handlers[pkg.Type](pkg.Data, conn)
}
func broadCastBlock() {
	for _, cli := range Peer.cliPeers {
		sendCurBlock(cli)
	}
	for _, svr := range Peer.svrPeers {
		sendCurBlock(svr)
	}
}

func broadCastIndex() {
	for _, cli := range Peer.cliPeers {
		sendIndex(blockchain.BlockChain.Index(), cli)
	}
	for _, svr := range Peer.svrPeers {
		sendIndex(blockchain.BlockChain.Index(), svr)
	}
}

func sendCurBlock(conn *tcpnetwork.Connection) {
	blockchain.BlockChain.Locker.Lock()
	defer func() {
		blockchain.BlockChain.Locker.Unlock()
	}()
	block := blockchain.BlockChain.CurBlock
	pkg_block := &block_pkg{
		Index:        block.Index,
		PreviousHash: block.PreviousHash,
		Timestamp:    block.Timestamp,
		Data:         block.Data,
		Hash:         block.Hash,
		Nonce:        block.Nonce,
	}
	data, err := json.Marshal(pkg_block)
	if err != nil {
		glog.Error(err.Error())
	}
	pkg := &json_pkg{
		Type:     "block",
		PeerName: Peer.peerName,
		Data:     data,
	}
	send_pkg, err := json.Marshal(pkg)
	if err != nil {
		glog.Error(err.Error())
	}
	conn.Send(send_pkg, 0)
}

func sendBlock(index int64, conn *tcpnetwork.Connection) {
	blockchain.BlockChain.Locker.Lock()
	defer func() {
		blockchain.BlockChain.Locker.Unlock()
	}()
	for idx, block := range blockchain.BlockChain.Blocks {
		if int64(idx) > index {
			pkg_block := &block_pkg{
				Index:        block.Index,
				PreviousHash: block.PreviousHash,
				Timestamp:    block.Timestamp,
				Data:         block.Data,
				Hash:         block.Hash,
				Nonce:        block.Nonce,
			}
			data, err := json.Marshal(pkg_block)
			if err != nil {
				glog.Error(err.Error())
			}
			pkg := &json_pkg{
				Type:     "block",
				PeerName: Peer.peerName,
				Data:     data,
			}
			send_pkg, err := json.Marshal(pkg)
			if err != nil {
				glog.Error(err.Error())
			}
			conn.Send(send_pkg, 0)
		}
	}

}

func sendIndex(index int64, conn *tcpnetwork.Connection) {
	pkg_index := &index_pkg{
		Index: index,
	}
	data, err := json.Marshal(pkg_index)
	if err != nil {
		glog.Error(err.Error())
	}
	pkg := &json_pkg{
		Type:     "index",
		PeerName: Peer.peerName,
		Data:     data,
	}
	send_pkg, err := json.Marshal(pkg)
	if err != nil {
		glog.Error(err.Error())
	}
	conn.Send(send_pkg, 0)
}

func onIndex(data []byte, conn *tcpnetwork.Connection) {
	pkg := &index_pkg{}
	err := json.Unmarshal(data, pkg)
	if err != nil {
		glog.Error(err.Error())
	}
	glog.Info("Index = ", pkg.Index)
	if blockchain.BlockChain.Index() <= pkg.Index {
		// 本地peer高度小于等于远端peer高度，不做处理
		return
	}
	sendBlock(pkg.Index, conn)
}

func onBlock(data []byte, conn *tcpnetwork.Connection) {
	pkg := &block_pkg{}
	err := json.Unmarshal(data, pkg)
	if err != nil {
		glog.Error(err.Error())
	}

	blk := block.NewBlock(pkg.Index, pkg.Timestamp, pkg.Nonce, pkg.PreviousHash, pkg.Data, pkg.Hash)

	if err := blockchain.BlockChain.AddBlock(blk); err != nil {
		glog.Error(err.Error())
		sendIndex(int64(0), conn)
		return
	}
	fmt.Println()

	// dump recv Block
	fmt.Println(blk.DumpBlock())
}
