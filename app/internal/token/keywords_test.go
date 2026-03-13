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

// Verifies that [token.LookupKeyword] reports the token kind for reserved Spot DSL keywords and rejects non-keywords.
func Test_LookupKeyword(t *testing.T) {
	t.Parallel()

	for tcName, tc := range map[string]struct {
		input    string
		wantKind token.Kind
		wantOK   bool
	}{
		"When the input matches `LANG`, `LookupKeyword` returns its matching kind.": {
			input:    "LANG",
			wantKind: token.LANG,
			wantOK:   true,
		},
		"When the input matches `EXTENSIONS`, `LookupKeyword` returns its matching kind.": {
			input:    "EXTENSIONS",
			wantKind: token.EXTENSIONS,
			wantOK:   true,
		},
		"When the input matches `SECTION`, `LookupKeyword` returns its matching kind.": {
			input:    "SECTION",
			wantKind: token.SECTION,
			wantOK:   true,
		},
		"When the input matches `DEFINE`, `LookupKeyword` returns its matching kind.": {
			input:    "DEFINE",
			wantKind: token.DEFINE,
			wantOK:   true,
		},
		"When the input matches `CHARSET`, `LookupKeyword` returns its matching kind.": {
			input:    "CHARSET",
			wantKind: token.CHARSET,
			wantOK:   true,
		},
		"When the input matches `VALUES`, `LookupKeyword` returns its matching kind.": {
			input:    "VALUES",
			wantKind: token.VALUES,
			wantOK:   true,
		},
		"When the input matches `RULE`, `LookupKeyword` returns its matching kind.": {
			input:    "RULE",
			wantKind: token.RULE,
			wantOK:   true,
		},
		"When the input matches `MATCH`, `LookupKeyword` returns its matching kind.": {
			input:    "MATCH",
			wantKind: token.MATCH,
			wantOK:   true,
		},
		"When the input matches `ERROR`, `LookupKeyword` returns its matching kind.": {
			input:    "ERROR",
			wantKind: token.ERROR,
			wantOK:   true,
		},
		"When the input matches `LITERAL`, `LookupKeyword` returns its matching kind.": {
			input:    "LITERAL",
			wantKind: token.LITERAL,
			wantOK:   true,
		},
		"When the input matches `SEQUENCE`, `LookupKeyword` returns its matching kind.": {
			input:    "SEQUENCE",
			wantKind: token.SEQUENCE,
			wantOK:   true,
		},
		"When the input matches `ENCLOSED_BY`, `LookupKeyword` returns its matching kind.": {
			input:    "ENCLOSED_BY",
			wantKind: token.ENCLOSED_BY,
			wantOK:   true,
		},
		"When the input matches `MUST_BE_FOLLOWED_BY`, `LookupKeyword` returns its matching kind.": {
			input:    "MUST_BE_FOLLOWED_BY",
			wantKind: token.MUST_BE_FOLLOWED_BY,
			wantOK:   true,
		},
		"When the input matches `CANNOT_BE_FOLLOWED_BY`, `LookupKeyword` returns its matching kind.": {
			input:    "CANNOT_BE_FOLLOWED_BY",
			wantKind: token.CANNOT_BE_FOLLOWED_BY,
			wantOK:   true,
		},
		"When the input is not a reserved keyword, `LookupKeyword` reports false.": {
			input:    "my_rule",
			wantKind: token.Illegal,
			wantOK:   false,
		},
		"When the input differs only by case, `LookupKeyword` reports false.": {
			input:    "lang",
			wantKind: token.Illegal,
			wantOK:   false,
		},
		"When the input is empty, `LookupKeyword` reports false.": {
			input:    "",
			wantKind: token.Illegal,
			wantOK:   false,
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel()

			// Act.
			gotKind, gotOK := token.LookupKeyword(tc.input)

			// Assert.
			assert.Equal(t, tcName, gotKind, tc.wantKind, "Kind")
			assert.Equal(t, tcName, gotOK, tc.wantOK, "OK")
		})
	}
}
