package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/kube-prompt/kube"
)

var (
	version  string
	revision string
)

func main() {
	defer fmt.Println("Goodbye!")
	p := prompt.New(
		kube.Executor,
		kube.Completer,
		prompt.OptionTitle("kube-prompt: interactive kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
