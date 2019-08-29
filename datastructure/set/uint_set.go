package main

import "math/bits"

const (
	bucket_size = 32
)

type UintSet []uint32

func NewUintSet(size int) UintSet {
	return make([]uint32, (size+bucket_size-1)/bucket_size)
}

func (s UintSet) offset(number uint32) (bucket, bit uint32) {
	bucket = number / bucket_size
	bit = 1 << (number % bucket_size)
	return bucket, bit
}

func (s UintSet) Add(number uint32) {
	bucket, bit := s.offset(number)
	s[bucket] |= bit
}

func (s UintSet) Clear(number uint32) {
	bucket, bit := s.offset(number)
	s[bucket] &= bits.Reverse32(bit)
}

func (s UintSet) Contains(number uint32) bool {
	bucket, bit := s.offset(number)
	return s[bucket]&bit != 0
}
