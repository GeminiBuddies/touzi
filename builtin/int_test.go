package builtin

import (
	"fmt"
	"sort"
	"testing"
	"touzi"
	"touzi/pcg"
)

func distribution(a *Int, count uint64, args []touzi.Argument, format string) {
	dist := make(map[string]uint64)

	for i := uint64(0); i < count; i += 1 {
		if result, err := a.Roll(args, format); err != nil {
			panic(err)
		} else {
			dist[string(result)] += 1
		}
	}

	var values []string
	for value := range dist {
		values = append(values, value)
	}

	sort.Strings(values)

	for _, value := range values {
		fmt.Printf("%s: %d(%.4f)\n", value, dist[value], float64(dist[value])/float64(count))
	}
}

func TestInt_Roll(t *testing.T) {
	a := &Int{
		source: pcg.New(),
	}

	distribution(a, 1048576*16, []touzi.Argument{"u1"}, "")
}
