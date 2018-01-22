package repl

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/Frank-gh/go-cmd-repl"
	"github.com/Frank-gh/go-cmd-repl/repl"
	"github.com/Frank-gh/simple_blockchain/command"
)

func echo(args ...cmd_repl.Argument) interface{} {
	echo := make([]string, len(args))
	for i, arg := range args {
		s, _ := arg.AsString()
		echo[i] = strings.Trim(s, "\n")
	}
	return strings.Join(echo, " ")
}

// client exit
func exit(args ...cmd_repl.Argument) interface{} {
	// Do sth.
	fmt.Println("Simple Blockchain exit !")
	os.Exit(0)
	return nil
}

func help(args ...cmd_repl.Argument) interface{} {
	return command.Comm.Help()
}

func mine(args ...cmd_repl.Argument) interface{} {
	param := make([]string, len(args))
	for i, arg := range args {
		s, _ := arg.AsString()
		param[i] = strings.Trim(s, "\n")
	}
	str := strings.Join(param, " ")
	if str == "" {
		return command.Comm.Help()
	}
	return command.Comm.Mine(str)
}

func open(args ...cmd_repl.Argument) interface{} {
	param := make([]string, len(args))
	for i, arg := range args {
		s, _ := arg.AsString()
		param[i] = strings.Trim(s, "\n")
	}

	if len(param) != 2 {
		return command.Comm.Help()
	}
	return command.Comm.Open(param[0], param[1])
}

func connect(args ...cmd_repl.Argument) interface{} {
	param := make([]string, len(args))
	for i, arg := range args {
		s, _ := arg.AsString()
		param[i] = strings.Trim(s, "\n")
	}

	if len(param) != 2 {
		return command.Comm.Help()
	}
	return command.Comm.Connect(param[0], param[1])
}

func StartRepl() {
	r := repl.New()
	r.RegisterCommand("echo", echo)
	r.RegisterCommand("exit", exit)
	r.RegisterCommand("help", help)
	r.RegisterCommand("mine", mine)
	r.RegisterCommand("open", open)
	r.RegisterCommand("connect", connect)
	go r.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Kill, os.Interrupt)
	<-quit
}
