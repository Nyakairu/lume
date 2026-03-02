# Homebrew Tap Setup

This directory contains the Homebrew formula template for Lume.

## How it works

1. Create a public repository: `Tyooughtul/homebrew-tap`
2. Add a GitHub secret `HOMEBREW_TAP_TOKEN` to the **lume** repo (a PAT with `repo` scope for `homebrew-tap`)
3. When you push a tag (e.g. `git tag v1.2.0 && git push --tags`), the release workflow will:
   - Build arm64 + amd64 binaries
   - Create a GitHub Release
   - Auto-update the formula in `homebrew-tap` with correct version + SHA256

## Manual formula update

If you need to update the formula manually:

```bash
VERSION="1.2.0"
SHA256_ARM64=$(shasum -a 256 lume-darwin-arm64 | awk '{print $1}')
SHA256_AMD64=$(shasum -a 256 lume-darwin-amd64 | awk '{print $1}')

sed -e "s/VERSION_PLACEHOLDER/${VERSION}/g" \
    -e "s/SHA256_ARM64_PLACEHOLDER/${SHA256_ARM64}/g" \
    -e "s/SHA256_AMD64_PLACEHOLDER/${SHA256_AMD64}/g" \
    Formula/lume.rb > /tmp/lume.rb
```

## Users install with

```bash
brew install Tyooughtul/tap/lume
```

Or:

```bash
brew tap Tyooughtul/tap
brew install lume
```
