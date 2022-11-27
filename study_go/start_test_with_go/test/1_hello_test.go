package ex

import (
	src "study_go_test/src"
	"testing"
)

func TestHello(t *testing.T) {
	got := src.Hello()
	want := "Hello world!"

	AssertCorrectMessage(t, got, want)
}
