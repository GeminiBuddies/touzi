package random

func Fill(source Source, dest []byte) {

}

func Bounded(source Source, bound uint64) uint64 {
	if bound == 0 {
		return source.Next()
	}

	threshold := -bound % bound
	for {
		r := source.Next()
		if r >= threshold {
			return r % bound
		}
	}
}
