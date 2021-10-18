package main

import (
	"regexp"
	"strings"
	"testing"
)

const refString = "Mary had a little lamb"

func Benchmark_Strings_Split(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.Split(refString, " ")
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Strings_SplitN_Under(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.SplitN(refString, " ", 3)
	}
	if len(words) != 3 {
		b.Errorf("got len(words) = %d, want 3", len(words))
	}
}

func Benchmark_Strings_SplitN_Equal(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.SplitN(refString, " ", 5)
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Strings_SplitN_Over(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.SplitN(refString, " ", 10)
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Strings_SplitAfter(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.SplitAfter(refString, " ")
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Strings_Fields(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.Fields(refString)
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Strings_FieldsFunc(b *testing.B) {
	splitFunc := func(r rune) bool {
		return strings.ContainsRune(" ", r)
	}
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.FieldsFunc(refString, splitFunc)
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Strings_FieldsFunc_Inline(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = strings.FieldsFunc(refString, func(r rune) bool {
			return strings.ContainsRune(" ", r)
		})
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Regexp_Split_PreComp(b *testing.B) {
	var words []string
	re := regexp.MustCompile("[ ]{1}")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		words = re.Split(refString, -1)
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}

func Benchmark_Regexp_Split_Inline(b *testing.B) {
	var words []string
	for n := 0; n < b.N; n++ {
		words = regexp.MustCompile("[ ]{1}").Split(refString, -1)
	}
	if len(words) != 5 {
		b.Errorf("got len(words) = %d, want 5", len(words))
	}
}
