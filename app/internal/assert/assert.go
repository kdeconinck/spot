// =====================================================================================================================
// = LICENSE:       Copyright (c) 2026 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package assert provides small helper functions for writing tests.
//
// The helpers in this package are intentionally minimal and work with Go's standard [testing] package without
// introducing external dependencies.
//
// Instead of accepting [testing.TB], the helpers use a small interface ([TB]).
// The [testing.TB] interface includes an unexported method, which prevents custom implementations outside the standard
// library.
// By accepting a minimal interface, the helpers remain compatible with [testing.T] and [testing.B] while still allowing
// strict test doubles.
package assert

import _ "testing"

// TB defines the minimal subset of methods from [testing.TB] that are required by this package.
//
// Both [testing.T] and [testing.B] satisfy this interface, allowing the same helpers to be used in tests and
// benchmarks.
type TB interface {
	// Helper marks the calling function as a test helper function.
	//
	// When a test fails, the file and line number reported will refer to the caller of the helper rather than the
	// helper itself.
	Helper()

	// Fatalf reports a formatted failure and stops test execution immediately.
	//
	// Fatalf is equivalent to Logf followed by FailNow in the testing package.
	Fatalf(format string, args ...any)
}
