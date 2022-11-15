package ex

import (
	src "study_go_test/src"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, src.Fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, src.Fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, src.Fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, src.Fibonacci1(3), "fibonacci1(3) should be 2")
	assert.Equal(233, src.Fibonacci1(13), "fibonacci1(13) should be 233")
}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, src.Fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, src.Fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, src.Fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, src.Fibonacci2(3), "fibonacci2(3) should be 2")
	assert.Equal(233, src.Fibonacci2(13), "fibonacci2(13) should be 233")
}

// 벤치마크 테스트
func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.Fibonacci1(20) // b.N 만큼 반복
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.Fibonacci2(20)
	}
}
