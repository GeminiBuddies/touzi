package random

func Fill(source Source, dest []byte) {
	var (
		val uint64
		rem = 0
		l   = len(dest)
	)

	for i := 0; i < l; i += 1 {
		if rem == 0 {
			val = source.Next()
			rem = 8
		}

		dest[i] = byte(val)
		val >>= 8
		rem -= 1
	}
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
