package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"strings"
	"touzi"
	"touzi/builtin"
	"touzi/pcg"
	"unsafe"
)

func render(w prompt.ConsoleWriter, index int, roll touzi.Roll, err error) {
	// get error type
	_, prefixError := err.(*touzi.ErrTouziNotFound)
	_, argumentsError := err.(*touzi.ErrInvalidArguments)
	_, argumentError := err.(*touzi.ErrInvalidArgument)

	// print index
	if err != nil {
		w.SetColor(prompt.Red, prompt.DefaultColor, true)
	} else {
		w.SetColor(prompt.DarkGray, prompt.DefaultColor, true)
	}
	w.WriteStr(fmt.Sprintf("[%d] ", index))

	// print prefix
	if prefixError {
		w.SetColor(prompt.Red, prompt.DefaultColor, true)
	} else {
		w.SetColor(prompt.DarkGreen, prompt.DefaultColor, false)
	}
	w.WriteStr(string(roll.Prefix))

	// print arguments
	if len(roll.Arguments) > 0 {
		if argumentsError {
			w.SetColor(prompt.Red, prompt.DefaultColor, true)
			w.WriteStr(strings.Join(*(*[]string)(unsafe.Pointer(&roll.Arguments)), ","))
		} else {
			invalidArgumentAt := -1
			if argumentError {
				invalidArgumentAt = err.(*touzi.ErrInvalidArgument).Position
			}

			for argIndex, arg := range roll.Arguments {
				if argIndex > 0 {
					w.SetColor(prompt.DarkGray, prompt.DefaultColor, false)
					w.WriteStr(",")
				}

				if argIndex == invalidArgumentAt {
					w.SetColor(prompt.Red, prompt.DefaultColor, true)
				} else {
					w.SetColor(prompt.DefaultColor, prompt.DefaultColor, false)
				}
				w.WriteStr(string(arg))
			}
		}
	}

	// print format
	if roll.Format != "" {
		w.SetColor(prompt.DarkGray, prompt.DefaultColor, false)
		w.WriteStr("#")
		w.SetColor(prompt.DefaultColor, prompt.DefaultColor, false)
		w.WriteStr(roll.Format)
	}

	// print result or error message
	if err == nil {
		w.SetColor(prompt.DarkGray, prompt.DefaultColor, false)
		w.WriteStr(" = ")
		w.SetColor(prompt.DefaultColor, prompt.DefaultColor, false)
		w.WriteStr(fmt.Sprintf("%s\n", roll.Result))
	} else {
		w.SetColor(prompt.Red, prompt.DefaultColor, false)
		w.WriteStr(fmt.Sprintf(" : %v\n", err))
	}

	w.SetColor(prompt.DefaultColor, prompt.DefaultColor, false)
}

func executor(w prompt.ConsoleWriter) func(string) {
	cup := touzi.NewCup(pcg.New())

	cup.Add(&builtin.Int{})
	cup.Add(&builtin.Coin{})

	return func(s string) {
		for index, r := range touzi.SplitRequests(s) {
			roll, err := cup.RollOne(r)

			render(w, index, roll, err)
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
		// prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
			return "[0] ", true
		}),
	).Run()
}
