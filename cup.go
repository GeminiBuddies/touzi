package touzi

import (
	"touzi/random"
	"unicode"
)

type Roll struct {
	Request         Request
	Touzi           Touzi
	Result          Result
	FormattedResult string
}

type Cup struct {
	source    random.Source
	touzi     map[rune]Touzi
	formatter Formatter
}

func NewCup(source random.Source, formatter Formatter, touzi ...Touzi) *Cup {
	c := &Cup{
		source:    source,
		touzi:     make(map[rune]Touzi),
		formatter: formatter,
	}

	for _, t := range touzi {
		t.InjectSource(c.source)

		prefix := t.Information().Prefix
		c.touzi[prefix] = t

		if lower := unicode.ToLower(prefix); lower != prefix {
			c.touzi[lower] = t
		}

		if upper := unicode.ToUpper(prefix); upper != prefix {
			c.touzi[upper] = t
		}
	}

	return c
}

func (c *Cup) RollOne(request Request) (roll Roll, err error) {
	roll.Request = request

	touzi, exists := c.touzi[request.Prefix]
	if !exists {
		err = ErrorTouziNotFound(request.Prefix)
		return
	}

	roll.Touzi = touzi
	roll.Result, err = touzi.Roll(request.Arguments)

	if err != nil {
		return
	}

	roll.FormattedResult = c.formatter.Format(roll.Result, roll.Request.Format)
	return
}
