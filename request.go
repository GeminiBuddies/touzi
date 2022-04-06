package touzi

import (
	"strings"
	"unicode/utf8"
	"unsafe"
)

type Request struct {
	Raw       string
	Prefix    rune
	Arguments []Argument
	Format    string
}

func ParseRequest(s string) (request Request) {
	request.Raw = s

	prefix, prefixLength := utf8.DecodeRuneInString(s)
	body := s[prefixLength:]
	request.Prefix = prefix

	formatIndex := strings.IndexRune(body, '#')
	if formatIndex >= 0 {
		request.Format = body[formatIndex+1:]
		body = body[:formatIndex]
	}

	if len(body) > 0 {
		var args = strings.Split(body, ",")
		request.Arguments = *(*[]Argument)(unsafe.Pointer(&args))
	}

	return
}

func SplitAndParseRequests(s string) (requests []Request) {
	for _, f := range strings.Fields(s) {
		for _, r := range strings.Split(f, ";") {
			if r != "" {
				requests = append(requests, ParseRequest(r))
			}
		}
	}

	return
}
