# 04: cjtrace serves a Round from a fixed scope

**What to build:** The drill. A Round of N Chars sampled from a scope, served one at a time,
where a wrong commit re-serves the same Char until the learner gets it right — so finishing a
Round means having actually typed every Char in it correctly. Ships the four traditional
Groups, uniform sampling as the default and frequency-weighted as a named alternative, and the
rule for what happens after a wrong commit as a replaceable policy.

It emits a message when a Round starts, when each Char is served, when each Char is finished
and when the Round completes, carrying raw counts rather than computed accuracy or speed so a
consumer can define those itself. The served-Char message is load-bearing rather than
convenience: the typer never learns the glyph, so this message is the only way anything
downstream can know what to draw.

It holds configuration but no lifecycle — no idle or finished state, and no start message.
Starting a Round is a method, callable again for the next one, so configuration survives
across Rounds while Round state is replaced. The host runs a fixed scope so this is demoable
before the menu exists.

**Blocked by:** 02, 03

**Status:** ready-for-agent

- [ ] Running the binary gives a 10-Char Round of 哲理類 with a summary at the end
- [ ] A wrong commit re-serves the same Char; the Round ends only once every Char is correct
- [ ] The number of attempts per Char appears in the result
- [ ] Uniform is the default Picker, and the frequency-weighted one can be selected instead
- [ ] A supplied Picker overrides the default, making a Round reproducible
- [ ] The emitted messages carry everything needed to render a Round without the built-in view
- [ ] The view is empty before a Round starts
- [ ] A second Round can be started on the same value, keeping its configuration
- [ ] cjtrace tests drive it with key presses against a small fake dictionary
