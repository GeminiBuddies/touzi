package pcg

import (
	cr "crypto/rand"
	"encoding/binary"
	"lukechampine.com/uint128"
	mr "math/rand"
)

type PCG struct {
	State uint128.Uint128
	Inc   uint128.Uint128
}

var pcgMul128 = uint128.New(4865540595714422341, 2549297995355413924)

func Factory(seed, seq uint128.Uint128) *PCG {
	var pcg = &PCG{
		State: uint128.Zero,
		Inc:   seq.Lsh(1).Or64(1),
	}

	pcg.Next()
	pcg.State = pcg.State.AddWrap(seed)
	pcg.Next()

	return pcg
}

func New() *PCG {
	var b [32]byte

	if _, err := cr.Read(b[:]); err != nil {
		return Factory(uint128.New(mr.Uint64(), mr.Uint64()), uint128.New(mr.Uint64(), mr.Uint64()))
	}

	return Factory(
		uint128.New(binary.LittleEndian.Uint64(b[0:8]), binary.LittleEndian.Uint64(b[8:16])),
		uint128.New(binary.LittleEndian.Uint64(b[16:24]), binary.LittleEndian.Uint64(b[24:])),
	)
}

func (p *PCG) Next() uint64 {
	p.State = p.State.MulWrap(pcgMul128).AddWrap(p.Inc)

	xorshifted := p.State.Hi ^ p.State.Lo
	rot := p.State.Rsh(122).Lo
	return (xorshifted >> rot) | (xorshifted << ((-rot) & 63))
}
