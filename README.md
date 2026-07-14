# slopstop-example

A small playground for learning the [slopstop](https://github.com/iansmith/slopstop)
workflow — a ticket-anchored, tests-first way to drive Claude Code.

It's a tiny **word-frequency** command-line tool that counts the most common
words in a text file. It ships with **three bugs and one missing feature**
(see [TICKETS.md](TICKETS.md)) — just enough to take a real change from ticket
to merged PR and see the whole slopstop loop once with your own hands.

The same program exists in **Python** (`python/`) and **Go** (`go/`), with the
same behavior and the same bugs. Work in whichever you prefer.

> **New here? Follow the guided walkthrough:**
> **https://github.com/iansmith/slopstop/blob/master/QUICKSTART.md**
>
> **Want to understand how it works, not just run it?** Read
> [HOW-IT-WORKS.md](HOW-IT-WORKS.md) — the building blocks, one primitive at a time.

---

## Run it yourself

**Python** (needs Python 3.11+ and `pip install pytest`):
```
cd python
python3 wordfreq.py ../data/sample.txt --top 5
python3 -m pytest            # baseline sanity tests
```

**Go** (needs Go 1.21+):
```
cd go
go run . ../data/sample.txt --top 5
go test ./...                # baseline sanity tests
```

Run it and read the output — the bugs are visible right away (`The` and `the`
counted separately, `dog.` split from `dog`, and `--top 5` giving you only 4
rows). Fixing them is the quickstart.

---

## Layout

```
.project-conf.toml   slopstop config — edit `key` to point at your copy
HOW-IT-WORKS.md      the building blocks explained, one primitive at a time
TICKETS.md           the three bugs + one feature, as ticket descriptions
design/              committed, durable design docs (start with design/wordfreq.md)
data/sample.txt      the text the tool counts
data/stopwords.txt   common words, for the --stopwords feature (WORD-4)
python/              wordfreq.py + tests
go/                  wordfreq.go + tests
```

Two more directories appear while you work, and both are **gitignored** — they
hold slopstop's short-term memory, not source:

```
.slopstop/           per-ticket tracking notes (task_plan / findings / progress)
scratch/             transient run artifacts (only from the larger :design/:run mode)
```

The split between committed `design/` and gitignored `.slopstop/` + `scratch/` is
the heart of the layout — see [HOW-IT-WORKS.md](HOW-IT-WORKS.md), "The three
directories."
