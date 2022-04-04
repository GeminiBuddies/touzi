package touzi

import (
	"strings"
	"touzi/random"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

type Roll struct {
	Request   Request
	Prefix    rune
	Arguments []Argument
	Format    string
	Touzi     Touzi
	Result    Result
}

type Cup struct {
	source random.Source
	touzi  map[rune]Touzi
}

func NewCup(source random.Source) *Cup {
	return &Cup{
		source: source,
		touzi:  make(map[rune]Touzi),
	}
}

func (d *Cup) Add(touzi Touzi) {
	touzi.InjectSource(d.source)

	prefix := touzi.Information().Prefix
	d.touzi[prefix] = touzi

	if lower := unicode.ToLower(prefix); lower != prefix {
		d.touzi[lower] = touzi
	}

	if upper := unicode.ToUpper(prefix); upper != prefix {
		d.touzi[upper] = touzi
	}
}

func (d *Cup) RollOne(request Request) (roll Roll, err error) {
	roll.Request = request

	prefix, prefixLength := utf8.DecodeRuneInString(string(request))
	body := string(request)[prefixLength:]
	roll.Prefix = prefix

	formatIndex := strings.IndexRune(body, '#')
	if formatIndex >= 0 {
		roll.Format = body[formatIndex+1:]
		body = body[:formatIndex]
	}

	if len(body) > 0 {
		var args = strings.Split(body, ",")
		roll.Arguments = *(*[]Argument)(unsafe.Pointer(&args))
	}

	touzi, exists := d.touzi[prefix]
	if !exists {
		err = ErrorTouziNotFound(prefix)
		return
	}

	roll.Touzi = touzi
	roll.Result, err = touzi.Roll(roll.Arguments, roll.Format)

	return
}
