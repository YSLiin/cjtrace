# Spec: module boundaries for cjtrace

Status: ready-for-agent

## Problem Statement

I have a working basic typer. It turns my keystrokes into Cangjie Letters, accepts the
Letter glyphs directly as well, caps input at five, and can print a debug line. It builds
and it runs. That is the whole program, and it is as far as this shape of code goes.

It doesn't know what I am supposed to be typing. There is no target Char, so nothing can
tell me whether what I typed is right — the red marking I have fires when the input
*reaches five Letters*, which is a stand-in for incorrectness rather than the real thing.
And there is no Cangjie data in the project at all: no Codes for characters, no frequency
information, nothing to practise against. So what I have is an input widget, not a trainer.

What I actually want is to practise **one radical group at a time** — 日月金木水火土 first,
then 竹戈十大中一弓, then combinations as I improve — drilling a round of characters that
can be typed using only the Letters I've chosen. There is nowhere in the current code to
put that: no notion of a scope, no notion of a round, nothing that picks characters.

The single package is the obstacle. It is named `cjtrace` but contains only the typer, so
the name already describes something it isn't, and the program is one root `main.go`.
There is no seam between typing one character, deciding what to ask next, and holding the
data — which means every one of those three things has to be built in the same place, and
I cannot reason about or test any of them separately.

Two smaller problems I want fixed while I'm in there: the typer resets to dark-terminal
styling on every keystroke, so it is unreadable in a light terminal and would silently
discard any styling a caller set; and the debug output is an ad-hoc field pair rather than
anything a consumer could rely on.

I also want other people to be able to use this — to embed the typer or the drill in their
own bubbletea program the way `bubbles` components are embedded, to build a game on top by
reacting to what the drill emits, or to just install a binary and practise. One package and
one root `main` leaves no way to offer the typer and the drill separately, and no room for
a second binary alongside the first.

## Solution

Split the spike into three packages with one-way dependencies, each owning one thing:

- **cjdict** owns the data and answers questions about it: which characters can be typed
  using only a given set of Cangjie Letters, and what Codes a character accepts. It has no
  UI dependencies at all.
- **cjtyper** owns typing one Char: translating keys to Letters, capping input at five,
  tracking whether what's typed can still become a valid Code, showing a hint when the
  learner stalls, and committing. It is told the Codes to accept and never learns which
  glyph it is, which is what leaves room for phrases later.
- **cjtrace** owns the drill: the scope presets, how characters are sampled, what happens
  after a wrong answer, and the Round. It emits messages at every step so a game can
  render its own screen and keep its own score.

Two binaries ship: one that is just the menu and the drill, and one example showing
gamification. A menu component selects a scope by multi-selecting the four traditional
Cangjie categories, drawing N random Letters, or entering Letters by hand.

The result: I get the group-scoped practice I wanted, the styling and debug-output problems
are fixed on the way through, and anyone else gets three embeddable pieces or an installable
binary.

## User Stories

### As a learner

1. As a learner, I want to practise only 哲理類 (日月金木水火土), so that I can build muscle
   memory for one coherent group of Letters before adding more.
2. As a learner, I want to select several categories at once, so that I can widen my
   practice gradually from 19 characters up to 5,611 without needing a separate preset for
   every rung.
3. As a learner, I want to see how many characters a scope contains before I start, so that
   I know whether I'm about to drill a pool of 15 or a pool of 500.
4. As a learner, I want a round of a fixed number of characters, so that a practice session
   has a defined end rather than running forever.
5. As a learner, I want every character in my scope to come up over repeated rounds, so that
   I actually learn the whole group instead of seeing the same handful of common characters.
6. As a learner, I want to type a character using any of its valid Codes, so that I'm not
   marked wrong for knowing an alternative spelling.
7. As a learner, I want my keystrokes shown as Cangjie Letters rather than Latin letters, so
   that I'm reading what I'd see in a real input method.
8. As a learner, I want to type the Letter glyphs directly as well as the Latin keys, so that
   the trainer works whichever way I'm inputting.
