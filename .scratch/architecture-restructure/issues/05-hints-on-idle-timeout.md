# 05: Hints when the learner stalls

**What to build:** When a learner stalls on a Char, reveal the next expected Letter; stall
again and reveal the whole Code. The smallest nudge that unblocks, rather than the answer. How
long counts as stalled is configurable and can be switched off entirely.

A wrong Letter does **not** trigger a hint — it only marks red. Stalling is the single trigger,
because that is the situation a learner can't get out of on their own.

Whether a hint was used travels through to the Char's result, so a game can score it; the typer
itself records nothing. Time comes from an injectable clock, which is what makes this testable
without real waiting.

**Blocked by:** 04 (the hint flag travels through the Char's result)

**Status:** ready-for-agent

- [ ] Pausing for the configured delay reveals the next expected Letter
- [ ] Stalling again reveals the full Code
- [ ] A wrong Letter marks red without revealing anything
- [ ] Setting the delay to zero disables hints
- [ ] A Char's result records that a hint was used
- [ ] Hint timing is tested against a fake clock, with no real waiting in the test suite
