package touzi

import (
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
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

func (d *Dispatcher) Dispatch(request Request) (result DispatchResult, err error) {
	result.Request = request

	prefix, prefixLength := utf8.DecodeRuneInString(string(request))
	body := string(request)[prefixLength:]
	result.Prefix = prefix

	formatIndex := strings.IndexRune(body, '#')
	if formatIndex >= 0 {
		result.Format = body[formatIndex+1:]
		body = body[:formatIndex]
	}

	if len(body) > 0 {
		var args = strings.Split(body, ",")
		result.Arguments = *(*[]Argument)(unsafe.Pointer(&args))
	}

	touzi, exists := d.touzi[prefix]
	if !exists {
		err = ErrorTouziNotFound(prefix)
		return
	}

	result.Touzi = touzi
	result.Result, err = touzi.Roll(result.Arguments, result.Format)

	return
}
