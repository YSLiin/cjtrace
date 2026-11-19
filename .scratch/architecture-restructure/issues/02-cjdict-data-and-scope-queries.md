# 02: cjdict — the data package and scope queries

**What to build:** A package that owns the Cangjie code table and the character frequency
table, and answers two questions: which Chars can be typed using only a given KeySet, and
which Codes a given Char accepts. It must keep **every** Code a Char accepts — the obvious
shape, a map holding one Code per Char, would mark a correct answer wrong for the 2,288 Chars
that accept more than one, which is the single most damaging bug available here. A malformed
row is not handled: the data is generated and trusted, and `go generate` fails first on bad
input, so construction does no validation. Scope queries are not memoised: a scan is fast enough for a
menu that queries as the learner ticks boxes (see README Benchmarks). The parsed tables are
built once on first use.

The package must stay free of UI dependencies, which is why KeySet belongs to this package's
API (passed as a string of Keys, per GLOSSARY) rather than beside the Letter table: owning it
next to the Letters would drag the terminal UI framework into the data layer.

The data files `cjdict/cjcode.csv` and `cjdict/freq.csv` are committed golden data and every
count below was measured against them; the upstream source files are not tracked (see README).

**Blocked by:** None (can start immediately)

**Status:** completed

- [x] A Char accepting several Codes reports all of them
- [x] Golden tests assert 20,902 Chars, 5,705 with a frequency entry, and 2,288 with more than one Code
- [x] Golden tests assert per-Group pools:
    - the full table - 哲理類: 243, 筆劃類: 360, 人體類: 58, 字形類: 99, 哲理類+筆劃類: 3206
    - the default pool - 哲理類: 66, 筆劃類 108, 人體類: 19, 字形類: 39, 哲理類+筆劃類: 877
- [x] The default candidate pool is Chars with a frequency entry; the full dictionary remains reachable for custom scopes
- [x] Generated data is trusted: `go generate` fails on bad input, and load skips nothing
- [x] The package imports no terminal UI library
- [x] Parsed tables are built once; benchmarks cover a query on each pool and the one-off load
