# Temporary scaleway-sdk-go owning-source YAML backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Original logical identity: `github.com/scaleway/scaleway-sdk-go`. Exact selected public distribution:
`v1.0.0-beta.35`, immutable upstream commit
`1fc11916a52027cd170adfc0689a1690009bd0b5`. Module sum `h1:8xfn1RzeI9yoCUuEwDy08F+No6PcKZGEDOQ6hrRyLts=`; go.mod sum `h1:47B1d/YXmSAxlJxUJxClzHR6b3T4M1WyCvwENPQNBWc=`.
Original module ZIP SHA-256: `9259fb6624ef0fa5b2cb9bbbdb16184a47dc5935eb6beb12d9bb208abc749216`. All 244 original
distribution files, licenses, copyright, tests and fixtures are retained;
`UPSTREAM-SHA256SUMS` anchors original bytes. Original nested CI metadata is
inactive source evidence and does not enable repository workflows.

Only matching maintained parser imports, their required module metadata,
compatible existing Testify selection and explicit owning-source consumer links
may differ. Monorepo-relative dependency replacements, when present, are removed
so independent modules resolve the original published requirements; dependencies'
replacements do not propagate. The reviewed policy pins exact allowed Go/metadata
deltas. License/copyright assets have no maintenance exception.

The native attribution guard requires the complete original and owned inventory,
checks every original file and rejects unreviewed deltas. Every owned Go/shell
file retains the common first-party header check, and owned source/documents
retain spell checking. Immutable original source is checked rather than rewritten
by first-party formatting/generation. The module stays registered for native
checks but is excluded from contrib release versioning.

Each independent consumer and actual generated distribution explicitly selects
this source. Remove the temporary owner and its selections after a reviewed
compatible maintained owning release passes the retained upstream/config/API
suites and actual distributions. Product OCB adoption is separate; this source
does not establish released/deployed product qualification.

## Shared patch reconciliation

The patch model was independently reviewed in StackVista/sts-toolbox PR73 at
`12a8f105e2a51ede32cafe95fd0cb1c70eb40c67`, also based on public beta35. This
fork selects beta35 in its actual graphs; no beta36/37 downgrade is applied.
The private nested owner cannot be fetched by the public checksum service and
no approved cross-private Contrib CI source-read path exists. The controller
therefore authorized republishing only the exact public original distribution
and matching maintained-v2 patch here. No private repository files, credentials,
company source or CI access configuration were copied or changed. Reconcile
future updates with the shared owner rather than allowing these patches to drift.
