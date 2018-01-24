package command

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Frank-gh/simple_blockchain/blockchain"
	"github.com/Frank-gh/simple_blockchain/p2p"
)

var (
	timer = time.NewTicker(100 * time.Millisecond)
)

func Help() string {
	help := "===  Welcome to Simple Blockchain ! ===\n"
	help += "  Command:\n"
	help += "    help       [command...]  \t\tProvides help for a given command.\n"
	help += "    exit                     \t\tExits application.\n"
	help += "    mine       <data>        \t\tMine a new block. Eg: mine 50$.\n"
	help += "    open       <host> <port> \t\tOpen port to accept incoming connections. Eg: open localhost 7365.\n"
	help += "    connect    <host> <port> \t\tConnect to a new peer. Eg: connect localhost 7365.\n"
	help += "    blockchain <host> <port> \t\tSee the current state of the blockchain.\n"
	return help
}

func Mine(data string) string {
	quit := make(chan bool)
	go ProgressBar(quit)
	newblock := blockchain.BlockChain.GenerateNextBlock(data)

	if err := blockchain.BlockChain.AddBlock(newblock); err != nil {
		return err.Error()
	}
	quit <- true
	fmt.Println()
	return blockchain.BlockChain.DumpBlockchain()
}

func Open(host, port string) string {
	return p2p.OpenPort(host, port)
}

func Connect(host, port string) string {
	return p2p.Connect(host, port)
}

func Blockchain() string {
	return blockchain.BlockChain.DumpBlockchain()
}

func ProgressBar(quit chan bool) {
	i := 0
	for {
		select {
		case <-quit:
			{
				return
			}
		case <-timer.C:
			{
				h := strings.Repeat(">", i)
				fmt.Printf("\r%s", h)
				os.Stdout.Sync()
				i++
			}
		}
	}
}
