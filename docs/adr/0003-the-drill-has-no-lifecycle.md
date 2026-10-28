# The drill holds configuration, not a lifecycle; the host composes menu and Round

`cjtrace` has no Idle or Finished state, no start message, and no key bindings for starting or
restarting. Its constructor holds configuration (dictionary, Picker, clock, Disposition) and a
start method begins a Round, callable again for the next one. Everything before the first Round
and between Rounds belongs to whichever program embeds it.

## Considered Options

The first design gave the drill three states with one view dispatching across them, plus a
`StartMsg` the menu would emit and an exported key map for Start and Restart. Following it
through, each piece turned out to have no owner but the host: the host already decides when to
show the menu, so it already knows when a Round should begin — and since it holds the Model, it
can call a method instead of sending a message.

## Consequences

The state machine did not disappear; it moved to the host, which is the point — a game's
menu-to-Round flow *should* differ from the bundled binary's. The cost is that each host writes
that flow itself, so an identical loop in both binaries would be real duplication worth
noticing.

Two failure modes follow from having no Idle state, and are deliberate: the command returned by
the start method must be propagated or the first Char is never served, and a Model that is never
started renders an empty view with no diagnostic.

Abandoning and restarting a Round, when they arrive, are methods the host binds keys to — not a
key map here. A key map in the drill would resume owning UX decisions and would compete for the
keystrokes the typer needs.
