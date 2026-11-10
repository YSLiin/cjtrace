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
| 哲理類 | 66 | 72% | 3,851× | weighted 28/66 · uniform 53/66 |
| 筆劃類 | 108 | 73% | 22,524× | weighted 33/108 · uniform 67/108 |
| 哲理類+筆劃類 | 877 | 30% | 24,362× | weighted 69/877 · uniform 95/877 |

(Mean of 300 simulated sessions, seed 42; each session is 10 Rounds of 10 distinct Chars.)

Weighted sampling serves 水 and 日 over and over while 汩, 塋, 杲, 鏜, 灝 and 溟 barely appear
— 27–51% fewer distinct Chars practised. The goal of a scoped drill is to cover the scope, not
to reproduce the statistics of running text.

Frequency-weighted sampling still ships as a named Picker, so this is a choice of default
rather than a removal. It uses Efraimidis–Spirakis weighted sampling (key `u^(1/w)`, take the
top N) rather than rejection sampling, which does not terminate in bounded time when a Round
of 10 is drawn from 人體類's pool of 19.
