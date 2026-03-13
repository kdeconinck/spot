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

import (
	"fmt"
	"strings"

	_ "testing"
)

// Formats the failure message for [Equal].
//
// The message consists of three aligned rows:
//
//	UT Name:   <test name>
//	Expected:  <expected value>
//	Actual:    <actual value>
//
// If label is not empty, it is appended to the "Expected" and "Actual" field names, for example "Expected (Item)".
// The values are aligned vertically, and at least one space is always inserted after the colon to avoid unnecessary
// horizontal padding.
func equalMessage[V any](testName, label string, want, got V) string {
	nameKey, expectedKey, actualKey := equalMsgKeys(label)
	maxKeyWidth := maxLen(nameKey, expectedKey, actualKey)

	formatRow := func(key string, value any) string {
		padding := strings.Repeat(" ", maxKeyWidth-len(key)+1)

		return fmt.Sprintf("%s:%s%v", key, padding, value)
	}

	formatRowWithColor := func(color, key string, value any) string {
		return fmt.Sprintf("%s%s\033[0m", color, formatRow(key, value))
	}

	return "\n\n" +
		formatRow(nameKey, testName) + "\n" +
		formatRowWithColor("\033[32m", expectedKey, want) + "\n" +
		formatRowWithColor("\033[31m", actualKey, got) + "\n\n"
}

// Constructs the field names used in the formatted failure message.
//
// The returned keys correspond to the rows:
//
//	UT Name
//	Expected
//	Actual
//
// If label is not empty, it is appended to the Expected and Actual keys in parentheses.
func equalMsgKeys(label string) (nameKey, expectedKey, actualKey string) {
	suffix := ""

	if label != "" {
		suffix = " (" + label + ")"
	}

	return "UT Name", "Expected" + suffix, "Actual" + suffix
}

// Returns the length of the longest string in values.
//
// This is used to compute the alignment width of the message keys so that
// the corresponding values appear in the same column.
func maxLen(values ...string) int {
	max := 0

	for _, v := range values {
		if len(v) > max {
			max = len(v)
		}
	}

	return max
}
