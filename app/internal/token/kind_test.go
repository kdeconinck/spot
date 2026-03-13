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
	"github.com/kdeconinck/spot/internal/token"
)

// Verifies that [token.Kind.String] returns the readable name for known token kinds, and a numeric fallback for
// unknown token kinds.
func Test_Kind_String(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		input token.Kind
		want  string
	}{
		"When the kind is `Illegal`, `String` returns its readable debug name.": {
			input: token.Illegal,
			want:  "ILLEGAL",
		},
		"When the kind is `EOF`, `String` returns its readable debug name.": {
			input: token.EOF,
			want:  "EOF",
		},
		"When the kind is `Identifier`, `String` returns its readable debug name.": {
			input: token.Identifier,
			want:  "IDENTIFIER",
		},
		"When the kind is `String`, `String` returns its readable debug name.": {
			input: token.String,
			want:  "STRING",
		},
		"When the kind is `LBrace`, `String` returns its readable debug name.": {
			input: token.LBrace,
			want:  "LBRACE",
		},
		"When the kind is `RBrace`, `String` returns its readable debug name.": {
			input: token.RBrace,
			want:  "RBRACE",
		},
		"When the kind is `LBRACKET`, `String` returns its readable debug name.": {
			input: token.LBracket,
			want:  "LBRACKET",
		},
		"When the kind is `RBRACKET`, `String` returns its readable debug name.": {
			input: token.RBracket,
			want:  "RBRACKET",
		},
		"When the kind is `LPAREN`, `String` returns its readable debug name.": {
			input: token.LParen,
			want:  "LPAREN",
		},
		"When the kind is `RPAREN`, `String` returns its readable debug name.": {
			input: token.RParen,
			want:  "RPAREN",
		},
		"When the kind is `COMMA`, `String` returns its readable debug name.": {
			input: token.Comma,
			want:  "COMMA",
		},
		"When the kind is `RANGE`, `String` returns its readable debug name.": {
			input: token.Range,
			want:  "RANGE",
		},
		"When the kind is `LANG`, `String` returns its readable debug name.": {
			input: token.LANG,
			want:  "LANG",
		},
		"When the kind is `EXTENSIONS`, `String` returns its readable debug name.": {
			input: token.EXTENSIONS,
			want:  "EXTENSIONS",
		},
		"When the kind is `SECTION`, `String` returns its readable debug name.": {
			input: token.SECTION,
			want:  "SECTION",
		},
		"When the kind is `DEFINE`, `String` returns its readable debug name.": {
			input: token.DEFINE,
			want:  "DEFINE",
		},
		"When the kind is `CHARSET`, `String` returns its readable debug name.": {
			input: token.CHARSET,
			want:  "CHARSET",
		},
		"When the kind is `VALUES`, `String` returns its readable debug name.": {
			input: token.VALUES,
			want:  "VALUES",
		},
		"When the kind is `RULE`, `String` returns its readable debug name.": {
			input: token.RULE,
			want:  "RULE",
		},
		"When the kind is `MATCH`, `String` returns its readable debug name.": {
			input: token.MATCH,
			want:  "MATCH",
		},
		"When the kind is `ERROR`, `String` returns its readable debug name.": {
			input: token.ERROR,
			want:  "ERROR",
		},
		"When the kind is `LITERAL`, `String` returns its readable debug name.": {
			input: token.LITERAL,
			want:  "LITERAL",
		},
		"When the kind is `SEQUENCE`, `String` returns its readable debug name.": {
			input: token.SEQUENCE,
			want:  "SEQUENCE",
		},
		"When the kind is `ENCLOSED_BY`, `String` returns its readable debug name.": {
			input: token.ENCLOSED_BY,
			want:  "ENCLOSED_BY",
		},
		"When the kind is `MUST_BE_FOLLOWED_BY`, `String` returns its readable debug name.": {
			input: token.MUST_BE_FOLLOWED_BY,
			want:  "MUST_BE_FOLLOWED_BY",
		},
		"When the kind is `CANNOT_BE_FOLLOWED_BY`, `String` returns its readable debug name.": {
			input: token.CANNOT_BE_FOLLOWED_BY,
			want:  "CANNOT_BE_FOLLOWED_BY",
		},
		"When the kind is outside the known range, `String` returns a numeric fallback.": {
			input: token.Kind(999),
			want:  "Kind(999)",
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Act.
			got := tc.input.String()

			// Assert.
			assert.Equal(t, tcName, got, tc.want, "String")
		})
	}
}
