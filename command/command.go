package command

import (
	"github.com/Frank-gh/simple_blockchain/blockchain"
	"github.com/Frank-gh/simple_blockchain/p2p"
)

type command struct{}

var Comm *command

func (this *command) Help() string {
	help := "===  Welcome to Simple Blockchain ! ===\n"
	help += "  Command:\n"
	help += "    help    [command...]  \t\tProvides help for a given command.\n"
	help += "    exit                  \t\tExits application.\n"
	help += "    mine    <data>        \t\tMine a new block. Eg: mine 50$.\n"
	help += "    open    <host> <port> \t\tOpen port to accept incoming connections. Eg: open localhost 7365.\n"
	help += "    connect <host> <port> \t\tConnect to a new peer. Eg: connect localhost 7365.\n"
	return help
}

func (this *command) Mine(data string) string {
	newblock := blockchain.BlockChain.GenerateNextBlock(data)

	if err := blockchain.BlockChain.AddBlock(newblock); err != nil {
		return err.Error()
	}
	return blockchain.BlockChain.DumpBlockchain()
}

func (this *command) Open(host, port string) string {
	return p2p.OpenPort(host, port)
}

func (this *command) Connect(host, port string) string {
	return p2p.Connect(host, port)
}

func init() {
	Comm = new(command)
}