9. As a learner, I want input capped at five Letters, so that the trainer enforces the same
   limit as real Cangjie.
10. As a learner, I want a wrong Letter marked immediately, so that I notice the error while I
    can still fix it.
11. As a learner, I want to turn that immediate marking off, so that I can practise in a harder
    mode where I only find out at submission.
12. As a learner, I want to correct a wrong Letter before I submit and have it count as a clean
    pass, so that noticing my own mistake is rewarded rather than punished.
13. As a learner, I want a hint when I stall on a character, so that I'm not stuck staring at a
    glyph I can't decompose.
14. As a learner, I want the hint to reveal the next Letter first and the whole Code only if I
    keep asking, so that I get the smallest nudge that unblocks me.
15. As a learner, I want to configure how long "stalled" means, so that the hint matches my pace
    rather than interrupting me.
16. As a learner, I want to submit with the space bar, so that I'm practising the same motion
    real Cangjie typing requires.
17. As a learner, I want a wrong submission to re-serve the same character, so that I leave the
    round having actually typed it correctly.
18. As a learner, I want a summary when the round ends, so that I can see how I did.
19. As a learner, I want the menu again after each round, so that I can change scope or go
    another round without restarting the program.
20. As a learner, I want the trainer readable in my terminal whether it's light or dark, so
    that I can actually see what I'm typing.
21. As a learner, I want to install one binary and start practising, so that I don't have to
    clone and build anything.
22. As a learner, I want to drill an arbitrary set of Letters I choose by hand, so that I can
    target the specific ones I keep getting wrong.
23. As a learner, I want to drill N randomly chosen Letters, so that I can get a varied scope
    without picking one.
24. As a learner, I want a random scope that is never too small to fill a round, so that I'm
    not served the same three characters repeatedly.

### As someone building on this

25. As a game author, I want to embed the drill as a bubbletea component, so that I can build
    my own interface around it the way I would with any `bubbles` component.
26. As a game author, I want a message telling me which Char is now being asked, so that I can
    draw the prompt myself — without it I cannot render a game at all, because the typer never
    learns the glyph.
27. As a game author, I want a message each time a Char is finished, carrying how many attempts
    it took and whether a hint was used, so that I can score it.
28. As a game author, I want a message when a Round starts and when it completes, carrying every
    result, so that I can run my own between-round screens.
29. As a game author, I want raw counts rather than computed accuracy or speed, so that I can
    define those metrics however my game needs.
30. As a game author, I want to supply my own sampling rule, so that I can add spaced repetition
    or suppress characters a player has already mastered across sessions.
31. As a game author, I want to decide what happens after a wrong answer, so that I can advance
    instead of re-serving if my game scores differently.
32. As a game author, I want to ignore the drill's own view entirely and render from its messages
    and accessors, so that my interface isn't constrained by its default look.
33. As a game author, I want the typer embeddable on its own, so that I can build something that
    drills Codes I supply without using the drill engine at all.
34. As a game author, I want to set my own styles and have them survive, so that the component
    matches the rest of my program.
35. As a library consumer, I want the data package free of UI dependencies, so that I can query
    Cangjie data without pulling a terminal UI framework into my program.
36. As a library consumer, I want a bad data row skipped rather than fatal, so that a malformed
    line can't kill my program.
37. As a library consumer, I want scope queries to be fast enough to call from a menu, so that
    selecting a scope feels instant.

### As the maintainer

38. As the maintainer, I want the typer, the drill and the data in separate packages, so that
    I can change one of them without reading the other two.
39. As the maintainer, I want tests I can run without a terminal, so that I can change the
    drill's rules without manually replaying a session.
40. As the maintainer, I want sampling and timing injectable, so that tests are deterministic.
41. As the maintainer, I want the data invariants asserted against the real data, so that a bad
    regeneration fails a test instead of silently degrading practice.
42. As the maintainer, I want the data generators unimportable and uninstallable, so that
    installing the project doesn't put `gencode` on anyone's `$PATH`.
