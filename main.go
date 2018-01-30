package main

import (
	"flag"

	"github.com/Frank-gh/simple_blockchain/repl"
)

func main() {
	flag.Parse()
	repl.StartRepl() // 启动交互式控制台
}
