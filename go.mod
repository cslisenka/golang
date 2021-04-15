// This file is placed in root and declares go module
// Module may have multiple packages
// Module has version and dependencies (with versions)
// Go automatically downloads all dependencies (resolves imports)
// We may get list of dependencies by "go list -m all"
// There is no central repository of go packages (like npm, pip or maven central), package is downloaded from URL
module example.com/hello

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	golang.org/x/tour v0.0.0-20210317163553-0a3a62c5e5c0 // indirect
	rsc.io/quote v1.5.2
)
