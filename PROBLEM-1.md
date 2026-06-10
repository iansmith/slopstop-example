# Part 1 — Date/time parser

## What to build

A single module (e.g. `dateparse.py`) that exports one public function:

```python
def parse(expr: str, now: datetime) -> datetime:
    ...
```

`expr` is a human-written date expression. `now` is the reference point for
relative expressions (always pass it in — never call `datetime.now()` inside
the parser, so the function is purely testable). Returns a `datetime` with
date and time both set, or raises `ValueError` with a helpful message if the
expression cannot be parsed.

## Expressions to support

Your parser must handle all of the following:

**Absolute dates (no time component — default to 00:00)**
- `2026-01-25` — ISO date
- `Jan 25`, `January 25`, `Jan 25 2027` — month-name forms
- `the 25th`, `25th` — day-of-month relative to the current month

**Relative days**
- `today`, `tomorrow`, `yesterday`
- `next Monday` … `next Sunday` — the named weekday in the coming week (never today)
- `last Monday` … `last Sunday` — the named weekday in the previous week
- `in N days`, `in N weeks` — N is a positive integer

**Time qualifiers (combine with any date expression above)**
- `9am`, `9:00`, `9:00am`, `14:30`, `2:30pm`
- `morning` (09:00), `noon` (12:00), `afternoon` (14:00), `evening` (18:00), `midnight` (00:00)
- `in N hours`, `in N minutes` — relative to `now`, ignoring any date part

**Combined expressions**
- `tomorrow 9am`, `next Friday 2:30pm`, `Jan 25 at 3pm`
- `in 3 hours` — relative to now (date and time both come from the offset)

Parsing is case-insensitive. Ignore leading/trailing whitespace.

## The complexity trap

If you write this as a single function that tries each format in sequence with
`if/elif`, you will end up with something like:

```python
def parse(expr, now):
    expr = expr.strip().lower()
    if expr == "today":
        ...
    elif expr == "tomorrow":
        ...
    elif expr == "yesterday":
        ...
    elif expr.startswith("next "):
        ...
    elif expr.startswith("last "):
        ...
    elif expr.startswith("in "):
        ...
    elif re.match(r'\d{4}-\d{2}-\d{2}', expr):
        ...
    elif re.match(r'(jan|feb|...)', expr):
        ...
    # ... 10 more branches ...
    else:
        raise ValueError(...)
```

That function will have cyclomatic complexity around 15–20. The slopstop CC gate
will flag it. That is the intended experience — read the gate output, understand
the score, and refactor.

A decomposed approach has much lower total CC even though it has more functions:

- `_parse_date_part(tokens)` — handles the date portion only
- `_parse_time_part(tokens)` — handles the time portion only  
- `_resolve_weekday(name, direction, now)` — next/last weekday logic
- `_resolve_relative(n, unit, now)` — in-N-days/hours/minutes logic
- `parse(expr, now)` — tokenizes and dispatches to the above

Each small function has CC 3–5. The total is higher but no single function
exceeds the threshold.

## Acceptance criteria

- All of the expression forms listed above parse correctly
- `parse` is a pure function: same inputs → same output; no calls to
  `datetime.now()` or any I/O inside the module
- Every expression form has at least one test in `tests/test_dateparse.py`
  (write the tests first — this is a perfect case for TDD since the input/output
  pairs are completely specified)
- No single function in the module has cyclomatic complexity above 10
  (the slopstop `:pr` gate will enforce this)
- `ValueError` is raised with a readable message for inputs that cannot be parsed

## Suggested test cases

```python
# Absolute
assert parse("2026-01-25", now) == datetime(2026, 1, 25, 0, 0)
assert parse("Jan 25", now) == datetime(now.year, 1, 25, 0, 0)
assert parse("the 3rd", now) == datetime(now.year, now.month, 3, 0, 0)

# Relative days
assert parse("today", now) == now.replace(hour=0, minute=0, second=0, microsecond=0)
assert parse("tomorrow", now) == (now + timedelta(days=1)).replace(hour=0, minute=0, second=0, microsecond=0)
assert parse("in 3 days", now).date() == (now + timedelta(days=3)).date()

# Times
assert parse("tomorrow 9am", now).hour == 9
assert parse("next Friday 2:30pm", now).hour == 14
assert parse("next Friday 2:30pm", now).minute == 30

# Relative time
base = datetime(2026, 6, 10, 13, 45, 0)
assert parse("in 3 hours", base) == datetime(2026, 6, 10, 16, 45, 0)
assert parse("in 90 minutes", base) == datetime(2026, 6, 10, 15, 15, 0)

# Errors
with pytest.raises(ValueError):
    parse("banana", now)
```
