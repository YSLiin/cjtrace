# 08: cjdemo renders the drill from messages alone

**What to build:** A second program that runs the same drill but draws none of it with the
drill's own view — every part of its screen comes from the messages the drill emits. The point
is proof rather than polish: the promise of this architecture is that someone can gamify the
trainer by reacting to what it emits, and this program is the test of that promise. If the
served-Char message were missing or incomplete, this program could not draw a prompt at all,
because the typer never learns the glyph.

Keep it deliberately plain. The game itself is a separate piece of work; this is the smallest
program that exercises the whole message contract.

**Blocked by:** 04

**Status:** ready-for-agent

- [ ] Runs a complete Round without calling the drill's view
- [ ] Draws the current Char from the served-Char message
- [ ] Shows something per-Char taken from the finished-Char message, such as attempts
- [ ] Shows a closing summary built from the Round-completed message
- [ ] Installs as its own binary alongside the practice one
