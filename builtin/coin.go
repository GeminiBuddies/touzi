package builtin

import (
	"touzi"
	"touzi/random"
)

const docCoin = `
Usage:

        c

`

type Coin struct {
	source random.Source
}

func (c *Coin) Information() touzi.Information {
	return touzi.Information{
		Prefix:        'c',
		Name:          "coin",
		Description:   "a coin-like touzi",
		Documentation: docCoin,
	}
}

func (c *Coin) InjectSource(source random.Source) {
	c.source = source
}

func (c *Coin) Roll(args []touzi.Argument, format string) (result touzi.Result, err error) {
	if len(args) > 0 {
		return "", touzi.ErrorInvalidArguments(args)
	}

	if c.source.Next()&1 == 0 {
		return "true", nil
	} else {
		return "false", nil
	}
}
