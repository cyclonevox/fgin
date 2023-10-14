//go:build go1.20
// +build go1.20

package util

import "unsafe"

// B2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
func B2s(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
