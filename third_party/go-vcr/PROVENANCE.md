# Temporary go-vcr test-owner backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Complete official module distribution `github.com/dnaeon/go-vcr@v1.2.0`, including all original licenses, fixtures and file modes. UPSTREAM-MODULE.json records original Go checksums and ZIP digest; UPSTREAM-SHA256SUMS anchors every original byte. Only matching maintained v2 imports and independently resolvable module metadata differ. The original Makefile and vendor manifest remain exact. Owned GNUmakefile selects module mode because the original distribution contains vendor metadata without vendor packages. Test matchers use published Gomega v1.38.2 and maintained v3.0.5. No production caller API or YAML node boundary changes.

This owner serves executed engineering/test paths only. Every executing independent consumer explicitly selects it. New owned Go/shell files retain first-party headers and strict checks. Original inventories and exact maintained deltas are enforced by the attribution guard. Remove this backport when a compatible same-generation maintained owning release passes retained and consumer contracts. No product adoption or release is implied.
