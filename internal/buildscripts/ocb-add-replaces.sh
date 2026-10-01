#!/usr/bin/env bash

# Copyright The OpenTelemetry Authors
# SPDX-License-Identifier: Apache-2.0

set -e

DIR="$1"
CONFIG_IN="cmd/$DIR/builder-config.yaml"
CONFIG_OUT="cmd/$DIR/builder-config-replaced.yaml"

cp "$CONFIG_IN" "$CONFIG_OUT"

local_mods=$(find . -type f -name "go.mod" -exec dirname {} \; | sort)
for mod_path in $local_mods; do
    module=$(awk '$1 == "module" {print $2}' "$mod_path/go.mod")
    echo "  - $module => ../../$mod_path" >> "$CONFIG_OUT"
done
echo "Wrote replace statements to $CONFIG_OUT"
