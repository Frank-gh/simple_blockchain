package repl

import (
	"os"
	"os/signal"
	"strings"

	"github.com/Frank-gh/go-cmd-repl"
	"github.com/Frank-gh/go-cmd-repl/repl"
)

func echo(args ...cmd_repl.Argument) interface{} {
	echo := make([]string, len(args))
	for i, arg := range args {
		s, _ := arg.AsString()
		echo[i] = strings.Trim(s, "\n")
	}
	return strings.Join(echo, " ")
}

func exit(args ...cmd_repl.Argument) interface{} {
	os.Exit(0)
	return nil
}

func enter(args ...cmd_repl.Argument) interface{} {
	return " "
}

func StartRepl() {
	r := repl.New()
	r.RegisterCommand("echo", echo)
	r.RegisterCommand("exit", exit)
	r.RegisterCommand("", enter)
	go r.Start()

	// Do other things

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Kill, os.Interrupt)
	<-quit
}