43. As the maintainer, I want one word per concept in the code and the docs, so that the
    vocabulary doesn't drift as the project grows.

## Implementation Decisions

### Packages and dependency direction

Three library packages — **cjdict**, **cjtyper**, **cjtrace** — plus an internal menu
component, two commands, and the existing generators moved under an internal tools
directory. One Go module.

Dependencies point one way: cjtrace depends on cjdict and cjtyper; the menu depends on
cjtrace. **cjtyper imports neither cjdict nor cjtrace**, and **cjdict imports nothing of
ours.**

**KeySet belongs to cjdict, not beside cjtyper's Letter table.** It has no named Go type: cjdict's
queries take it as a string of Keys. This resolves a conflict: cjdict's scope query needs it, and if it lived with the Letters then cjdict
would import cjtyper and drag bubbletea and lipgloss into the data layer. cjtyper has no use
for KeySet — it receives Codes, never scopes.

### Domain vocabulary

| Term | Meaning |
|---|---|
| **Char** (字) | one character — the unit of practice |
| **Code** (倉頡碼) | the full Letter sequence for a Char |
| **Letter** (倉頡字母) | one of the 25, paired 1:1 with a **Key** |
| **KeySet** | a set of Letters, passed as Keys and evaluated as a bitmask |
| **Group** | a named preset KeySet |
| **Round** | an ordered sample of Chars drawn from a scope |
| **Picker** | how a Round's Chars are sampled |
| **Disposition** | what happens after a wrong commit |
| **Attempt** | one commit for a Char |

"Word" is not a term: the data is per-character and no 詞 source exists. "Mistake" is not a
term: a self-corrected keystroke is not a mistake. "Item" is not a term: it was Char under a
second name.

### cjdict

Owns the embedded data and three operations: load, candidates-for-a-KeySet, and codes-for-a-
Char. Chars are keyed by rune; each entry (`CjCode`) carries all the Char's Codes and their
precomputed bitmasks. Frequency counts live in a separate table (`GetFreqs`), and the frequency
rank is dropped as derivable. Codes-for-a-Char is a lookup in the full table. The returned
tables are shared and read-only.

- **All Codes per Char must be preserved.** The obvious shape — a map keyed by Char holding
  one Code — would mark a correct answer wrong for the 2,288 Chars (of 20,902) that accept more than one.
- The data is generated and **trusted**: `go generate` fails on bad input, so library load
  does no row validation and skips nothing.
- The default candidate pool is Chars with both a Code and a frequency entry (5,701 of 20,902). The full
  dictionary stays available for custom scopes.
- Containment is tested as "the Char's Code uses no Letter outside the KeySet". **No index**:
  a scan of the default pool is fast enough to run on every tick of a menu box (see README
  Benchmarks), so queries are **not memoised**; the parsed tables are built once on first use.
  If the full dictionary ever needs interactive queries, the upgrade is complement-plus-inverted-bitmap
  behind this same interface.

### cjtyper

Owns typing one Char. Told the acceptable Codes via a setter; **never told the glyph.** Owns
the Letter table and exports it, so cjtrace gets glyphs for preset labels without a fourth
package.

Behaviour: translate Key to Letter, accept the Letter glyphs as input too, cap at five, track
whether the buffer is still a prefix of *some* acceptable Code, and commit on the commit key.

- Hints fire on **idle timeout only**. A wrong Letter turns the Code red but does not hint.
  Hint content is progressive: next Letter, then the whole Code.
- **No mistake counting.** A Char self-corrected before commit is a clean single pass. This
  removes the viability-transition counting rule entirely.
- Emits a commit message carrying the input, whether it was correct, whether a hint was used,
  and elapsed time; plus a message when a hint is shown.
- **No per-keystroke message**: a game already receives raw key messages from the root, and a
  viability accessor reports buffer health, so a third path would duplicate two.
- The ad-hoc debug field and log string are **dropped**: diagnostics are not part of a
  component's contract, and a consumer that wants them has the messages.

### cjtrace

