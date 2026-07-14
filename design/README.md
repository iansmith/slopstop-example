# `design/` — the durable, committed design record

This directory is **committed to git**. It holds the human-curated design docs
that outlive any single ticket: the shape of the program, decisions and their
rationale, invariants the code must keep. If you would want a teammate (or
future-you) to read it six months from now, it goes here.

Contrast it with the two directories slopstop keeps **out** of git:

| Directory | Git | Lifespan | Holds |
|---|---|---|---|
| `design/` | **committed** | durable | design docs, decisions, invariants (this dir) |
| `.slopstop/` | gitignored | per-ticket | working notes for a ticket in flight (`task_plan.md`, `findings.md`, `progress.md`) |
| `scratch/` | gitignored | per-run | transient artifacts from `:design` / `:run` (PRDs, charters, run state) |

The rule of thumb: **`design/` is the record you keep; `.slopstop/` and
`scratch/` are the machine's short-term memory.** Nothing per-ticket or per-run
is ever committed — no stale landmines for future sessions. When a bigger piece
of work produces a PRD or a charter in `scratch/`, its durable conclusions get
distilled into a design doc here (or attached to the umbrella ticket); the
scratch copy is disposable.

For the wider picture — how a ticket flows through the tracking directory and
the plan → red-test → code → PR → merge loop — see
[`../HOW-IT-WORKS.md`](../HOW-IT-WORKS.md).

## What's here

- [`wordfreq.md`](wordfreq.md) — the design of the word-frequency tool: what it
  does, the counting pipeline, and the invariants the four tickets each touch.
  Read it before WORD-1 and you'll see how a ticket's "Definition of Done" traces
  back to a stated invariant here.
