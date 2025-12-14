#!/bin/sh
set -e

# Fix git safe.directory for GitHub Actions
# The workspace is mounted at /github/workspace and may have different ownership
git config --global --add safe.directory '*'

exec "$@"
