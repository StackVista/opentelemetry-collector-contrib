# Explicit shared owner selections

Tracking: https://github.com/StackVista/stackstate/issues/717

## Ordered map

Logical owner `github.com/wk8/go-ordered-map/v2` remains at selected original
v2.1.8, commit `85ca4a2b29d3241fa4513f82be3d38fe2392a791`. The reviewed shared
source is StackVista/stackstate-agent PR535, exact
`d6ac30666e85e56fe8e0b992ebb6d1177a69cdd7`, nested `third_party/go-ordered-map`.
Its canonical addressable replacement requires the `/v2` major suffix:

- Path: `github.com/StackVista/stackstate-agent/third_party/go-ordered-map/v2`
- Go-resolved version: `v2.0.0-20261001130553-d6ac30666e85`
- Module sum: `h1:9OL/wtLq7AITxqTOuAbg6Jrl3KvwK68Fuo0ZR4/wJhc=`
- go.mod sum: `h1:JzMZgMx7Cu1HiXUR9XQoyBhBUf/nKfjsmelqUNnKya8=`

The replacement version's v2.0.0 prefix is the canonical commit-derived pseudo
version, not a downgrade of its exact v2.1.8 source. It preserves ordered JSON
and YAML and coherently selects maintained v3 typed YAML methods. It is selected
explicitly by root and all requiring independent consumers. Actual OCB generation
also carries root's versioned remote owner selections; dependency replacements do
not propagate. This source is reused rather than copied here. Remove the temporary
selection after a compatible maintained owning release passes these consumers.
Product builders must adopt reviewed selections independently; this fork's
selection does not qualify product delivery or deployment.
