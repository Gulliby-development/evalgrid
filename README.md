<p align="center"><img src="docs/logo.svg" width="96" alt="" /></p>

# evalgrid

[![CI](https://github.com/Gulliby-development/evalgrid/actions/workflows/ci.yml/badge.svg)](https://github.com/Gulliby-development/evalgrid/actions/workflows/ci.yml) [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

evalgrid is a standard-library-only Go command-line tool for evaluation-harness
teams. It reads reference answers and model samples from a plain-text ledger
on standard input, then reports exact-match, token F1 and ROUGE-L scores,
unbiased pass@k estimates, per-question comparisons, or a CI score gate.

## Build and run

Build from the repository root with Go:

```sh
go build -o evalgrid .
```

On Windows, use `go build -o evalgrid.exe .` and invoke `evalgrid.exe`.
The examples below use POSIX shell syntax. In cmd.exe, use
`evalgrid.exe check --min 80 < answers.ledger`; in PowerShell, use
`./evalgrid.exe check --min 80 < answers.ledger`. Input redirection is
provided by the shell on all three systems.

For example, with a built executable and a ledger file:

```sh
./evalgrid check --min 80 < answers.ledger
```

## Commands

```text
usage: evalgrid score [--metric NAME]
       evalgrid passk K
       evalgrid show ID
       evalgrid check [--min PCT]
```

The first argument selects a case-sensitive command. `help`, `-h`, or
`--help` as the first argument prints the synopsis; `--version` prints
`evalgrid 0.1.0`. Either informational form ignores later arguments.
Without a command, the synopsis is written to standard error.

- `score` prints one row per model with samples, in bytewise model-name
  order. By default it reports sample count, distinct question count and
  mean EM, F1 and ROUGE-L. `--metric em|f1|rouge` selects one metric.
  A valid ledger without answers produces no output.
- `passk K` reports the number of eligible questions, skipped questions and
  mean pass@k for each model. K accepts one or two ASCII digits and must be
  1–99; output uses the canonical integer K. Questions with fewer than K
  samples are skipped. If none are eligible, the value is `-`.
- `show ID` prints the reference text and then each answer for that question
  in ledger order, with integer EM and four-decimal F1 and ROUGE-L. A
  syntactically valid ID without a reference produces a no-case error.
- `check` reports each model's mean EM as a percentage to two decimals.
  `--min PCT` sets the threshold (0–100; default 50). The comparison uses
  the exact percentage before rounding. A below-threshold model is marked
  `below`; otherwise it is `ok`. With no answers it prints `no answers`.

Command-line arguments are validated before the ledger is read. Words after
the command are classified left to right: a word beginning with `-` and
containing more than one character is an option; every other word is
positional, including `-` and the empty word. The first word that breaks a
rule determines the error. Only `score` accepts `--metric`, and only
`check` accepts `--min`; other options are unknown. A repeated option is
rejected when encountered. An option consumes the next word as its value
even if that word begins with `-`; a missing next word or an empty value is
rejected. Extra positional arguments are rejected as encountered. Once all
words have been read, validation proceeds in this order: report a missing
required positional argument (`passk` needs K and `show` needs ID), validate
the option value, then validate K or ID. `--metric` accepts only `em`, `f1`
or `rouge`. `--min` must contain one to three ASCII digits and have a numeric
value no greater than 100; leading zeros are allowed. K must contain one or
two ASCII digits and have a value from 1 through 99. Error forms include
`unknown command WORD`, `unknown option WORD`, `WORD given twice`,
`WORD needs a value`, `unexpected argument WORD`, `passk needs K`,
`show needs ID`, `bad metric TEXT`, `--min needs an integer from 0 to 100`,
`bad k TEXT`, and `bad id TEXT`.

## Ledger format

The ledger is read as bytes from standard input. Each line is separated by
LF; one CR immediately before LF is removed. A trailing final LF does not
create an extra line, and a nonempty unterminated final line is included.
Every byte must be a tab or printable ASCII (space through tilde). Other
bytes, including non-ASCII bytes, controls and an extra CR, cause
`line N: bad character` before any other interpretation of that line.

Words are separated by runs of spaces and tabs. Blank lines and lines whose
first word begins with `#` are ignored. A comment or blank line is still
subject to byte validation. All other lines have one of these forms:

```text
ref ID TEXT
ans ID MODEL TEXT
```

IDs are 1–20 characters from lowercase ASCII letters, digits, underscore and
hyphen. Model names are 1–40 characters from ASCII letters, digits, period,
underscore, colon, slash and hyphen. Reference and answer text is the
remainder after the fixed fields, trimmed of leading and trailing spaces or
tabs; internal blanks are preserved. Text is at most 60 characters. Reference
text cannot be empty; answer text may be empty. Each ID has exactly one
reference. Answers may precede their reference, but after parsing the whole
ledger every answer must have one. Thus a later syntax error takes precedence
over a missing-reference error.

Ledger validation is performed in line order, and the first invalid line
stops parsing. After byte validation, each line is split into words; a blank
line or comment is ignored, otherwise the first word must be the exact kind
`ref` or `ans`. For a `ref` line, checks are ordered: an ID must be present,
the ID must be valid, the reference text must be nonempty, the text must not
exceed 60 characters, and that ID must not already have a reference. For an
`ans` line, checks are ordered: an ID must be present, the ID must be valid,
a model name must be present, the model name must be valid, and the text must
not exceed 60 characters. Answer text may be empty. Only after every line
passes those checks are answers checked for references, in answer-line order;
the first answer whose ID has no reference produces the missing-reference
error. Consequently, later syntax errors take precedence over missing
references.

Ledger failures are reported as `line N: MESSAGE`. The possible messages
are `unknown line kind WORD`, `ref needs an id`, `bad id WORD`,
`ref ID needs text`, `text too long`, `ref for ID defined twice`,
`ans needs an id`, `ans needs a model name`,
`bad model name WORD`, and `no reference for ID`. The show command may
also report `no case ID` after the complete ledger has parsed. A standard
input read failure is reported as `cannot read standard input`.

## Normalization and metrics

Text is converted to tokens by lowercasing ASCII A–Z, deleting every ASCII
punctuation character, splitting on runs of spaces and tabs, and removing
the tokens `a`, `an` and `the`. Punctuation deletion occurs before
splitting: for example, a hyphenated word becomes one token.

- EM is 1 when the token lists are identical and 0 otherwise; two empty lists
  match.
- Token F1 uses multiset overlap C (the sum of each token's minimum count).
  For nonempty lists it is 2C divided by the combined list lengths. Two empty
  lists score 1; exactly one empty list scores 0.
- ROUGE-L uses the same formula, with C equal to the longest common
  subsequence length.
- For n samples with c exact matches, pass@k is
  1 − choose(n − c, k) / choose(n, k); the quotient is zero when n − c < k.
  It is computed exactly, not using floating point.

Score means are over all of a model's samples. Pass@k means are over eligible
questions only. Metrics and means are calculated as exact integer/rational
values, and only the final display is rounded, half up. Scores use four
decimal places; check percentages use two. Models are sorted by byte order,
which places capital letters before lowercase letters. The show command keeps
sample order from the ledger.

## Output, errors and statuses

Successful commands write results only to standard output and no standard
error. Errors write no standard output and one line beginning
`evalgrid: ` to standard error. Every output line ends in LF. Exit status is
0 for success, 1 for ledger-dependent errors, 2 for command-line errors, and
3 when check finds at least one model below the threshold (its report remains
on standard output). A check with no answers exits 0. If standard input
cannot be read, the entry point writes `evalgrid: cannot read standard input`
to standard error and exits 1.

## Examples

These examples use POSIX shell quoting and the executable built above.

```text
$ printf 'ref q1 The Eiffel Tower\nref q2 a blue whale\nans q1 gpt-x eiffel tower\nans q1 gpt-x Tower in Paris\nans q2 gpt-x whale\nans q2 mini the blue whale\nans q1 mini Eiffel Tower\n' | ./evalgrid score
gpt-x samples 3 cases 2 em 0.3333 f1 0.6889 rouge 0.6889
mini samples 2 cases 2 em 1.0000 f1 1.0000 rouge 1.0000
```

```text
$ printf 'ref q1 The Eiffel Tower\nref q2 a blue whale\nans q1 gpt-x eiffel tower\nans q1 gpt-x Tower in Paris\nans q2 gpt-x whale\nans q2 mini the blue whale\nans q1 mini Eiffel Tower\n' | ./evalgrid passk 1
gpt-x cases 2 skipped 0 pass@1 0.2500
mini cases 2 skipped 0 pass@1 1.0000
```

```text
$ printf 'ref q1 The Eiffel Tower\nref q2 a blue whale\nans q1 gpt-x eiffel tower\nans q1 gpt-x Tower in Paris\nans q2 gpt-x whale\nans q2 mini the blue whale\nans q1 mini Eiffel Tower\n' | ./evalgrid show q1
ref The Eiffel Tower
gpt-x em 1 f1 1.0000 rouge 1.0000
gpt-x em 0 f1 0.4000 rouge 0.4000
mini em 1 f1 1.0000 rouge 1.0000
```

```text
$ printf 'ref q1 The Eiffel Tower\nref q2 a blue whale\nans q1 gpt-x eiffel tower\nans q1 gpt-x Tower in Paris\nans q2 gpt-x whale\nans q2 mini the blue whale\nans q1 mini Eiffel Tower\n' | ./evalgrid check --min 60; echo "exit $?"
gpt-x em 33.33% below
mini em 100.00% ok
exit 3
```

## Development

Run the project checks from the repository root:

```sh
gofmt -l .
go vet ./...
go test ./...
```

The module declares Go `1.27.0` in `go.mod`. This conflicts with the
requested Go 1.22 minimum; explicit resolution is required before claiming
Go 1.22 support. Do not infer Go 1.22 compatibility from a successful build
with another toolchain.

## License

This project is distributed under the [MIT License](LICENSE).

---

Built with [BAKED](https://github.com/Gulliby-development/baked), an open-source coding agent (bench round R86).