Owns the drill. Ships the four traditional categories as presets, **multi-selectable** — their
15 unions form the whole progression from 19 to 5,611 Chars, so there are no cumulative
presets and no frequency-sharded ones. 難 is in no category, so Chars needing it are reachable
only through a custom scope.

- Two Pickers ship. **Uniform is the default** — the frequency table has already selected for
  frequency, so weighting the sample by frequency applies the same criterion twice and costs
  25–35% of drill coverage. Uniform is a partial Fisher–Yates; the frequency-weighted one uses
  Efraimidis–Spirakis weighted sampling, which is one pass with no rejection loop — this
  matters because rejection sampling does not terminate in bounded time when drawing 10 Chars
  from 人體類's pool of 15.
- **No randomness option.** Randomness belongs to the Picker, and supplying a Picker already
  gives tests determinism. The shipped Pickers are constructors that accept a source.
- Disposition decides advance-or-retry after a commit; the default is **retry until correct**.
  So the typer reports the verdict and cjtrace decides the consequence — the Code comparison
  lives in exactly one place.
- **No lifecycle states and no start message.** The constructor holds configuration; a start
  method begins a Round and is callable again for the next one, so configuration survives
  across Rounds while Round state is replaced. The host owns "before the first Round" and
  "between Rounds", which is also when it shows the menu. The view is empty when no Round is
  active, and the accessors reduce to active/index/total.

The emitted messages:

| Message | Carries | Why |
|---|---|---|
| Round started | scope, total | lets a game open its own round screen |
| **Char served** | Char, index, total | **load-bearing** — the typer never learns the glyph, so this is a game's only way to know what to draw |
| Char completed | the Char's result | per-Char scoring |
| Round completed | all results, elapsed | end-of-round screens |

A Char's result carries the Char, attempts, whether a hint was used, and elapsed time. **Raw
counts only** — accuracy and speed are derived by the consumer.

Events travel as commands routed through the root: a key press reaches cjtrace, then cjtyper,
which returns a command carrying the commit; the root delivers it back to cjtrace, which
advances. That costs one extra event-loop cycle, accepted so that commits are observable from
outside.

### Menu and binaries

The menu offers category multi-select, random-N-Letters and custom Letter entry, shows one
total pool size (scopes are multi-selectable, so one total rather than per-option counts), and
emits **its own** selection message — it never needs to know cjtrace exists, keeping the
dependency direction clean. The practice binary turns that selection into a start call and
shows the menu again when a Round completes. The demo binary stays a placeholder until the
game is designed.

### API conventions

- **Exported fields** for independent knobs whose zero values are usable: hint delay, the
  wrong-Letter marking flag (default on), the commit key, the Round size, the key map.
- **A setter/getter pair for styles**, because a style change must invalidate render state —
  matching how `textinput` handles it.
- **Options only where the zero value is invalid**: the dictionary, the Picker, the clock, the
  disposition. The rule: field if the zero value is usable and the knob is independent; option
  if construction must establish an invariant.
- Styles use **adaptive** colours, and the error state toggles only the error fields instead of
  reassigning the whole set — the current code's unconditional dark-styles reset is why the
  typer is unreadable in a light terminal.
- **One view method per component**, dispatching internally. Multiple exported view methods
  would force every consumer to replicate the dispatch.
- **Naming**: no `Cj` prefix inside packages — the import path already says Cangjie. Every
  message type ends in `Msg`. No `Get` prefix on accessors. Message names are past
  participles, so the whole set reads as "this happened": `Served`, `Completed`, `Started`.
  This is why the served message is not named for a state like "ready" (an adjective, and
  already used in this repo for a loading flag) or with process vocabulary like "spawned".

## Testing Decisions

A good test here drives a component the way bubbletea drives it — feeding messages in and
asserting on the messages and view that come out — and never reaches for internal state. The
emitted messages *are* the contract a game depends on, so asserting on them is asserting on
external behaviour. There is no prior art: the repo has no tests, so these become it.

**Three seams**, all new, pitched as high as they go:

