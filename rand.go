package rtrand

import (
	"math/rand"
	"time"
)

func ExpFloat64() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.ExpFloat64()
}

func Float32() float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float32()
}

func Float64() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()
}

func Int() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

func Int31() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31()
}

func Int31n(n int32) int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31n(n)
}

func Int63() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}

func Int63n(n int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(n)
}

func Intn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func NormFloat64() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.NormFloat64()
}

func Perm(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Perm(n)
}

func Uint32() uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint32()
}
