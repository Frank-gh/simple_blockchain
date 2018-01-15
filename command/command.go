package command

import (
	"github.com/Frank-gh/simple_blockchain/blockchain"
)

type command struct{}

var Comm *command

func (this *command) Help() string {
	help := "===  Welcome to Simple Blockchain ! ===\n"
	help += "  Command:\n"
	help += "    help [command...]\t\tProvides help for a given command.\n"
	help += "    exit             \t\tExits application.\n"
	return help
}

func (this *command) Mine(data string) string {
	newblock := blockchain.BlockChain.GenerateNextBlock(data)
	if err := blockchain.BlockChain.IsValidNewBlock(newblock, blockchain.BlockChain.CurBlock); err != nil {
		return err.Error()
	}
	blockchain.BlockChain.AddBlock(newblock)
	return newblock.Hash
}
func init() {
	Comm = new(command)
}
