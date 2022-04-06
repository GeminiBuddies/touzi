package builtin

import (
	"strconv"
	"touzi"
	"touzi/random"
	"unicode"
	"unicode/utf8"
)

const docInt = `
Usage:

        d<arguments>

<arguments> can be one of the following:

	<max>
	<min>,<max>
	<min>,<max>,<step>
	i<bits> or u<bits>

The default <max> is 6. The default <min> and <step> is 1.
`

type Int struct {
	source random.Source
}

func (i *Int) Information() touzi.Information {
	return touzi.Information{
		Prefix:        'd',
		Name:          "int",
		Description:   "a touzi generating integers",
		Documentation: docInt,
	}
}

func (i *Int) InjectSource(source random.Source) {
	i.source = source
}

func (i *Int) rollIUBits(args []touzi.Argument) (rolled bool, result touzi.Result, err error) {
	if len(args) != 1 || args[0] == "" {
		return
	}

	c, s := utf8.DecodeRuneInString(string(args[0]))
	b := string(args[0][s:])

	if c = unicode.ToLower(c); c != 'i' && c != 'u' {
		return
	}

	rolled = true

	var bits uint64
	if bits, err = strconv.ParseUint(b, 10, 8); err != nil || bits > 64 {
		err = touzi.ErrorInvalidArgument(args, 0)
		return
	}

	shift := 64 - bits
	rand := i.source.Next()

	if c == 'i' {
		result = int64(rand) >> shift
	} else {
		result = rand >> shift
	}

	return
}

func (i *Int) rollMinMaxStep(args []touzi.Argument) (rolled bool, result touzi.Result, err error) {
	if len(args) > 3 {
		return
	}

	rolled = true

	var (
		start int64 = 6
		end   int64 = 1
		step  int64 = 1
		ustep uint64
	)

	for index, arg := range args {
		if number, ok := arg.AsInt(); !ok {
			err = touzi.ErrorInvalidArgument(args, index)
			return
		} else {
			switch index {
			case 0:
				start = number
			case 1:
				end = number
			case 2:
				step = number
			}
		}
	}

	if step == 0 {
		err = touzi.ErrorInvalidArgument(args, 2)
		return
	} else if step < 0 {
		step = -step
	}

	ustep = uint64(step)
	if ustep == 1 {
		if start > end {
			start, end = end, start
		}

		result = start + int64(random.Bounded(i.source, uint64(end)-uint64(start)+1))
	} else {
		if start <= end {
			result = start + int64(random.Bounded(i.source, (uint64(end)-uint64(start))/ustep+1)*ustep)
		} else {
			result = start - int64(random.Bounded(i.source, (uint64(start)-uint64(end))/ustep+1)*ustep)
		}
	}

	return
}

func (i *Int) Roll(args []touzi.Argument) (result touzi.Result, err error) {
	var rolled bool

	if rolled, result, err = i.rollIUBits(args); rolled {
		return
	}

	if rolled, result, err = i.rollMinMaxStep(args); rolled {
		return
	}

	err = touzi.ErrorInvalidArguments(args)
	return
}
