# A Round samples Chars uniformly, not weighted by character frequency

The Picker defaults to uniform sampling without replacement, even though the project holds
character frequency data and weighting by it was the original intent. The reason is that the
candidate pool is *already* a frequency filter — it is the set of Chars that appear in the
frequency table at all — so weighting the sample by frequency applies the same criterion
twice.

## Consequences

Double-counting frequency measurably costs drill coverage, because a scoped pool is small and
steeply skewed:

| Scope | Pool | Top 10 Chars' share of weight | Heaviest ÷ lightest | Distinct Chars over 10 Rounds of 10 |
| --- | --- | --- | --- | --- |
| 哲理類 | 42 | 73% | 117× | weighted 31/42 · uniform 40/42 |
| 筆劃類 | 73 | 72% | 704× | weighted 35/73 · uniform 55/73 |
| 哲理類+筆劃類 | 514 | 29% | 761× | weighted 76/514 · uniform 95/514 |

Weighted sampling serves 水 and 日 over and over while 灶, 肖, 梢, 棘, 杰 and 炙 barely appear
— 25–35% fewer distinct Chars practised. The goal of a scoped drill is to cover the scope, not
to reproduce the statistics of running text.

Frequency-weighted sampling still ships as a named Picker, so this is a choice of default
rather than a removal. It uses Efraimidis–Spirakis weighted sampling (key `u^(1/w)`, take the
top N) rather than rejection sampling, which does not terminate in bounded time when a Round
of 10 is drawn from 人體類's pool of 15.
