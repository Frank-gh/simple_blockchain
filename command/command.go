package command

import (
	"fmt"
	"os"
	"strings"
	"time"

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

func (this *command) Open(host, port string) string {
	return p2p.OpenPort(host, port)
}

func (this *command) Connect(host, port string) string {
	return p2p.Connect(host, port)
}

func init() {
	Comm = new(command)
}

func ProgressBar(quit chan bool) {
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
		fmt.Printf("\r%.0f%%[%s]", float64(i)/49*100, h)
		os.Stdout.Sync()
	}
	<-quit
}
