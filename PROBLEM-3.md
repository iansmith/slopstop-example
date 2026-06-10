# Part 3 — Day view

## What to build

The `cal day` subcommand and a `display.py` module that handles all formatted
output for the program.

```
cal day [date-expression]
```

With no argument, show today. With an argument, show that day.

## Output format

```
$ cal day tomorrow
Wednesday 2026-06-11
  09:00  Team standup (#1)
  14:00  Afternoon check-in (#3)

$ cal day "next Monday"
Monday 2026-06-15
  (no events)
```

Rules:
- Header line: weekday name, then the ISO date
- Each event line: two-space indent, `HH:MM`, two spaces, title, space, `(#id)`
- Events sorted by time ascending
- If no events: two-space indent, `(no events)`

## The `display.py` module

Move all output formatting into `display.py`. Functions to export:

```python
def format_day(date: date, events: list[Event]) -> str:
    """Return the formatted day view as a string (caller prints it)."""
    ...
```

Returning a string instead of printing directly makes this trivially testable —
you can assert on the exact output without capturing stdout.

## What to watch for

**Test the formatter, not just the command.** `format_day` takes a plain list
of events and returns a string. Test it with hand-constructed event lists. You
do not need a database to test the display logic.

**Keep the subcommand handler thin.** `cal day` should:
1. Parse the date expression
2. Query the store for events on that day
3. Call `display.format_day`
4. Print the result

If you find logic accumulating in the subcommand handler itself (beyond those
four steps), move it into `display.py` or `store.py`.

The CC gate is less likely to fire on this part than on Part 1, but a formatter
that uses a deep chain of conditionals to handle edge cases (no events, single
event, multi-event, past dates, future dates) can still get unwieldy. Keep each
formatter function focused on one output concern.

## Acceptance criteria

- `cal day` (no argument) shows today correctly
- `cal day <expr>` shows that day's events sorted by time
- `(no events)` shown when the day is empty
- Header uses the full weekday name
- All formatting logic lives in `display.py`
- `format_day` is tested directly (not via subprocess)
- Subcommand handler in `cal.py` is ≤ 15 lines excluding the argument registration