1. **cjtrace** — the behavioural seam. Driven by real key-press messages with a fake
   dictionary, a fixed Picker and a fixed clock. Covers the Round lifecycle, retry-until-
   correct, the Char-to-Codes handoff, and the four emitted messages.
2. **cjtyper** — the component seam. It is a published API that third parties import
   *without* cjtrace, so its contract is tested directly: keystroke translation, the
   five-Letter cap, viability tracking, the verdict, and idle-timeout hints against a fake
   clock at the layer where the clock actually lives.
3. **cjdict** — the data seam, and unavoidable: its correctness is data-dependent, and a fake
   dictionary at the cjtrace seam cannot prove any of it. Golden tests against the real
   embedded data assert the counts measured during design: 20,902 Chars, 5,705 with frequency
   entries, 2,288 with more than one Code, and the per-Group pool sizes (66, 108, 19, 39,
   and 877 for the 哲理類+筆劃類 union; 243, 360, 58, 99 and 3,206 over the full dictionary).

No seam for the menu: it is UI over KeySet union arithmetic that cjdict's tests already cover.

Determinism comes from the injected Picker and clock, which is the reason those two are
options rather than fields.

The cjtyper seam was initially left as an open question, on the grounds that cjtrace exercises
the typer anyway and fewer seams is better. Breaking the work into tickets settled it: the
typer lands two tickets before the drill exists, so without its own seam those tickets would
ship with no automated proof at all, and hint timing would be tested two layers from its
clock. The delivery argument decided what the API argument had left balanced.

## Out of Scope

- **The demo game.** Its binary ships as a placeholder.
- **Auto-commit on an unambiguous prefix** — committing automatically once the typed prefix can
  only become one Code.
- **Abandoning and restarting a Round.** Deferred as methods the host binds keys to, emitting
  their own messages, rather than a key map inside cjtrace — a key map there would resume
  owning UX decisions and would compete for the keystrokes the typer needs. One semantic
  question is deliberately left for then: does a restart re-serve the same Chars, favouring
  mastery, or resample, favouring variety?
- **The dataset rework**: collapsing the two generators into one tool producing a single
  artifact, keyed per Char with space-separated Codes and a frequency count where zero means
  "no frequency entry" — one file holding all 20,902 Chars at less than today's two files,
  with codepoint, bitmask and frequency-percentage columns dropped as derivable. (The bitmask
  column makes the file both larger *and* slower to load, since computing the mask is a few
  shift-ORs against an integer parse.) Source URLs, retrieval date and corpus total go in
  metadata so percentages stay derivable. The artifact must be **committed**: `go install` has
  no bundling mechanism, so only embedded, committed files reach an installed binary; CI's job
  is to re-download, regenerate and fail on a diff, on a schedule rather than per build.
- **The generator's sanity-check command**, including the codepoint check that the dropped
  column no longer carries.
- **Redistribution terms**, before the first tagged release. The Code source is MIT and needs
  its copyright line and permission text carried in a notice file; the frequency data is not
  explicitly licensed, so the source's terms need checking — noting that raw statistics are
  generally uncopyrightable while a compilation's selection and arrangement may not be.
- **Multi-character phrases.** No 詞 data exists. The typer taking Codes only rather than
  glyphs is what leaves this possible later, but it is a new concept needing a new source.

## Further Notes

### Why the four traditional categories, not frequency-sharded presets

The bitmask was built to make frequency-sharded presets possible, so dropping them deserves
the measurement (default pool of 5,701 Chars; Letter *usage* is the share of running-text
weight carried by Chars whose first Code contains the Letter). Usage is flat — 27% for the
heaviest Letter down to 8% for the lightest, with 難 at 1% — so there is no head to shard on.
More decisively, Letter usage does not predict *typeable* Chars, because a Char needs all its
Letters in scope: the top 6 Letters by usage yield 105 Chars, the coverage-optimal 6 yield 141,
and the traditional 筆劃類 — 7 Letters — yields 108. At 12 Letters, the top 12 by usage yield
883 against 877 for the 14 Letters of 哲理類+筆劃類. The traditional grouping is about as good
as usage ranking by the numbers (a coverage-optimised small set does win, but is an arbitrary
bag of Letters) and better pedagogically: the glyphs in each category share visual logic and
the categories have canonical names.

