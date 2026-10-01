# Temporary datadog-secrets-mock test-owner backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Complete official module distribution `github.com/DataDog/datadog-agent/comp/core/secrets/mock@v0.73.0-devel.0.20251030121902-cd89eab046d6`, including all original licenses, fixtures and file modes. UPSTREAM-MODULE.json records original Go checksums and ZIP digest; UPSTREAM-SHA256SUMS anchors every original byte. Only matching maintained v2 imports and independently resolvable module metadata differ. Original monorepo sibling replacements cannot resolve outside the original monorepo and are removed. The DataDog definition requirement selects the same cd89 generation used by the executing setup module. No production caller API or YAML node boundary changes.

This owner serves executed engineering/test paths only. Every executing independent consumer explicitly selects it. New owned Go/shell files retain first-party headers and strict checks. Original inventories and exact maintained deltas are enforced by the attribution guard. Remove this backport when a compatible same-generation maintained owning release passes retained and consumer contracts. No product adoption or release is implied.
