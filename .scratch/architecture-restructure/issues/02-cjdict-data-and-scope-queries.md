# 02: cjdict — the data package and scope queries

**What to build:** A package that owns the Cangjie code table and the character frequency
table, and answers two questions: which Chars can be typed using only a given KeySet, and
which Codes a given Char accepts. It must keep **every** Code a Char accepts — the obvious
shape, a map holding one Code per Char, would mark a correct answer wrong for the 7,883 Chars
that accept more than one, which is the single most damaging bug available here. A malformed
row is skipped silently rather than killing the host program, so construction cannot fail.
Scope results are memoised, because a menu calls them as the learner ticks boxes.

The package must stay free of UI dependencies, which is why KeySet is declared here rather
than beside the Letter table: declaring it next to the Letters would drag the terminal UI
framework into the data layer.

Note that the two data files currently sit uncommitted in the working tree; they land with
this ticket, and every count below was measured against them.

**Blocked by:** None (can start immediately)

**Status:** ready-for-agent

- [ ] A Char accepting several Codes reports all of them
- [ ] Golden tests assert 92,865 Chars, 2,912 with a frequency entry, and 7,883 with more than one Code
- [ ] Golden tests assert per-Group pools: 哲理類 42, 筆劃類 73, 人體類 15, 字形類 26, and 514 for 哲理類+筆劃類
- [ ] The default candidate pool is Chars with a frequency entry; the full dictionary remains reachable for custom scopes
- [ ] A malformed row is skipped without failing construction
- [ ] The package imports no terminal UI library
- [ ] A repeated query for the same KeySet is served from the memo rather than rescanned
