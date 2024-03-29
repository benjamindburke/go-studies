package main

import "fmt"

func main() {
	////////
	// Strings

	// text strings are conventionally interpreted as UTF-8-encoded sequences of Unicode code points (runes).
	s := "hello, world"

	// the built-in len function returns the number of bytes (not runes) in a string
	fmt.Println(len(s)) // "12"

	// and the index operation s[i] accesses the i-th byte of string s where 0 < i < len(s)
	// notice that not all Unicode runes are exactly one byte long, which is why the distinction
	// between runes and bytes is important
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

	// if accessing an out of range byte, go panics
	// c := s[len(s)] // causes panic

	// substring operations s[i:j] yield a new string consisting of they bytes of the original string starting at
	// index i and continuing up to, but not including, the byte at index j, and the result contains j-i bytes.
	// go also panics if j < i or either index is out of bounds.
	fmt.Println(s[0:5]) // "hello"
	fmt.Println(s[:5])  // "hello"
	fmt.Println(s[7:])  // "world
	fmt.Println(s[:])   // "hello, world"

	// String values are immutable: the byte sequence contained in a string value can never be changed.
	// To append one string to another, for instance we can use a temporary variable to hold the old string's memory.
	s = "left foot"
	t := s

	// this does not modify the string that s originally held but causes s to hold the new string formed by the += statement
	s += ", right foot"

	fmt.Println(s) // "left foot, right foot"

	// meanwhile, t holds onto the old string value
	fmt.Println(t) // "left foot"

	// modifications in place obviously are not allowed with immutable string values:
	// s[0] = 'L' // compile error: cannot assign to s[0]

	// Immutability means that it is safe for two copies of a string to share the same underlying memory, making it cheap
	// to copy strings of any length. Similarly, a string and substring s[7:] may safely share the same data, so the substring
	// operation is also cheap. No new memory is allocated in either case.

	////////
	// String Literals

	// Arbitrary bytes can also be included in literal strings using hexadecimal or octal escapes.
	// A hexadecimal escape is written \xhh, with exactly two hexadecimal digits h in upper or lower case.
	// An octal escape is written \ooo with exactly three octal digits o (0 through 7) not exceeding \377
	// Both denote a single byte with the specified value.
	h := '\xFF'
	fmt.Println(h) // "255"
	o := '\377'
	fmt.Println(o) // "255"

	// A raw string literal is written `...`, using backquotes instead of double quotes.
	// Within a raw string literal, no escape sequences are processed; the contents are taken literally,
	// including backslashes and newlines, so a raw string literal may spread over several lines in the program source.
	// The only processing is that carriage returns are deleted so that the value of the string is the same on all platforms,
	// including those that conventionally put carriage returns in text files.

	// Raw literals are a convenient way to compose regular expressions, which tend to have a lot of backslashes.
	fmt.Println(`
		abcdefghijklmnopqrstuvwxyz
	`)
}
