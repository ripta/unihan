// The archive version is pinned deliberately.
// Deriving it from unicode.Version would tie output to the Go toolchain.
//go:generate go run ../../cmd/generate ../../tmp/Unihan_17.0.0.zip ./data.gen.go

package data
