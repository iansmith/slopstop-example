// Baseline sanity tests for wordfreq.
//
// These exercise behavior that is orthogonal to the three known bugs (they use
// lowercase, unpunctuated input and never request more rows than exist), so
// they stay green on the current code and give slopstop a clean regression
// baseline to start from.
package main

import (
	"reflect"
	"testing"
)

func TestCountsSimpleWords(t *testing.T) {
	got := wordFrequencies("cat cat dog")
	want := map[string]int{"cat": 2, "dog": 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wordFrequencies = %v, want %v", got, want)
	}
}

func TestEmptyInputHasNoWords(t *testing.T) {
	if got := wordFrequencies(""); len(got) != 0 {
		t.Fatalf("wordFrequencies(\"\") = %v, want empty", got)
	}
}

func TestTopWordsOrdersByCountThenAlphabetically(t *testing.T) {
	counts := map[string]int{"cat": 2, "dog": 1, "ant": 1}
	got := topWords(counts, 10)
	want := []wordCount{{"cat", 2}, {"ant", 1}, {"dog", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("topWords = %v, want %v", got, want)
	}
}
