#!/usr/bin/env bash

# header.sh
# Adds SPDX license header to all .go files that are missing it.

set -euo pipefail

HEADER="// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors
"

ROOT="${1:-.}"
ADDED=0
SKIPPED=0

while IFS= read -r -d '' file; do
  if grep -q "SPDX-License-Identifier" "$file"; then
    ((SKIPPED++))
    continue
  fi

  tmp=$(mktemp)
  printf '%s\n' "$HEADER" | cat - "$file" > "$tmp"
  mv "$tmp" "$file"

  echo "Added header: $file"
  ((ADDED++))
done < <(find "$ROOT" -name "*.go" -type f -print0)

echo ""
echo "Done. Added: $ADDED | Skipped (already had header): $SKIPPED"
