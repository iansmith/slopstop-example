# slopstop-example

A hands-on example project for learning the [slopstop](https://github.com/iansmith/slopstop)
workflow. You will build a small but real command-line calendar program, broken
into 4 parts. Each part is a GitHub issue in your fork. You will use slopstop to
plan, implement, review, and merge each part — the same workflow used on production
software.

---

## Prerequisites

Before you start, make sure you have:

- **Claude Code** — the Anthropic CLI. You need a Claude account that can run
  Claude Code (claude.ai/code). Install it with:
  ```
  npm install -g @anthropic-ai/claude-code
  ```
- **gh** — the GitHub CLI, on your PATH and authenticated.
  ```
  brew install gh      # macOS
  gh auth login
  ```
- **slopstop plugin** — install it into Claude Code:
  ```
  claude plugin marketplace add iansmith/slopstop
  claude plugin install slopstop@slopstop
  ```

---

## Setup (do this once)

### 1. Fork the repo

Go to https://github.com/iansmith/slopstop-example and click **Fork**.

### 2. Clone your fork

```bash
git clone git@github.com:<your-username>/slopstop-example.git
cd slopstop-example
```

All remaining steps happen inside this directory.

### 3. Create `.project-conf.toml`

Slopstop needs to know which ticket system and repo to use. Create this file at
the root of your clone:

```toml
system  = "github"
key     = "EXAMPLE"
prefix  = "EXAMPLE"

[github]
owner = "<your-username>"
repo  = "slopstop-example"

[status_labels]
in_progress = "in-progress"
```

Replace `<your-username>` with your GitHub username. The `status_labels` section
tells slopstop which labels to move issues through — create those labels in your
fork if they don't exist:

```bash
gh label create "in-progress" --color "0075ca" --description "Work in progress"
```

### 4. Run GitHub initialization

Start a Claude Code session in the repo directory and run:

```
/slopstop:gh-init
```

This will verify your configuration, confirm the label setup, and write any
missing scaffolding. Answer the prompts — it is safe to re-run if anything
needs fixing.

### 5. Create the problem tickets

The problem is described in [PROBLEM.md](PROBLEM.md) and broken into four parts
in [PROBLEM-1.md](PROBLEM-1.md) through [PROBLEM-4.md](PROBLEM-4.md).

Claude can create all four GitHub issues for you automatically. In your Claude
Code session run:

```
/slopstop:start
```

Then tell Claude: **"Create GitHub issues for each part of PROBLEM.md"** and it
will read the problem files and open the issues in your fork. Once the issues
exist you can see them at `https://github.com/<your-username>/slopstop-example/issues`.

---

## Working through the problem

Once your issues are created, use the slopstop lifecycle for each one:

| Step | Command | What it does |
|------|---------|--------------|
| Start a ticket | `/slopstop:start` | Creates a branch, marks the issue in-progress |
| Plan the work | `/slopstop:plan` | Writes red tests first, proposes an implementation plan |
| Open a PR | `/slopstop:pr` | Runs tests, complexity gate, CodeRabbit review |
| Merge | `/slopstop:merge` | Squash-merges, closes the issue, cleans up |

Work the parts in order (Part 1 → Part 4): each part builds on the previous one.

### What to watch for

- **Part 1 (date parser)** is where cyclomatic complexity will push back on you.
  A naive `if/elif` chain over all the supported date formats will trigger the
  slopstop CC gate. That is intentional — it is showing you exactly what the gate
  is for. Decompose the parser into smaller functions and watch the score drop.

- **Part 3 and 4 (views)** are where the design decisions from Part 1 and 2
  pay off. If your parser returns a clean `datetime`, the display code is simple.
  If it returns a tangle of strings, you will feel it here.

---

## Reference

- [slopstop repo](https://github.com/iansmith/slopstop) — full documentation
- [slopstop cold-start guide](https://github.com/iansmith/slopstop/blob/master/START-HERE.md) — detailed first-time setup
- [PROBLEM.md](PROBLEM.md) — full problem description
- PROBLEM-1.md through PROBLEM-4.md — per-part specs