### The 15 unions

Default pool (Chars with a frequency entry). The Groups' Letters are

- 哲理類: `a`–`g`,
- 筆劃類: `h`–`n`,
- 人體類: `o`–`r`,
- 字形類: `s`–`w` and `y` (卜);

難 (`x`) belongs to no Group. `% running text` is the summed frequency of the pool Chars in scope
over the summed frequency of the whole table.

| scope | Letters | Chars | % running text |
|---|---|---|---|
| 人體類 | 4 | 19 | 2.3 |
| 字形類 | 6 | 39 | 1.4 |
| 哲理類 | 7 | 66 | 1.6 |
| 筆劃類 | 7 | 108 | 5.6 |
| 人體類+字形類 | 10 | 211 | 6.5 |
| 哲理類+人體類 | 11 | 245 | 7.8 |
| 筆劃類+人體類 | 11 | 556 | 16.6 |
| 哲理類+字形類 | 13 | 418 | 7.2 |
| 筆劃類+字形類 | 13 | 561 | 15.0 |
| 哲理類+筆劃類 | 14 | 877 | 22.2 |
| 哲理類+人體類+字形類 | 17 | 1131 | 22.1 |
| 筆劃類+人體類+字形類 | 17 | 1754 | 36.8 |
| 哲理類+筆劃類+人體類 | 18 | 2175 | 48.1 |
| 哲理類+筆劃類+字形類 | 20 | 2703 | 50.9 |
| all four | 24 | 5611 | 99.2 |

### Why uniform sampling is the default

| scope | pool | top 10 Chars' share of weight | heaviest ÷ lightest | distinct Chars over 10 Rounds of 10 |
|---|---|---|---|---|
| 哲理類 | 66 | 72% | 3,851× | weighted 28/66 · uniform 53/66 |
| 筆劃類 | 108 | 73% | 22,524× | weighted 33/108 · uniform 67/108 |
| 哲理類+筆劃類 | 877 | 30% | 24,362× | weighted 69/877 · uniform 95/877 |

Distinct Chars are the mean over 300 simulated sessions (seed 42), each of 10 Rounds drawing
10 distinct Chars.

### Data hygiene

Four frequency entries have no Code — `○`, `Y`, `S` and `~` — so the pool is 5,701 of the
5,705 frequency entries; 15,201 Chars have Codes but no frequency entry. The Code generator
keeps only U+4E00–U+9FFF, so nothing is outside the BMP and no Code contains `z`.

### Verification

Build, then the tests. Then by hand: select 哲理類 alone and
confirm the menu reports 66 Chars; run a Round and confirm a wrong Letter marks immediately,
an idle pause reveals the next Letter, a wrong submission re-serves the same Char, and the
menu returns at the end. Multi-select 哲理類+筆劃類 and confirm 877. Run once in a light
terminal and once in a dark one. Install and confirm only the two intended binaries land on
`$PATH`.

### Domain docs this spec produces

`GLOSSARY.md` from the vocabulary table, and three ADRs — each hard to reverse, surprising
without context, and the outcome of a real trade-off:

1. **Traditional categories over frequency-sharded presets** — surprising precisely because
   the bitmask exists to make the sharding possible.
2. **Uniform sampling as the default** — a reader will ask why a typing trainer ignores
   character frequency; the answer is that the pool already encodes it.
3. **cjtrace has no lifecycle** — why there is no start message, no idle view and no restart
   binding, and why the host owns menu-to-Round composition.

### Data files

`cjdict/cjcode.csv`, `cjdict/freq.csv` and both generators (`internal/tools/gencode`,
`internal/tools/genfreq`) are committed; the CSVs are golden data and every count quoted here
was measured against them. The upstream sources (`cj5.json`, `BIAU1.TXT`) are untracked; see
README for regenerating.
