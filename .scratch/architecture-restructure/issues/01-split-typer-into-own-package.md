# 01: Prefactor — split the typer into its own package, move the program under cmd

**What to build:** Nothing changes for the learner — the binary types Cangjie Letters exactly
as it does today. This is pure rearrangement, so that the rest of the work has somewhere to
go. The typer moves out of the package named `cjtrace`, which currently contains nothing else
and so describes something it isn't, into its own `cjtyper` package, exporting the Cangjie
Letter table so other packages can label Groups with glyphs. The program moves from a root
`main` into `cmd/cjtrace`, and the two data generators move under `internal/tools` so they can
neither be imported nor installed by anyone. This frees the name `cjtrace` for the drill
engine, which has nowhere to live while the typer occupies it.

Review this one as a no-op: if behaviour changed, something went wrong.

**Blocked by:** None (can start immediately)

**Status:** completed

- [x] Everything builds, and running the binary types Letters exactly as before the move
- [x] The Cangjie Letter table is exported from `cjtyper`
- [x] Installing every command in the module puts only the practice binary on `$PATH` — not the generators
- [x] The README's generator commands point at their new location
- [x] No behaviour change anywhere
