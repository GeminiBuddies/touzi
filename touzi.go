package touzi

import (
	"strconv"
	"touzi/random"
)

// Argument argument
type Argument string

func (p Argument) IsEmpty() bool {
	return p == ""
}

func (p Argument) AsInt() (value int64, ok bool) {
	if v, err := strconv.ParseInt(string(p), 0, 64); err == nil {
		return v, true
	} else {
		return 0, false
	}
}

func (p Argument) AsUint() (value uint64, ok bool) {
	if v, err := strconv.ParseUint(string(p), 0, 64); err == nil {
		return v, true
	} else {
		return 0, false
	}
}

func (p Argument) AsString() string {
	return string(p)
}

// Result result
type Result interface{}

// Information about a Touzi
type Information struct {
	Prefix        rune
	Name          string
	Description   string
	Documentation string
}

// Touzi a dice
type Touzi interface {
	Information() Information
	InjectSource(source random.Source)
	Roll(args []Argument) (result Result, err error)
}
