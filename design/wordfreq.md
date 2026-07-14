# Design — the `wordfreq` tool

`wordfreq` reads a text file and prints the most frequent words. One program,
two identical implementations (`python/wordfreq.py`, `go/wordfreq.go`).

This is a **design doc**, so it states how the tool is *supposed* to behave. The
four tickets in [`../TICKETS.md`](../TICKETS.md) exist because the shipped code
violates the invariants below — each ticket restores one of them. That is the
whole shape of the workflow: the invariant is stated here, the ticket names the
gap, the red test pins the invariant as an executable check, and the fix makes
the code satisfy it.

## The pipeline

Counting runs as a small pipeline. Read it top to bottom:

```
read file
  → normalize   lowercase each token, strip surrounding punctuation
  → tokenize    split on whitespace into words
  → count       tally occurrences into a map
  → filter      (optional) drop stopwords
  → rank        sort by count desc, ties broken alphabetically
  → take N      emit the top N rows
```

The current code collapses `normalize` into nothing (it splits raw whitespace
tokens and counts them verbatim) and has an off-by-one in `take N`. The tickets
walk the pipeline back into shape.

## Invariants

These are the promises the tool makes. Each maps to a ticket.

1. **Case-insensitive counting.** `The`, `the`, and `THE` are the same word,
   counted as `the`. *(WORD-1 — the `normalize` step does not lowercase.)*

2. **Punctuation-insensitive counting.** Surrounding punctuation is not part of a
   word: `dog`, `dog.`, and `dog,` all count as `dog`. Interior marks (e.g. an
   apostrophe in `don't`) are left alone. *(WORD-2 — `normalize` does not strip.)*

3. **`--top N` emits exactly N rows** (or fewer only when the text has fewer than
   N distinct words). *(WORD-3 — `take N` is off by one.)*

4. **Deterministic order.** Rank is by count descending; ties break
   alphabetically. This invariant already holds — keep it. A fix that makes the
   output non-deterministic (e.g. relying on map iteration order) breaks it.

5. **Optional stopword filtering.** With `--stopwords`, words in the common-word
   list (`data/stopwords.txt`, matched case-insensitively) are dropped before
   ranking. *(WORD-4 — the `filter` step does not exist yet.)*

## Why write this down

You could fix all four bugs by reading the code alone. The design doc earns its
place the moment the answer is *not* obvious — when "the right behavior" is a
decision, not a deduction. Invariant 2 is the example: should `don't` become
`dont`? The code can't tell you; the design doc decides (no — interior marks
stay), and the red test for WORD-2 encodes that decision so no later change can
quietly reverse it.

That is the durable half of the workflow. The transient half — the plan, the
findings, the session log for one ticket — lives in `.slopstop/` and is thrown
away when the ticket lands. See [`../HOW-IT-WORKS.md`](../HOW-IT-WORKS.md).
