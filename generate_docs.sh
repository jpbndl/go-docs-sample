#!/bin/bash
set -e

SRC_BASE="./docs/localhost:6060"
DEST_DIR="./out"

HTML_SRC="$SRC_BASE/pkg/jayps.com/go-docs/user"
LIB_SRC="$SRC_BASE/lib"

echo "Cleaning $DEST_DIR..."
rm -rf "$DEST_DIR"
mkdir -p "$DEST_DIR"

echo "Copying index.html..."
cp "$HTML_SRC/index.html" "$DEST_DIR/"

echo "Copying lib directory..."
cp -r "$LIB_SRC" "$DEST_DIR/"

echo "Fixing relative paths in index.html..."

# Replace '../../../../lib/godoc/' with 'lib/godoc/' in index.html
sed -i.bak 's|\.\./\.\./\.\./\.\./lib/godoc/|lib/godoc/|g' "$DEST_DIR/index.html"

echo "Removing backup file..."
rm "$DEST_DIR/index.html.bak"

echo "Done. Files ready in $DEST_DIR"
