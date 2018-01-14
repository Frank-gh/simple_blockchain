package repl

import (
	"os"
	"os/signal"
	"strings"

	"github.com/garslo/go-cmd-repl"
	"github.com/garslo/go-cmd-repl/repl"
)

func Echo(args ...cmd_repl.Argument) interface{} {
	echo := make([]string, len(args))
	for i, arg := range args {
		s, _ := arg.AsString()
		echo[i] = strings.Trim(s, "\n")
	}
	return strings.Join(echo, " ")
}

func StartRepl() {
	r := repl.New()
	r.RegisterCommand("echo", Echo)
	go r.Start()

	// Do other things

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Kill, os.Interrupt)
	<-quit
}
