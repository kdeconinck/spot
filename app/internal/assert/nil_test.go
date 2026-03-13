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

// Verify the public API of the assert package.
//
// Tests in this package are written against the exported API only.
// This ensures that validation behavior is tested through the same surface that external consumers would use.
package assert_test

import (
	"testing"

	"github.com/kdeconinck/spot/internal/assert"
)

// Used to construct typed nil pointer values in nil assertion tests.
type nilTestStruct struct {
	// TODO: Intentionally left empty.
}

// Verifies that [assert.Nil] reports failures only when values differ from nil, and that it formats the failure message
// as expected.
func Test_Nil(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		testName   string
		label      string
		gotInput   any
		wantMsg    string
		wantHelp   int
		wantFatalf int
	}{
		"When the value is <nil>, no failure is reported.": {
			testName:   "nil check",
			label:      "Value",
			gotInput:   nil,
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `pointer`, no failure is reported.": {
			testName:   "nil check",
			label:      "Value",
			gotInput:   (*nilTestStruct)(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `slice`, no failure is reported.": {
			testName:   "nil check",
			label:      "Value",
			gotInput:   []int(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `map`, no failure is reported.": {
			testName:   "nil check",
			label:      "Value",
			gotInput:   map[string]int(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `func`, no failure is reported.": {
			testName:   "nil check",
			label:      "Value",
			gotInput:   (func())(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `channel`, no failure is reported.": {
			testName:   "nil check",
			label:      "Value",
			gotInput:   (chan int)(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is different from <nil>, a failure is reported.": {
			testName:   "nil check",
			label:      "Value",
			gotInput:   true,
			wantMsg:    "\n\nUT Name:          nil check\n\033[32mExpected (Value): <nil>\033[0m\n\033[31mActual (Value):   NOT <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the label is empty, the field names do not include a suffix.": {
			testName:   "nil check",
			label:      "",
			gotInput:   true,
			wantMsg:    "\n\nUT Name:  nil check\n\033[32mExpected: <nil>\033[0m\n\033[31mActual:   NOT <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			tbSpy := new(TBSpy)

			// Act.
			assert.Nil(tbSpy, tc.testName, tc.gotInput, tc.label)

			// Assert.
			if tbSpy.failureMsg != tc.wantMsg {
				t.Fatalf("Failure message = %q, want %q", tbSpy.failureMsg, tc.wantMsg)
			}

			if tbSpy.helperCalls != tc.wantHelp {
				t.Fatalf("Helper calls = %d, want %d", tbSpy.helperCalls, tc.wantHelp)
			}

			if tbSpy.fatalfCalls != tc.wantFatalf {
				t.Fatalf("Fatalf calls = %d, want %d", tbSpy.fatalfCalls, tc.wantFatalf)
			}
		})
	}
}

// Verifies that [assert.Nilf] reports failures only when values differ from nil, and that it forwards the
// caller-provided format string and arguments to Fatalf.
func Test_Nilf(t *testing.T) {
	t.Parallel()

	const msgFmt = "Not <nil> - got %v, want <nil>."

	for tcName, tc := range map[string]struct {
		gotInput   any
		wantMsg    string
		wantHelp   int
		wantFatalf int
	}{
		"When the value is <nil>, no failure is reported.": {
			gotInput:   nil,
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `pointer`, no failure is reported.": {
			gotInput:   (*nilTestStruct)(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `slice`, no failure is reported.": {
			gotInput:   []int(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `map`, no failure is reported.": {
			gotInput:   map[string]int(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `func`, no failure is reported.": {
			gotInput:   (func())(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `channel`, no failure is reported.": {
			gotInput:   (chan int)(nil),
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is different from <nil>, a failure is reported.": {
			gotInput:   true,
			wantMsg:    "Not <nil> - got true, want <nil>.",
			wantHelp:   1,
			wantFatalf: 1,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			tbSpy := new(TBSpy)

			// Act.
			assert.Nilf(tbSpy, tc.gotInput, msgFmt, tc.gotInput)

			// Assert.
			if tbSpy.failureMsg != tc.wantMsg {
				t.Fatalf("Failure message = %q, want %q", tbSpy.failureMsg, tc.wantMsg)
			}

			if tbSpy.helperCalls != tc.wantHelp {
				t.Fatalf("Helper calls = %d, want %d", tbSpy.helperCalls, tc.wantHelp)
			}

			if tbSpy.fatalfCalls != tc.wantFatalf {
				t.Fatalf("Fatalf calls = %d, want %d", tbSpy.fatalfCalls, tc.wantFatalf)
			}
		})
	}
}

// Verifies that [assert.NotNil] reports failures only when values are nil, and
// that it formats the failure message as expected.
func Test_NotNil(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		testName   string
		label      string
		gotInput   any
		wantMsg    string
		wantHelp   int
		wantFatalf int
	}{
		"When the value is NOT <nil>, no failure is reported.": {
			testName:   "not nil check",
			label:      "Value",
			gotInput:   true,
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `pointer`, a failure is reported.": {
			testName:   "not nil check",
			label:      "Value",
			gotInput:   (*nilTestStruct)(nil),
			wantMsg:    "\n\nUT Name:          not nil check\n\033[32mExpected (Value): NOT <nil>\033[0m\n\033[31mActual (Value):   <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `slice`, a failure is reported.": {
			testName:   "not nil check",
			label:      "Value",
			gotInput:   []int(nil),
			wantMsg:    "\n\nUT Name:          not nil check\n\033[32mExpected (Value): NOT <nil>\033[0m\n\033[31mActual (Value):   <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `map`, a failure is reported.": {
			testName:   "not nil check",
			label:      "Value",
			gotInput:   map[string]int(nil),
			wantMsg:    "\n\nUT Name:          not nil check\n\033[32mExpected (Value): NOT <nil>\033[0m\n\033[31mActual (Value):   <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `func`, a failure is reported.": {
			testName:   "not nil check",
			label:      "Value",
			gotInput:   (func())(nil),
			wantMsg:    "\n\nUT Name:          not nil check\n\033[32mExpected (Value): NOT <nil>\033[0m\n\033[31mActual (Value):   <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `channel`, a failure is reported.": {
			testName:   "not nil check",
			label:      "Value",
			gotInput:   (chan int)(nil),
			wantMsg:    "\n\nUT Name:          not nil check\n\033[32mExpected (Value): NOT <nil>\033[0m\n\033[31mActual (Value):   <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is <nil>, a failure is reported.": {
			testName:   "not nil check",
			label:      "Value",
			gotInput:   nil,
			wantMsg:    "\n\nUT Name:          not nil check\n\033[32mExpected (Value): NOT <nil>\033[0m\n\033[31mActual (Value):   <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the label is empty, the field names do not include a suffix.": {
			testName:   "not nil check",
			label:      "",
			gotInput:   nil,
			wantMsg:    "\n\nUT Name:  not nil check\n\033[32mExpected: NOT <nil>\033[0m\n\033[31mActual:   <nil>\033[0m\n\n",
			wantHelp:   1,
			wantFatalf: 1,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			tbSpy := new(TBSpy)

			// Act.
			assert.NotNil(tbSpy, tc.testName, tc.gotInput, tc.label)

			// Assert.
			if tbSpy.failureMsg != tc.wantMsg {
				t.Fatalf("Failure message = %q, want %q", tbSpy.failureMsg, tc.wantMsg)
			}

			if tbSpy.helperCalls != tc.wantHelp {
				t.Fatalf("Helper calls = %d, want %d", tbSpy.helperCalls, tc.wantHelp)
			}

			if tbSpy.fatalfCalls != tc.wantFatalf {
				t.Fatalf("Fatalf calls = %d, want %d", tbSpy.fatalfCalls, tc.wantFatalf)
			}
		})
	}
}

// Verifies that [assert.NotNilf] reports failures only when values are nil, and that it forwards the caller-provided
// format string and arguments to Fatalf.
func Test_NotNilf(t *testing.T) {
	t.Parallel()

	const msgFmt = "<nil> - got %v, want NOT <nil>."

	for tcName, tc := range map[string]struct {
		gotInput   any
		wantMsg    string
		wantHelp   int
		wantFatalf int
	}{
		"When the value is NOT <nil>, no failure is reported.": {
			gotInput:   true,
			wantMsg:    "",
			wantHelp:   1,
			wantFatalf: 0,
		},
		"When the value is a typed <nil> `pointer`, a failure is reported.": {
			gotInput:   (*nilTestStruct)(nil),
			wantMsg:    "<nil> - got <nil>, want NOT <nil>.",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `slice`, a failure is reported.": {
			gotInput:   []int(nil),
			wantMsg:    "<nil> - got [], want NOT <nil>.",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `map`, a failure is reported.": {
			gotInput:   map[string]int(nil),
			wantMsg:    "<nil> - got map[], want NOT <nil>.",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `func`, a failure is reported.": {
			gotInput:   (func())(nil),
			wantMsg:    "<nil> - got <nil>, want NOT <nil>.",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is a typed <nil> `channel`, a failure is reported.": {
			gotInput:   (chan int)(nil),
			wantMsg:    "<nil> - got <nil>, want NOT <nil>.",
			wantHelp:   1,
			wantFatalf: 1,
		},
		"When the value is different from <nil>, a failure is reported.": {
			gotInput:   nil,
			wantMsg:    "<nil> - got <nil>, want NOT <nil>.",
			wantHelp:   1,
			wantFatalf: 1,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Arrange.
			tbSpy := new(TBSpy)

			// Act.
			assert.NotNilf(tbSpy, tc.gotInput, msgFmt, tc.gotInput)

			// Assert.
			if tbSpy.failureMsg != tc.wantMsg {
				t.Fatalf("Failure message = %q, want %q", tbSpy.failureMsg, tc.wantMsg)
			}

			if tbSpy.helperCalls != tc.wantHelp {
				t.Fatalf("Helper calls = %d, want %d", tbSpy.helperCalls, tc.wantHelp)
			}

			if tbSpy.fatalfCalls != tc.wantFatalf {
				t.Fatalf("Fatalf calls = %d, want %d", tbSpy.fatalfCalls, tc.wantFatalf)
			}
		})
	}
}
