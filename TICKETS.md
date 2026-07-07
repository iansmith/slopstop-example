# The four tickets

This repo ships with three bugs and one missing feature. In the quickstart you
create these as four GitHub issues (the quickstart gives you a block to paste
into Claude Code, which files them for you via `/slopstop:create-gh`). Because
`/slopstop:create-gh` makes the ticket key equal the issue number, the first
four issues become `WORD-1` … `WORD-4`.

The bugs live identically in both `python/wordfreq.py` and `go/wordfreq.go` —
fix whichever language you launched Claude Code in.

---

## WORD-1 (bug) — Word counting is case-sensitive

`wordfreq` counts differently-cased forms of the same word as different words.
Run it on the sample text and you'll see both `The` and `the` as separate rows.

**Expected:** counting is case-insensitive — `The`, `the`, and `THE` all count
as the same word, `the`.

**Repro:**
```
python3 wordfreq.py ../data/sample.txt --top 5      # (or: go run . ../data/sample.txt --top 5)
```
Observe `The` and `the` as separate entries.

---

## WORD-2 (bug) — Punctuation is not stripped from words

Words with attached punctuation are counted separately from the bare word, so
`dog`, `dog.`, and `dog,` are three different entries.

**Expected:** surrounding punctuation (`. , ; : ! ?` and quotes) is stripped
from each word before counting, so `dog.` and `dog,` both count as `dog`.

**Repro:**
```
python3 wordfreq.py ../data/sample.txt --top 10
```
Observe `dog`, `dog.`, `dog,` (and similar) as separate entries.

---

## WORD-3 (bug) — `--top N` returns one fewer result than requested

`--top N` prints `N-1` rows instead of `N`.

**Expected:** `--top 5` prints exactly 5 rows (or fewer only when there aren't 5
distinct words).

**Repro:**
```
python3 wordfreq.py ../data/sample.txt --top 5
```
Count the rows — there are only 4.

---

## WORD-4 (feature) — Add `--stopwords` to filter out common words

Add a `--stopwords` flag. When passed, words that appear in a common-word
stoplist are excluded from the counts, so the output surfaces meaningful content
words instead of `the`, `and`, `a`, and so on.

Use the provided list in `data/stopwords.txt` (one word per line). Matching is
case-insensitive.

**Expected:**
```
python3 wordfreq.py ../data/sample.txt --top 5 --stopwords
```
prints the top 5 words with stopwords removed (no `the`, `and`, `a`, `is`,
`over`).
