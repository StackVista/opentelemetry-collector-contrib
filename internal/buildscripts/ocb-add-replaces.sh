#!/usr/bin/env bash

# Copyright The OpenTelemetry Authors
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

DIR="$1"
CONFIG_IN="cmd/$DIR/builder-config.yaml"
CONFIG_OUT="cmd/$DIR/builder-config-replaced.yaml"

cp "$CONFIG_IN" "$CONFIG_OUT"

local_mods=$(find . -type f -name "go.mod" -exec dirname {} \; | sort)
for mod_path in $local_mods; do
    module=$(awk '$1 == "module" {print $2}' "$mod_path/go.mod")
    echo "  - $module => ../../$mod_path" >> "$CONFIG_OUT"
done
go mod edit -json | jq -r '.Replace[] // empty | select(.New.Version != null) |
    "  - " + .Old.Path + (if .Old.Version then " " + .Old.Version else "" end) +
    " => " + .New.Path + " " + .New.Version' >> "$CONFIG_OUT"
echo "Wrote replace statements to $CONFIG_OUT"
