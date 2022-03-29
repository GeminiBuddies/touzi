package random

type Source interface {
	Next() uint64
}
