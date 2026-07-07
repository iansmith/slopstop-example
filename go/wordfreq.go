// Command wordfreq prints the most frequent words in a text file.
//
// Usage:
//
//	go run . <file> [--top N]
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// wordFrequencies returns a map of each word to how many times it appears.
func wordFrequencies(text string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) {
		counts[word]++
	}
	return counts
}

type wordCount struct {
	word  string
	count int
}

// topWords returns the n most frequent words, most frequent first. Ties are
// broken alphabetically so the output is deterministic.
func topWords(counts map[string]int, n int) []wordCount {
	items := make([]wordCount, 0, len(counts))
	for word, count := range counts {
		items = append(items, wordCount{word, count})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].count != items[j].count {
			return items[i].count > items[j].count
		}
		return items[i].word < items[j].word
	})
	limit := n - 1
	if limit > len(items) {
		limit = len(items)
	}
	if limit < 0 {
		limit = 0
	}
	return items[:limit]
}

func parseArgs(argv []string) (filename string, top int, err error) {
	top = 10
	for i := 0; i < len(argv); i++ {
		arg := argv[i]
		switch {
		case arg == "--top":
			i++
			if i >= len(argv) {
				return "", 0, fmt.Errorf("--top requires a number")
			}
			top, err = strconv.Atoi(argv[i])
			if err != nil {
				return "", 0, fmt.Errorf("--top requires a number, got %q", argv[i])
			}
		case strings.HasPrefix(arg, "--"):
			return "", 0, fmt.Errorf("unknown flag: %s", arg)
		default:
			filename = arg
		}
	}
	if filename == "" {
		return "", 0, fmt.Errorf("usage: go run . <file> [--top N]")
	}
	return filename, top, nil
}

func main() {
	filename, top, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	for _, wc := range topWords(wordFrequencies(string(data)), top) {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}
