# cjtrace

A 倉頡 (Cangjie) typing trainer. A learner picks a set of Cangjie Letters to practise, and
the trainer serves characters that can be typed using only those Letters.

## Language

### The data

**Char**:
One Chinese character — the unit of practice. A Round serves Chars one at a time.
_Avoid_: Word, 詞, Item, Glyph

**Code**:
The full sequence of Letters that types a Char, such as `ohml`. A Char may accept more than
one Code, and every one of them is correct.
_Avoid_: Shape, Spelling

**Letter**:
One of the 25 倉頡字母, such as 日 or 竹. Letters compose into a Code.
_Avoid_: Radical (that is 部首, a different thing), Shape, Code

**Key**:
The keyboard key that produces a Letter, such as `a` for 日. Letters and Keys correspond
one to one.

### Scoping a drill

**KeySet**:
A set of Letters, held as a bitmask. A Char is typeable within a KeySet when one of its
Codes uses no Letter outside it.
_Avoid_: Bitmask (that is how a KeySet is stored, not what it is), Keymap

**Group**:
One of the four traditional named KeySets — 哲理類, 筆劃類, 人體類, 字形類. Groups are
combinable, and their unions form the whole difficulty progression.
_Avoid_: Category, Preset, Tier, Shard

**Scope**:
The KeySet a Round draws its Chars from. A Scope is always a KeySet; the word names the
role, not a second kind of thing.

### Running a drill

**Round**:
An ordered sample of Chars drawn from a Scope, with a fixed length and a definite end.
_Avoid_: Session, Game, Level, Set

**Picker**:
How a Round's Chars are sampled from the Chars a Scope makes available.
_Avoid_: Sampler, Shuffler, Selector, Chooser

**Disposition**:
What happens to a Char after a wrong Commit — serve it again, or move on.
_Avoid_: Policy, Strategy

**Commit**:
Submitting the typed Code for judgement, by pressing the commit key. One Char may take
several Commits.
_Avoid_: Submit, Enter, Send

**Attempt**:
One Commit for a Char. A Char needing three Attempts was got wrong twice.
_Avoid_: Try, Mistake (a Letter corrected before Commit is not counted at all)

**Verdict**:
Whether a committed Code was correct. The typer decides the Verdict; the drill decides the
consequence.
_Avoid_: Judgement, Validation, Result

**Hint**:
A progressive reveal of the Code for the Char being typed — the next Letter first, the whole
Code on asking again. Triggered by stalling.
