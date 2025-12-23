# Releasing Parley

## Quick Release

To create a new release:

```bash
# Patch version (0.1.0 -> 0.1.1)
./scripts/release.sh patch

# Minor version (0.1.0 -> 0.2.0)
./scripts/release.sh minor

# Major version (0.1.0 -> 1.0.0)
./scripts/release.sh major
```

This will:
1. Bump the version in `wails.json`
2. Create a git tag
3. Push to GitHub
4. Trigger GitHub Actions to build binaries for:
   - macOS (Universal - M1 & Intel)
   - Linux (amd64)
   - Windows (amd64)

## Manual Build

To build locally:

```bash
# Development build
wails dev

# Production build (current platform)
wails build

# Build for specific platform
wails build -platform darwin/universal  # macOS
wails build -platform linux/amd64       # Linux
wails build -platform windows/amd64     # Windows
```

## Distribution

Binaries are created in `build/bin/`:
- **macOS**: `Parley.app` (bundled as .zip)
- **Linux**: `parley` binary (bundled as .tar.gz)
- **Windows**: `parley.exe` (bundled as .zip)

## First Release

Before your first release:
1. Create a GitHub repository (already done: `jonnynabors/parley`)
2. Push your code: `git push origin main`
3. Run: `./scripts/release.sh patch` to create v0.1.0

## Requirements

- Go 1.21+
- Node.js 16+
- Wails CLI
- GitHub repository with Actions enabled
