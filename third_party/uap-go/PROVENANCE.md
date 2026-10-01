# Temporary UAParser YAML v2 backport

Logical module: `github.com/ua-parser/uap-go`. Original source:
`v0.0.0-20240611065828-3a4781585db6`, commit
`3a4781585db6719d364e35638acf15745b9f7f91` from https://github.com/ua-parser/uap-go.

Go checksum database verified distribution:

- Module sum: `h1:SIKIoA4e/5Y9ZOl0DCe3eVMLPOQzJxgZpfdHHeauNTM=`
- go.mod sum: `h1:BUbeWZiieNxAuuADTBNb3/aeje6on3DhU3rpWsQSB1E=`
- Module ZIP SHA-256: `6b8ab3549d3a28ff142dbfdb8e683aeb0f27dd31331aeb228dc5c4ec91b2ea6e`

All module distribution files, licenses, embedded regex database and upstream tests
are retained. `UPSTREAM-SHA256SUMS` describes the original distribution. The two
YAML imports (production and tests) and module metadata are the only semantic
upstream source changes: matching maintained `go.yaml.in/yaml/v2 v2.4.3`, requiring Go 1.15.
Existing caller APIs, LRU version, database loading and scalar semantics are preserved.
Repository integration adds a Makefile and focused loading/scalar/error tests.
Copied upstream CI files are inactive source evidence, not repository workflows.
The existing repository generation/formatting pass also normalizes import grouping
in `test.go`, variable declaration grouping in `benchmark_test.go`, and field
alignment in `parser.go`. These are cosmetic; no upstream behavior changes.
There are no custom YAML node methods requiring translation.

The omitted upstream Git submodule fixtures are included at the original gitlink
`ua-parser/uap-core df56280c9e2b42dd64be2b750f803c58feb3f94a`. Its licenses and
complete database/test source are retained. `UAP-CORE-SHA256SUMS` records every
fixture file; the retrieved GitHub archive SHA-256 is
`5fe6bebbf2fab92af6537a0cc6f7404f5baf0082fc68dd068963999e9ed064d0`.

As inspected October 1, 2026, the latest owner revision
`v0.0.0-20260529044130-17c35e68e58c` still selects old YAML v3 and changes the
parsing/cache API. It is not the matching-major maintenance fix for this selected source.

Each independent consumer explicitly replaces the UAParser owner module with this
nested source; generated distributions read actual module identities during OCB
replacement generation. This is not a legacy parser module-name replacement.
The module is excluded from contrib release versioning but included in module
checks; it is explicitly tidied first because crosslink omits foreign namespaces. Remove this backport and all consumer replacements
when an independently reviewed compatible maintained owning release preserves
these contracts and passes the retained suites and actual distribution checks.
Product builders currently select upstream contrib; this backport does not propagate
to independently generated product binaries or establish release/deployment adoption.

Tracking: https://github.com/StackVista/stackstate/issues/717

## Required attribution gate

This nested owner's native `checklicense`/`lint` path first runs the fail-closed
`internal/attribution` guard on the host platform, including during Windows
cross-lint. Original inventories are pinned by their manifest SHA-256 values;
every original file is required and checked byte-for-byte. The six reviewed
parser/metadata/cosmetic maintenance files have exact permitted hashes in the
guard. License/copyright assets have no maintenance exception. Any new upstream
change requires explicit review and updating its narrow hash anchor.

`OWNED-FILES` accounts separately for additions. Missing, overlapping, symlinked
or unaccounted files fail. Every owned Go/shell file is passed to the unchanged
common first-party header checker; adding owned source requires inventory
registration and genuine first-party headers. The module stays registered for
native lint. Root/common source selection, parser modules and upstream bytes are
unchanged by this integration. Tests exercise lost/altered attribution, source,
manifest and inventory failures; native header checks also reject headerless
new owned Go/shell files.
