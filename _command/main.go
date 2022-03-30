package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"touzi"
	"touzi/builtin"
	"touzi/pcg"
)

func executor(w prompt.ConsoleWriter) func(string) {
	source := pcg.New()
	dispatcher := &touzi.Dispatcher{}
	intTouzi := &builtin.Int{}

	intTouzi.InjectSource(source)
	dispatcher.Register(intTouzi)

	return func(s string) {
		for index, r := range touzi.SplitRequests(s) {
			w.SetColor(prompt.DarkGray, prompt.DefaultColor, true)
			w.WriteStr(fmt.Sprintf("[%d] %s = ", index, r))

			result, err := dispatcher.Dispatch(r)

			if err == nil {
				w.SetColor(prompt.DefaultColor, prompt.DefaultColor, false)
				w.WriteStr(fmt.Sprintf("%s\n", result.Result))
			} else {
				w.SetColor(prompt.Red, prompt.DefaultColor, false)
				w.WriteStr(fmt.Sprintf("%v\n", err))
				w.SetColor(prompt.DefaultColor, prompt.DefaultColor, false)
			}
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
