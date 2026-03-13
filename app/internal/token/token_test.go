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

// Verify the public API of the token package.
//
// Tests in this package are written against the exported API only.
// This ensures that token behavior is tested through the same surface that external consumers would use.
package token_test

import (
	"testing"

	"github.com/kdeconinck/spot/internal/assert"
	"github.com/kdeconinck/spot/internal/loc"
	"github.com/kdeconinck/spot/internal/token"
)

// Verifies that [token.Token.Is] reports whether the token has the requested kind.
func Test_Token_Is(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		input token.Token
		kind  token.Kind
		want  bool
	}{
		"When the token kind matches the requested kind, `Is` returns true.": {
			input: token.Token{
				Kind:   token.String,
				Lexeme: `"json"`,
				Span:   loc.Span{},
			},
			kind: token.String,
			want: true,
		},
		"When the token kind does not match the requested kind, `Is` returns false.": {
			input: token.Token{
				Kind:   token.String,
				Lexeme: `"json"`,
				Span:   loc.Span{},
			},
			kind: token.Identifier,
			want: false,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Act.
			got := tc.input.Is(tc.kind)

			// Assert.
			assert.Equal(t, tcName, got, tc.want, "Is")
		})
	}
}
