package touzi

import (
	"unicode"
	"unicode/utf8"
)

type DispatchResult struct {
	Request   Request
	Prefix    rune
	Touzi     Touzi
	Arguments []Argument
	Format    string
	Result    Result
}

type Dispatcher struct {
	touzi map[rune]Touzi
}

func (d *Dispatcher) Register(touzi Touzi) {
	if d.touzi == nil {
		d.touzi = make(map[rune]Touzi)
	}

	prefix := touzi.Information().Prefix
	d.touzi[prefix] = touzi

	if lower := unicode.ToLower(prefix); lower != prefix {
		d.touzi[lower] = touzi
	}

	if upper := unicode.ToUpper(prefix); upper != prefix {
		d.touzi[upper] = touzi
	}
}

func (d *Dispatcher) Dispatch(request Request) (DispatchResult, error) {
	prefix, prefixLength := utf8.DecodeRuneInString(string(request))
	body := string(request)[prefixLength:]

}
