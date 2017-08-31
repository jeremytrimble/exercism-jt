package pangram

import "strings"

const testVersion = 2

type BitVec struct {
	bits uint32
}

func NewBitvec() BitVec {
	return BitVec{bits: 0}
}

func (bv *BitVec) SetRune(r rune) {
	bv.Set(uint(strings.ToUpper(string(r))[0]) - uint('A'))
}

func (bv *BitVec) Set(bit uint) {
	bv.bits |= (1 << bit)
}

func (bv *BitVec) IsComplete() bool {
	return bv.bits == 0x3FFFFFF
}

func IsPangram(in string) bool {
	var bv BitVec

	for _, r := range in {
		bv.SetRune(r)
	}

	return bv.IsComplete()
}
