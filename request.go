package touzi

import "strings"

type Request string

func SplitRequests(s string) (requests []Request) {
	for _, f := range strings.Fields(s) {
		for _, r := range strings.Split(f, ";") {
			if r != "" {
				requests = append(requests, Request(r))
			}
		}
	}

	return
}
