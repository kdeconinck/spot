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

// Package token defines the lexical token model used by the Spot DSL scanner and parser.
//
// A token represents one classified unit from a `.spot` source file, such as a keyword, identifier, string literal,
// punctuation mark, or special marker like end-of-file.
package token

// Maps reserved Spot DSL words to their token kinds.
var keywords = map[string]Kind{
	"LANG":                  LANG,
	"EXTENSIONS":            EXTENSIONS,
	"SECTION":               SECTION,
	"DEFINE":                DEFINE,
	"CHARSET":               CHARSET,
	"VALUES":                VALUES,
	"RULE":                  RULE,
	"MATCH":                 MATCH,
	"ERROR":                 ERROR,
	"LITERAL":               LITERAL,
	"SEQUENCE":              SEQUENCE,
	"ENCLOSED_BY":           ENCLOSED_BY,
	"MUST_BE_FOLLOWED_BY":   MUST_BE_FOLLOWED_BY,
	"CANNOT_BE_FOLLOWED_BY": CANNOT_BE_FOLLOWED_BY,
}

// LookupKeyword reports the token kind for lexeme.
//
// The boolean result is true when lexeme is a reserved Spot DSL keyword. Non-keyword lexemes should usually be
// classified as [Identifier] by the scanner.
func LookupKeyword(lexeme string) (Kind, bool) {
	kind, ok := keywords[lexeme]

	return kind, ok
}
