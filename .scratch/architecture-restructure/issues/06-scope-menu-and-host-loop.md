# 06: Scope menu and the menu-to-Round loop

**What to build:** The feature this whole restructure was for. A screen where the learner ticks
any combination of the four traditional Groups and sees how many Chars that combination makes
available, then practises exactly those. When the Round finishes the menu comes back, so the
learner can widen the scope or go again without restarting the program.

Combinations matter more than the individual Groups: the 15 unions of four Groups form the
entire difficulty ladder from 15 Chars up to 2,872, which is why there are no separate
cumulative presets. One total count is shown rather than a count per Group, since the Groups
are additive.

The menu emits its own selection message and knows nothing about the drill — the host turns a
selection into a Round. That keeps the dependency pointing one way and means a game can replace
the menu wholesale.

**Blocked by:** 04

**Status:** ready-for-agent

- [ ] Any combination of the four Groups can be ticked
- [ ] The total shown matches the measured unions — 42 for 哲理類, 514 for 哲理類+筆劃類, 2,872 for all four
- [ ] Confirming a selection starts a Round over exactly that scope
- [ ] The menu returns when the Round completes, and another Round can be run immediately
- [ ] Selecting nothing cannot start an empty Round
- [ ] The menu component does not depend on the drill engine
