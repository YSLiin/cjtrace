# 07: Random and hand-picked scopes

**What to build:** Two more ways to choose what to practise, beyond the four Groups: draw N
Letters at random, or type in the Letters you want.

Random draws need a floor guard. A random 6-Letter set yields anywhere from 19 to 57 practisable
Chars depending on which Letters come up, so a draw whose pool is too small to fill a Round is
rejected and redrawn rather than served — otherwise the learner gets the same three Chars over
and over.

Hand-picked Letters are also the only route to a specific gap: 難 belongs to none of the four
Groups, so the Chars whose Codes need it are unreachable until a learner can name their own
Letters.

**Blocked by:** 06

**Status:** ready-for-agent

- [ ] N Letters can be drawn at random and practised
- [ ] A draw too small to fill a Round is redrawn rather than served
- [ ] Letters can be entered by hand, by key or by glyph
- [ ] A hand-picked set including 難 serves the Chars that need it
- [ ] Input that names no valid Letters is rejected without crashing
