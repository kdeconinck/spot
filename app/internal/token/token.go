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

import "github.com/kdeconinck/spot/internal/loc"

// Token represents one lexical unit produced by the Spot DSL scanner.
//
// Kind identifies the token category.
// Lexeme stores the exact source text for the token.
// Span records the half-open source range occupied by the token.
type Token struct {
	// Kind identifies the lexical class of the token.
	Kind Kind

	// Lexeme stores the exact source text from which the token was scanned.
	Lexeme string

	// Span records the source range occupied by the token.
	Span loc.Span
}

// Is reports whether t has the provided [Kind].
func (t Token) Is(kind Kind) bool {
	return t.Kind == kind
}
