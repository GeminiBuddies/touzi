package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"touzi"
)

func executor(w prompt.ConsoleWriter) func(string) {
	return func(s string) {
		for index, r := range touzi.SplitRequests(s) {
			w.SetColor(prompt.DarkGray, prompt.DefaultColor, true)
			w.WriteStr(fmt.Sprintf("[%d] %s = ", index, r))
			w.SetColor(prompt.DefaultColor, prompt.DefaultColor, false)
			w.WriteStr(fmt.Sprintf("%s\n", s))
		}

		_ = w.Flush()
	}
}

func completer() func(prompt.Document) []prompt.Suggest {
	return func(document prompt.Document) []prompt.Suggest {
		return nil
	}
}

func main() {
	w := prompt.NewStdoutWriter()

	prompt.New(
		executor(w),
		completer(),
		prompt.OptionPrefix(">>> "),
	).Run()
}
