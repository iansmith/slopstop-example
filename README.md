# slopstop-example

A hands-on example project for learning the [slopstop](https://github.com/iansmith/slopstop)
workflow. You will direct Claude Code to build a small but real command-line calendar
program, broken into 4 parts. Each part is a GitHub issue in your fork. You will use
slopstop to direct Claude Code through the full cycle — planning, implementation, review,
and merge — the same workflow used on production software.

---

## Prerequisites

Before you start, make sure you have:

- **A Claude account** capable of running Claude Code (claude.ai/code or the
  desktop app). A Pro or Team plan is required.

- **Claude Code CLI** — the terminal interface to Claude Code:
  ```
  npm install -g @anthropic-ai/claude-code
  ```

- **gh** — the GitHub CLI, on your PATH and authenticated:
  ```
  brew install gh      # macOS; see https://cli.github.com for other platforms
  gh auth login
  ```

### Installing slopstop

Pick the method that matches how you run Claude Code.

**Claude Code CLI (terminal)** — two-step marketplace install:

```
claude plugin marketplace add iansmith/slopstop
claude plugin install slopstop@slopstop
```

Commands appear as `/slopstop:start`, `/slopstop:plan`, etc.

**Claude Desktop** — the desktop app does not yet support `/plugin install`, so
use the shell installer instead:

```bash
curl -fsSL https://raw.githubusercontent.com/iansmith/slopstop/master/install-for-claude-desktop.sh | bash
```

This downloads each skill file into `~/.claude/commands/`. Commands appear as
`/slopstop-start`, `/slopstop-plan`, etc. (hyphen instead of colon — same
skills, different namespace).

To pin to the current release (recommended):

```bash
SLOPSTOP_REF=v2.2.0 bash <(curl -fsSL https://raw.githubusercontent.com/iansmith/slopstop/master/install-for-claude-desktop.sh)
```

> The rest of this guide uses the CLI (`/slopstop:verb`) form. If you are on
> Desktop, substitute the hyphen form (`/slopstop-verb`) wherever you see a
> slash command.

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

Replace `<your-username>` with your GitHub username.

The `status_labels` section tells slopstop which labels to move issues through.
This example uses the **3-state workflow** (Todo → In Progress → Done), which only
requires an `in_progress` label. The 4-state workflow adds an `in_review` label
between In Progress and Done — do not add that here, or slopstop will expect an
extra review transition that this example does not use.

Create the label in your fork if it doesn't exist:

```bash
gh label create "in-progress" --color "0075ca" --description "Work in progress"
```

### 4. Run GitHub initialization

Start a Claude Code session in the repo directory and run:

```
/slopstop:gh-init --workflow 3
```

The `--workflow 3` flag tells slopstop to use the 3-state workflow (Todo →
In Progress → Done) and skips the interactive workflow question. If you omit
the flag, gh-init will ask "Which workflow? [3/4]" — answer `3`. Answering
`4` will add an `in_review` label and an extra transition step that this
example does not use.

gh-init will verify your configuration, confirm the repo, create the
`in-progress` label if it does not exist, and write `.project-conf.toml`.
It is safe to re-run if anything needs fixing.

### 5. Create the problem tickets

The problem is described in [PROBLEM.md](PROBLEM.md) and broken into four parts
in [PROBLEM-1.md](PROBLEM-1.md) through [PROBLEM-4.md](PROBLEM-4.md).

Run `/slopstop:create-gh` once for each part. The skill creates a GitHub issue
and assigns it the canonical `EXAMPLE-N` ticket key that the rest of the slopstop
workflow uses. In your Claude Code session:

```
/slopstop:create-gh --title "Part 1 — Date/time parser" --body "See PROBLEM-1.md"
/slopstop:create-gh --title "Part 2 — Event storage and CRUD" --body "See PROBLEM-2.md"
/slopstop:create-gh --title "Part 3 — Day view" --body "See PROBLEM-3.md"
/slopstop:create-gh --title "Part 4 — Week view" --body "See PROBLEM-4.md"
```

Or tell Claude: **"Create one GitHub issue for each part in PROBLEM-1.md through
PROBLEM-4.md using /slopstop:create-gh"** and it will run the skill four times,
reading each problem file for the title and body.

Once created, the issues are at `https://github.com/<your-username>/slopstop-example/issues`.

---

## Working through the problem

Once your issues are created, use the slopstop lifecycle for each one. Each
command is something you type; Claude Code does the work:

| Step | Command | What Claude Code does |
|------|---------|----------------------|
| Start a ticket | `/slopstop:start` | Creates a branch, marks the issue in-progress |
| Plan the work | `/slopstop:plan` | Writes red tests first, proposes an implementation plan |
| Implement | *(see below)* | Writes the code, runs tests, iterates until green |
| Open a PR | `/slopstop:pr` | Runs tests, complexity gate, CodeRabbit review |
| Merge | `/slopstop:merge` | Squash-merges, closes the issue, cleans up |
| Archive | `/slopstop:archive` | Moves the ticket's local tracking state from active to archive |

**On the implement step:** slopstop does not prescribe a single command here
because different projects have different build and test setups. After
`/slopstop:plan` completes, direct Claude Code to execute it:

- Say **"implement the plan"** — Claude Code will follow the plan it just
  produced, write all the code, run the tests, and iterate until they pass.
- If your project has a project-specific skill for this, use that instead.

Work the parts in order (Part 1 → Part 4): each part builds on the previous one.

### What to watch for

- **Part 1 (date parser)** is where the slopstop complexity gate will fire.
  A naive `if/elif` chain over all the supported date formats produces
  cyclomatic complexity well above the threshold — Claude Code will be asked
  to refactor before the PR can proceed. That is intentional: watch Claude
  decompose the parser into smaller functions and see the score drop.

- **Parts 3 and 4 (views)** are where the design decisions from Parts 1 and 2
  pay off. If Claude Code produced a clean parser boundary in Part 1, the
  display code writes itself. If the parser leaked complexity forward, you will
  see it here.

---

## Reference

- [slopstop repo](https://github.com/iansmith/slopstop) — full documentation
- [slopstop cold-start guide](https://github.com/iansmith/slopstop/blob/master/START-HERE.md) — detailed first-time setup
- [PROBLEM.md](PROBLEM.md) — full problem description
- PROBLEM-1.md through PROBLEM-4.md — per-part specs
