# Part 2 — Event storage and CRUD

## What to build

A storage module (e.g. `store.py`) and the `add`, `del`, and `list` subcommands
of the `cal` program.

The storage backend is SQLite. Use a single file, defaulting to `~/.cal.db`.
Allow the path to be overridden via a `CAL_DB` environment variable (this makes
testing trivial — point it at a temp file).

## Data model

One table:

```sql
CREATE TABLE IF NOT EXISTS events (
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    title     TEXT    NOT NULL,
    event_at  TEXT    NOT NULL   -- ISO 8601: "2026-06-11T09:00:00"
);
```

No recurrence, no duration, no timezone for now. An event is a title and a moment.

## The `cal` entry point

Wire up argument parsing (use `argparse` or equivalent) so these work:

```
cal add <title> <date-expression>
cal del <id>
cal list [date-expression]
```

For `add`, pass `<date-expression>` through the parser from Part 1. Store the
resulting `datetime` as an ISO string.

For `list` with no argument, default to today. With an argument, show all events
whose `event_at` falls on that calendar day (midnight-to-midnight).

## Output format

```
$ cal add "Team standup" "tomorrow 9am"
Added event #4: Team standup — 2026-06-11 09:00

$ cal list tomorrow
2026-06-11
  #4  09:00  Team standup

$ cal del 4
Deleted event #4: Team standup

$ cal list tomorrow
2026-06-11
  (no events)
```

## What to watch for

This part is simpler than Part 1 — the complexity lives in the parser, not here.
The thing to notice is how cleanly `store.py` can be if `parse()` returns a proper
`datetime`. If you find yourself doing string manipulation in `store.py` to work
around ambiguity from Part 1, that is a signal the parser boundary is leaking.

The slopstop plan step will ask you to write tests before implementing. For this
module, use the `CAL_DB` env var to point at a temp file in your test fixture so
you never touch the real `~/.cal.db` during testing.

## Acceptance criteria

- `cal add` stores an event and prints the assigned id and resolved datetime
- `cal del <id>` removes the event and prints confirmation; prints a clear error if
  the id does not exist
- `cal list` (no argument) shows today's events
- `cal list <date-expression>` shows events on that day
- All storage operations go through `store.py`; `cal.py` only handles argument
  parsing and calls into `store.py`
- Tests use a temp SQLite file, never `~/.cal.db`
- All functions in `store.py` have cyclomatic complexity below 10
