#!/bin/bash
set -e

# Get version from argument or default to patch bump
VERSION_TYPE=${1:-patch}

# Get current version from git tags
CURRENT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")

# Parse version numbers
CURRENT_VERSION=${CURRENT_VERSION#v}
IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT_VERSION"

# Bump version
case $VERSION_TYPE in
  major)
    MAJOR=$((MAJOR + 1))
    MINOR=0
    PATCH=0
    ;;
  minor)
    MINOR=$((MINOR + 1))
    PATCH=0
    ;;
  patch)
    PATCH=$((PATCH + 1))
    ;;
  *)
    echo "Usage: ./scripts/release.sh [major|minor|patch]"
    exit 1
    ;;
esac

NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"

echo "🚀 Releasing version $NEW_VERSION"
echo ""

# Update wails.json version
jq ".info.productVersion = \"${NEW_VERSION#v}\"" wails.json > wails.json.tmp
mv wails.json.tmp wails.json

# Commit version change
git add wails.json
git commit -m "chore: bump version to $NEW_VERSION" || true

# Create and push tag
git tag -a "$NEW_VERSION" -m "Release $NEW_VERSION"
git push origin main
git push origin "$NEW_VERSION"

echo ""
echo "✅ Release $NEW_VERSION initiated!"
echo "📦 GitHub Actions will build and create the release"
echo "🔗 Check progress at: https://github.com/jonnynabors/parley/actions"
