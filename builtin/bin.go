package builtin

import (
	"touzi"
	"touzi/random"
)

const docBin = `
Usage:

        b<bytes>

`

type Bin struct {
	source random.Source
}

func (b *Bin) Information() touzi.Information {
	return touzi.Information{
		Prefix:        'b',
		Name:          "bin",
		Description:   "a touzi generating binary data",
		Documentation: docBin,
	}
}

func (b *Bin) InjectSource(source random.Source) {
	b.source = source
}

func (b *Bin) Roll(args []touzi.Argument) (result touzi.Result, err error) {
	if len(args) != 1 {
		return nil, touzi.ErrorInvalidArguments(args)
	}

	if by, ok := args[0].AsUint(); !ok {
		return nil, touzi.ErrorInvalidArgument(args, 0)
	} else {
		r := make([]byte, by, by)
		random.Fill(b.source, r)

		return r, nil
	}
}
