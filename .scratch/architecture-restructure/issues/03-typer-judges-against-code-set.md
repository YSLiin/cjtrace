# 03: The typer judges a Char against its Code set

**What to build:** The first slice where the thing is a trainer rather than an input widget.
The typer is told which Codes to accept and judges what is typed against them — while never
learning which Char it is, which is what leaves room for phrases later. Wrong Letters are
marked as you type, and that marking can be switched off for a harder mode where nothing is
revealed until you commit. Input caps at five Letters. The space bar commits, emitting a
message carrying what was typed, whether it was correct, and how long it took.

Two existing problems get fixed here. Styles become adaptive and settable, replacing the
unconditional dark-terminal reset that makes the typer unreadable in a light terminal and
silently discards anything a caller set; the error state toggles only the error fields rather
than reassigning the whole set. And the ad-hoc debug field and log string are dropped —
diagnostics are not part of a component's contract.

The host wires up one real Char from the dictionary, so this is demoable end to end before any
Round exists.

**Blocked by:** 01 (the package must exist), 02 (for a real Char's Codes)

**Status:** ready-for-agent

- [ ] Running the binary shows one Char, accepts typing, marks wrong Letters, and reports right or wrong on commit
- [ ] Any one of a Char's valid Codes is accepted as correct
- [ ] Marking can be switched off, and then nothing is revealed until commit
- [ ] A Letter corrected before commit leaves no trace in the result — it counts as a clean pass
- [ ] Input cannot exceed five Letters
- [ ] Styles set by a consumer survive subsequent keystrokes
- [ ] Readable in both a light and a dark terminal
- [ ] cjtyper tests cover keystroke translation, the cap, viability tracking and the verdict
- [ ] cjtyper imports neither the data package nor the drill
