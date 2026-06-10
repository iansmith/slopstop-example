# The Problem: a command-line calendar

You are going to build `cal` — a simple, text-based calendar program that runs
entirely from the command line. No GUI, no server, no web framework. Just a
program you can run to add events, delete them, and see what is on your schedule.

The interesting part is **how dates are entered**. Real people do not type ISO
timestamps. They type things like "tomorrow", "next Tuesday at 3pm", or "in two
hours". Your program has to accept all of that and store it correctly.

## What you will build

A single executable called `cal` (or `cal.py`, whatever suits your language)
with these subcommands:

```
cal add <title> <date-expression>       # add an event
cal del <id>                            # delete an event by id
cal list [date-expression]              # list events (default: today)
cal day  [date-expression]              # show one day's events with times
cal week [date-expression]              # show a week grid
```

Example session:

```
$ cal add "Team standup" "tomorrow 9am"
Added event #1: Team standup — 2026-06-11 09:00

$ cal add "Dentist" "next Friday 2:30pm"
Added event #2: Dentist — 2026-06-19 14:30

$ cal add "Code review" "in 3 hours"
Added event #3: Code review — 2026-06-10 16:45

$ cal day tomorrow
Wednesday 2026-06-11
  09:00  Team standup (#1)

$ cal week
Mon 2026-06-09  (nothing)
Tue 2026-06-10  16:45  Code review (#3)
Wed 2026-06-11  09:00  Team standup (#1)
Thu 2026-06-12  (nothing)
Fri 2026-06-13  (nothing)
Sat 2026-06-14  (nothing)
Sun 2026-06-15  (nothing)

$ cal del 1
Deleted event #1: Team standup
```

## The four parts

Work through these in order. Each one is a separate GitHub issue in your repo.

| Part | Topic | File |
|------|-------|------|
| 1 | Date/time parser | [PROBLEM-1.md](PROBLEM-1.md) |
| 2 | Event storage and CRUD | [PROBLEM-2.md](PROBLEM-2.md) |
| 3 | Day and list views | [PROBLEM-3.md](PROBLEM-3.md) |
| 4 | Week view | [PROBLEM-4.md](PROBLEM-4.md) |

## Language and dependencies

Use any language you like. Python is suggested because the standard library has
everything you need (`datetime`, `re`, `sqlite3`, `argparse`) and there is no
setup friction.

**Do not use a third-party date-parsing library** for Part 1. The entire point of
that part is to write the parser yourself and watch how its structure affects the
complexity score. You can add a library later if you want — but write it from
scratch first.

## What you are learning

This problem is chosen because **Part 1 has a complexity trap**. The obvious
implementation — one big function with an `if/elif` chain for each date format —
will produce cyclomatic complexity well above the slopstop gate threshold. You
will see the gate fire, understand why, refactor into smaller pieces, and watch
the score drop. That is the lesson.

Parts 2–4 then show how the design choices you made in Part 1 propagate forward.
A clean parser makes the rest easy. A tangled one makes the rest painful.
