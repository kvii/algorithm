package main

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

// calculate directly
func complexMultiplication1(a, b, c, d int) complex128 {
	return complex(float64(a), float64(b)) * complex(float64(c), float64(d))
}

// use multiply formula
func complexMultiplication2(a, b, c, d int) complex128 {
	return complex(float64(a*c-b*d), float64(a*d+b*c))
}

// use algorithm
func complexMultiplication3(a, b, c, d int) complex128 {
	var (
		s1 = a * (c + d)
		s2 = d * (b + a)
		s3 = c * (b - a)
	)
	// s1 - s2 = ac + ad - bd - ad = ac - bd (real)
	// s1 + s3 = ac + ad + bc - ac = ad + bc (imag)
	return complex(float64(s1-s2), float64(s1+s3))
}

func TestComplex(t *testing.T) {

	// test fn1 == fn2 == fn3
	var test = func(a, b, c, d int) bool {
		return complexMultiplication1(a, b, c, d) == complexMultiplication2(a, b, c, d) &&
			complexMultiplication2(a, b, c, d) == complexMultiplication3(a, b, c, d)
	}

	// to prevent args overflow.
	var filter = func(v []reflect.Value, r *rand.Rand) {
		for i := range v {
			v[i] = reflect.ValueOf(r.Intn(100_0000))
		}
	}

	conf := &quick.Config{Values: filter}

	assert.NoError(t, quick.Check(test, conf))
}

func Benchmark427_1(B *testing.B) {
	var fn = func() int { return rand.Intn(100_0000) }

	for i := 0; i < B.N; i++ {
		a, b, c, d := fn(), fn(), fn(), fn()
		for j := 0; j < 10000; j++ {
			_ = complexMultiplication1(a, b, c, d)
		}
	}
}
func Benchmark427_2(B *testing.B) {
	var fn = func() int { return rand.Intn(100_0000) }

	for i := 0; i < B.N; i++ {
		a, b, c, d := fn(), fn(), fn(), fn()
		for j := 0; j < 10000; j++ {
			_ = complexMultiplication2(a, b, c, d)
		}
	}
}
func Benchmark427_3(B *testing.B) {
	var fn = func() int { return rand.Intn(100_0000) }

	for i := 0; i < B.N; i++ {
		a, b, c, d := fn(), fn(), fn(), fn()
		for j := 0; j < 10000; j++ {
			_ = complexMultiplication3(a, b, c, d)
		}
	}
}
