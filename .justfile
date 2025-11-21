set dotenv-load

# Test the Go sources (Units).
test:
    @GOTOOLCHAIN=go1.25.4+auto go test -v -coverprofile=.coverprofile.out ./...
