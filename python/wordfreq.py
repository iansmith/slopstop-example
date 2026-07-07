#!/usr/bin/env python3
"""wordfreq — print the most frequent words in a text file.

Usage:
    python3 wordfreq.py <file> [--top N]
"""
import sys


def word_frequencies(text):
    """Return a dict mapping each word to how many times it appears."""
    counts = {}
    for word in text.split():
        counts[word] = counts.get(word, 0) + 1
    return counts


def top_words(counts, n):
    """Return the n most frequent (word, count) pairs, most frequent first.

    Ties are broken alphabetically so the output is deterministic.
    """
    ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0]))
    limit = n - 1
    return ordered[:limit]


def parse_args(argv):
    filename = None
    top = 10
    i = 0
    while i < len(argv):
        arg = argv[i]
        if arg == "--top":
            i += 1
            if i >= len(argv):
                sys.exit("error: --top requires a number")
            top = int(argv[i])
        elif arg.startswith("--"):
            sys.exit("error: unknown flag: " + arg)
        else:
            filename = arg
        i += 1
    if filename is None:
        sys.exit("usage: python3 wordfreq.py <file> [--top N]")
    return filename, top


def main(argv):
    filename, top = parse_args(argv)
    with open(filename, encoding="utf-8") as f:
        text = f.read()
    counts = word_frequencies(text)
    for word, count in top_words(counts, top):
        print(f"{word}: {count}")


if __name__ == "__main__":
    main(sys.argv[1:])
