# How slopstop works — the building blocks

The [quickstart](https://github.com/iansmith/slopstop/blob/master/QUICKSTART.md)
gets you through the loop in 15 minutes by running five commands. This document
is for the reader who, after that, wants to know **what those commands actually
do** — the way you might learn git by understanding blobs, trees, and refs
instead of memorizing `git commit`.

Nothing here is magic. Every command is a scripted sequence of ordinary
operations — writing a file, making a git commit, adding a label, opening a PR —
and you can watch each one happen. The point of this doc is to name the pieces so
you can see the machine, not just drive it.

---

## The one idea

Slopstop makes two rules mechanical:

1. **Ticket-first.** Every change is anchored to a ticket. The ticket is the unit
   of work, and its number is the work's identity — the branch, the tracking
   notes, and the PR all carry it.
2. **Tests-first.** The failing test is written and **committed before** the fix.
   That commit is frozen: the test that proves the bug can't later be quietly
   edited to pass. This is the load-bearing primitive, and it gets its own section
   below.

Everything else is plumbing around those two rules.

---

## The primitives

### 1. The config — `.project-conf.toml`

The one file slopstop reads to know what project it is working on. Three required
fields answer three questions:

```toml
system = "github"                      # where do tickets live?
key    = "you/slopstop-example"        # which project?
prefix = "WORD"                        # what are its tickets called? (WORD-1, WORD-2, …)
```

Everything else in the file tunes behavior (which review backend, where tracking
notes go, how GitHub encodes "in progress"). It is committed, so the whole team
shares one answer. When a command says "read the config," this is the file.

### 2. The ticket — the unit of work

On GitHub the **issue number is the ticket's identity**. Issue #1 is `WORD-1`;
the branch is `fix/WORD-1`; the tracking notes live in `.slopstop/ticket-active/WORD-1/`.
They are all the same number on purpose — so at any moment you can go from a
branch name to its issue to its notes without a lookup table.

(GitHub draws issue *and* PR numbers from one counter, which is why after four
issues your first PR is #5 — not a mistake, just the shared counter.)

### 3. The tracking directory — a ticket's working memory

When you start a ticket, slopstop creates a directory named after it and writes
three files. This is the ticket's short-term memory — everything Claude knows
about the work in progress, on disk, in plain Markdown you can open:

```
.slopstop/ticket-active/WORD-1/
├── task_plan.md    the plan + the Definition of Done (the contract for "done")
├── findings.md     what Claude learned investigating (files, root cause, decisions)
└── progress.md     a session-by-session log, appended as the work proceeds
```

`task_plan.md` is the important one: it holds the **Definition of Done**, the
explicit list of what must be true before the ticket can close. The red tests are
written to satisfy it, and `:pr`'s gates check against it. Open these files while
a ticket is in flight — they are the clearest window into what the machine is
thinking.

At `:merge`, the whole directory moves to `.slopstop/ticket-archive/WORD-1/`. The
notes are never thrown away, but they never clutter the repo either — which is the
next primitive.

### 4. The three directories — what is kept, what is scratch

Slopstop uses three directories, and the only thing you need to remember is
**which ones git tracks**:

| Directory | Git | Lifespan | Holds |
|---|---|---|---|
| `design/` | **committed** | durable | design docs, decisions, invariants |
| `.slopstop/` | gitignored | per-ticket | the tracking dir above (active + archive) |
| `scratch/` | gitignored | per-run | transient artifacts from `:design` / `:run` (PRDs, charters, run state) |

The line runs between `design/` and the other two. `design/` is the record you
**keep** — read [`design/wordfreq.md`](design/wordfreq.md) for a worked example.
`.slopstop/` and `scratch/` are the machine's **short-term memory**: useful while
work is happening, disposable after. Because they are gitignored, they never end
up in a diff or a commit, and there are no stale per-ticket files rotting in the
repo months later.

A design doc and a tracking file can describe the same feature. The difference is
who they are for and how long they live: the design doc is for a future reader and
is committed; the tracking file is for the machine mid-task and is thrown away.

### 5. The frozen red test — why "tests-first" is more than a slogan

This is the primitive that makes the whole thing trustworthy, so it is worth
seeing exactly.

When `:plan` writes the failing test for WORD-1, it does three concrete things:

1. Runs the test and confirms it **fails** on the current (buggy) code.
2. Commits *just that test*, with a subject like
   `[WORD-1] Phase 0: red tests for case-insensitive counting`.
3. Treats that commit as **frozen**: from here on, the test may not be weakened,
   its expected value may not be changed, and it may not be deleted or skipped.

Why the ceremony? Because the easiest way to make a failing test pass is to edit
the test, and an AI eager to see green is exactly the thing that would do it — with
a confident commit message explaining why. Freezing the red commit turns "the
tests pass" into a claim you can actually trust: the test that proves the bug is
pinned in git history, and `:pr` mechanically diffs your later commits against it.
If an assertion changed, that is caught and refused. Making the test *pass* and
making the test *agree with your code* are different acts, and only the first one
is allowed.

You can watch this yourself: after `:plan`, run `git log --oneline` and you will
see the `Phase 0: red tests` commit sitting there before any fix.

### 6. The loop — five commands, and what each actually does

The commands are thin. Here is the machine underneath each:

| Command | What actually happens |
|---|---|
| `/slopstop:start WORD-1` | Adds the `status:in-progress` label to issue #1; creates branch `fix/WORD-1` off the default branch; creates `.slopstop/ticket-active/WORD-1/` with the three files. |
| `/slopstop:plan` | Writes the red test, **commits it frozen** (primitive #5), writes `task_plan.md` (plan + Definition of Done) and `findings.md` — then **stops**. It does not write the fix. |
| *(implement)* | You (or Claude, guided by the plan) write the fix until the red test goes green. Leave it **uncommitted** — `:pr`'s first step runs against the working tree. |
| `/slopstop:pr` | Simplifies the change; runs the tests; checks the frozen tests were not tampered with; commits; pushes; opens the PR; runs a Claude code review. |
| `/slopstop:merge` | Merges the PR; advances the ticket one state (3-state: closes #1, removes the label); archives the tracking dir to `.slopstop/ticket-archive/`. |

Two things surprise people, and both are deliberate:

- **`:plan` stops before implementing.** In this interactive flow it produces the
  plan and the frozen red test and hands back to you. The fix is a separate step.
  (When slopstop runs a fleet of autonomous agents via `:run`, the same `:plan`
  keeps going and implements — but that is the automated path, not this one.)
- **You leave the fix uncommitted until `:pr`.** `:pr` begins with a simplify pass
  that reads the working tree, so it needs your changes unstaged. `:pr` makes the
  commit, not you.

---

## The bigger machine (optional)

The five commands above are the everyday loop for one ticket. Slopstop has a
larger mode for planning and building many tickets at once — `:design` (turn an
idea into a PRD + charter), `:tickets` (cut the PRD into a tree of leaf tickets),
and `:run` (launch a fleet of isolated agents, one per ticket, and integrate their
work). That mode is where `scratch/runs/<run-id>/` comes from: a run's PRD,
charter, and state all live there, gitignored, and its durable conclusions are
distilled into `design/` or attached to the umbrella ticket.

You do not need any of it for this example. But it is the same primitives scaled
up: still ticket-first, still tests-first, still design-kept-and-scratch-thrown-away.

---

## Where to look next

- [`QUICKSTART.md`](https://github.com/iansmith/slopstop/blob/master/QUICKSTART.md)
  — the 15-minute hands-on path (start here if you haven't run the loop yet).
- [`design/wordfreq.md`](design/wordfreq.md) — a real design doc; see how each
  ticket's Definition of Done traces back to a stated invariant.
- [`TICKETS.md`](TICKETS.md) — the four tickets you'll drive through the loop.
- [`CONFIG.md`](https://github.com/iansmith/slopstop/blob/master/CONFIG.md) — every
  knob in `.project-conf.toml`.
