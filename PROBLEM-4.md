# Part 4 — Week view

## What to build

The `cal week` subcommand. It shows a seven-day grid anchored to the Monday of
the week containing the given date (or this week if no date is given).

```
cal week [date-expression]
```

## Output format

```
$ cal week
Mon 2026-06-09  (nothing)
Tue 2026-06-10  16:45  Code review (#3)
Wed 2026-06-11  09:00  Team standup (#1)
               14:00  Afternoon check-in (#4)
Thu 2026-06-12  (nothing)
Fri 2026-06-13  (nothing)
Sat 2026-06-14  (nothing)
Sun 2026-06-15  (nothing)
```

Rules:
- Seven lines, always Monday through Sunday
- Each line: three-letter weekday abbreviation, space, ISO date, two spaces, then
  either `(nothing)` or the first event's `HH:MM  title (#id)`
- If a day has more than one event, additional events appear on continuation lines
  indented to align with the first event's time column (see Wed above)
- Events sorted by time ascending within each day
- The anchor date defaults to today; `cal week "next week"` should advance by
  7 days before finding the Monday anchor

## The complexity trap for this part

The tricky piece is the **multi-event continuation line**. The naive approach is
a nested loop with a flag variable and manual column arithmetic. That works but
quickly becomes hard to follow. Consider building a `_week_rows` helper that
returns a list of `(day_label, time_label, event_label)` tuples — then the
formatting loop becomes trivial.

This part is also a good place to notice whether your `store.py` API is
convenient. You will want something like `get_events_for_range(start, end)` that
returns events grouped by day, or at least sorted by `event_at` so the week view
can iterate through them cleanly. If that query is awkward to write, it is
probably worth adding a small helper to `store.py` rather than doing the
grouping in `display.py`.

## Acceptance criteria

- `cal week` shows the Monday-to-Sunday week containing today
- `cal week <expr>` shows the week containing the given date
- Multi-event days show correctly with aligned continuation lines
- `(nothing)` shown for empty days
- Week grid is always exactly 7 lines
- All formatting logic lives in `display.py`
- `format_week` (or equivalent) is tested directly with a hand-constructed
  event list, not via subprocess
- The subcommand handler in `cal.py` is thin (parse → query → format → print)

## Stretch goals (optional)

If you finish Parts 1–4 and want more:

- **Recurrence** — add a `--repeat daily|weekly|"first Monday"` flag to `cal add`.
  Recurring events are stored as rules and expanded on query. This is a second
  complexity bomb: the expansion logic for "nth weekday of month" forms will spike
  CC if written as a monolith.

- **Duration** — `cal add "Meeting" "tomorrow 9am" --duration 1h`. Store end time
  alongside start time and show duration in the day/week view.

- **Time zones** — `cal add "Call with London" "tomorrow 9am" --tz Europe/London`.
  Store in UTC, display in local time. Use `zoneinfo` (Python 3.9+).
