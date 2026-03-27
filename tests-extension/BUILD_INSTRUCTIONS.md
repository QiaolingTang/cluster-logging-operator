# Building the OTE Test Extension

The migration has successfully completed code generation and test migration. The remaining step is resolving Go module dependencies due to the `github.com/openshift/origin` package having checksum database issues.

## Current Status

✅ **Completed:**
- All test files migrated and annotated
- Main.go generated with proper OTE integration
- Testdata fixtures copied and bindata configuration ready
- go.mod created with dependencies and replace directives
- go.sum partially populated (1150 entries)

⚠️ **Issue:**
The `github.com/openshift/origin` package version is not in the Go checksum database, preventing normal `go build`.

## Solution: Build from Parent Directory

The cluster-logging-operator parent repository already has working Go modules. Build the extension using the parent's module resolution:

```bash
# Navigate to parent repository
cd /Volumes/GoWorkspace/go/src/github.com/openshift/cluster-logging-operator

# Add tests-extension as a Go workspace (Go 1.18+)
cat > go.work << 'WORKSPACE'
go 1.25

use .
use ./tests-extension
WORKSPACE

# Now build from tests-extension
cd tests-extension
go build -o bin/cluster-logging-operator-tests-ext ./cmd
```

## Alternative: Use Vendor Mode (Recommended)

If workspace approach fails, use vendor mode:

```bash
cd tests-extension

# Copy go.sum from parent if needed
cp ../go.sum ./go.sum.parent
cat go.sum.parent >> go.sum
sort -u go.sum -o go.sum

# Vendor dependencies (ignore warnings)
go mod vendor 2>&1 | grep -v "does not contain package"

# Build with vendor
go build -mod=vendor -o bin/cluster-logging-operator-tests-ext ./cmd
```

## Verify the Binary

Once built:

```bash
# Check binary
ls -lh bin/cluster-logging-operator-tests-ext
file bin/cluster-logging-operator-tests-ext

# Test it
./bin/cluster-logging-operator-tests-ext --help
./bin/cluster-logging-operator-tests-ext list | head -10
```

## Troubleshooting

### Error: "404 Not Found" for origin package
- **Cause:** The origin version is from a fork/branch not in the official checksum database
- **Fix:** Use `GOSUMDB=off` or the vendor approach above

### Error: "toolchain checksum disabled"
- **Cause:** GOSUMDB=off prevents toolchain verification
- **Fix:** Use `GOTOOLCHAIN=local` with your installed Go version

### Error: "go.mod requires go >= 1.25"
- **Fix:** Install Go 1.25+ from https://go.dev/dl/

## Next Steps After Build

1. **Test the extension:**
   ```bash
   ./bin/cluster-logging-operator-tests-ext list
   ```

2. **Commit the changes:**
   ```bash
   cd ..
   git add tests-extension/
   git status
   ```

3. **Create PR** with the migrated tests

