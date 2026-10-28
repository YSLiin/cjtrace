# Drill scopes are the four traditional Letter Groups, not frequency-ranked shards

The generated Cangjie data carries a per-Code bitmask specifically so that drill scopes could
be built from the most-used Letters — tiers like "the 6 most-used Letters", "the next 6". We
measured it and shipped the four traditional Groups (哲理類, 筆劃類, 人體類, 字形類) instead,
multi-selectable, because frequency ranking turned out to measure the wrong thing.

## Considered Options

Frequency-ranked tiers failed on two counts:

- **There is no head to shard.** Weighted Letter usage across running text runs from 27% for
  the heaviest Letter down to 8% for the lightest, 難 aside — a factor of 3.4, declining
  smoothly. Any cut through it is arbitrary.
- **Letter usage does not predict typeable Chars**, because a Char needs *all* of its Letters
  in scope. The top 6 Letters by usage yield 53 practisable Chars; the coverage-optimal 6
  yield 70; traditional 筆劃類, with 7 Letters, yields 73 — beating the frequency-optimised
  set. At 12 Letters the top-12 yield 517 against 514 for 哲理類+筆劃類.

So the traditional grouping is as good by the numbers, and better otherwise: the glyphs in
each Group share visual logic, the Groups have canonical names, and learners meet them in
that order already.

## Consequences

Making Groups multi-selectable replaced a separate cumulative-progression feature: the 15
unions of four Groups form a ladder from 15 to 2,872 Chars, which includes the cumulative
path (42 → 514 → 1,210 → 2,872) as four rungs on it.

The bitmask remains exactly as useful — scope membership is still a bitmask containment test.
Only the preset table changed.

難 belongs to no Group, so the handful of Chars whose Codes need it are reachable only through
a custom KeySet.
