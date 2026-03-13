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

import "strconv"

// Kind identifies the lexical class of a scanned token.
type Kind uint16

const (
	// Illegal identifies a token that could not be classified by the scanner.
	Illegal Kind = iota

	// EOF marks the logical end of the source stream.
	EOF

	// Identifier identifies a non-keyword name.
	Identifier

	// String identifies a quoted string literal.
	String

	// LBrace identifies the `{` punctuation token.
	LBrace

	// RBrace identifies the `}` punctuation token.
	RBrace

	// LBracket identifies the `[` punctuation token.
	LBracket

	// RBracket identifies the `]` punctuation token.
	RBracket

	// LParen identifies the `(` punctuation token.
	LParen

	// RParen identifies the `)` punctuation token.
	RParen

	// Comma identifies the `,` punctuation token.
	Comma

	// Range identifies the `..` punctuation token.
	Range

	// LANG identifies the `LANG` keyword.
	LANG

	// EXTENSIONS identifies the `EXTENSIONS` keyword.
	EXTENSIONS

	// SECTION identifies the `SECTION` keyword.
	SECTION

	// DEFINE identifies the `DEFINE` keyword.
	DEFINE

	// CHARSET identifies the `CHARSET` keyword.
	CHARSET

	// VALUES identifies the `VALUES` keyword.
	VALUES

	// RULE identifies the `RULE` keyword.
	RULE

	// MATCH identifies the `MATCH` keyword.
	MATCH

	// ERROR identifies the `ERROR` keyword.
	ERROR

	// LITERAL identifies the `LITERAL` keyword.
	LITERAL

	// SEQUENCE identifies the `SEQUENCE` keyword.
	SEQUENCE

	// ENCLOSED_BY identifies the `ENCLOSED_BY` keyword.
	ENCLOSED_BY

	// MUST_BE_FOLLOWED_BY identifies the `MUST_BE_FOLLOWED_BY` keyword.
	MUST_BE_FOLLOWED_BY

	// CANNOT_BE_FOLLOWED_BY identifies the `CANNOT_BE_FOLLOWED_BY` keyword.
	CANNOT_BE_FOLLOWED_BY
)

// Maps each token kind to its readable debug name.
//
// The table is indexed directly by [Kind] so token kinds remain efficient integer values internally while still being
// easy to inspect in tests, logs, and diagnostics.
var kindNames = [...]string{
	Illegal:               "ILLEGAL",
	EOF:                   "EOF",
	Identifier:            "IDENTIFIER",
	String:                "STRING",
	LBrace:                "LBRACE",
	RBrace:                "RBRACE",
	LBracket:              "LBRACKET",
	RBracket:              "RBRACKET",
	LParen:                "LPAREN",
	RParen:                "RPAREN",
	Comma:                 "COMMA",
	Range:                 "RANGE",
	LANG:                  "LANG",
	EXTENSIONS:            "EXTENSIONS",
	SECTION:               "SECTION",
	DEFINE:                "DEFINE",
	CHARSET:               "CHARSET",
	VALUES:                "VALUES",
	RULE:                  "RULE",
	MATCH:                 "MATCH",
	ERROR:                 "ERROR",
	LITERAL:               "LITERAL",
	SEQUENCE:              "SEQUENCE",
	ENCLOSED_BY:           "ENCLOSED_BY",
	MUST_BE_FOLLOWED_BY:   "MUST_BE_FOLLOWED_BY",
	CANNOT_BE_FOLLOWED_BY: "CANNOT_BE_FOLLOWED_BY",
}

// String returns the readable name of k.
//
// It is intended for diagnostics, debug output, and test failures.
// If k is outside the known token kind range, String returns a fallback value that includes the numeric token ID.
func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}

	return "Kind(" + strconv.FormatUint(uint64(k), 10) + ")"
}
