package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

const testString = "test"

func BenchmarkConcat(b *testing.B) {
	if testing.Short() {
		b.Skip("Skip long-running BenchmarkConcat")
	}
	var str string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += testString
	}
	b.StopTimer()
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString(testString)
	}
	b.StopTimer()
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], testString)
	}
	b.StopTimer()
}

func Benchmark_Strings_Builder_WriteString(b *testing.B) {
	var sb strings.Builder
	for n := 0; n < b.N; n++ {
		sb.WriteString(testString)
	}
}

func Benchmark_Strings_Builder_as_IoWriter(b *testing.B) {
	var sb strings.Builder
	var w io.Writer = &sb
	for n := 0; n < b.N; n++ {
		w.Write([]byte(testString))
	}
}
