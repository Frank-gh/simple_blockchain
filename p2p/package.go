package p2p

type json_pkg struct {
	Type     string `json:type`
	PeerName string `json:peerName`
	Data     []byte `json:data`
}

type index_pkg struct {
	Index int64 `json:index`
}

type block_pkg struct {
	Index        int64  `json:index`
	PreviousHash string `json:previousHash`
	Timestamp    int64  `json:timestamp`
	Data         string `json:data`
	Hash         string `json:hash`
	Nonce        int64  `json:nonce`
}
