package main

import (
	"fmt"
    "unicode/utf8"
)

// Thanks to these properties of UTF-8 in Go, many string operations don't require decoding. We can test whether one string
// contains another as a prefix:
func hasPrefix(s, prefix string) bool {
    return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// or as a suffix:
func hasSuffix(s, suffix string) bool {
    return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// or as a substring:
func contains(s, substr string) bool {
    for i := 0; i < len(s); i++ {
        if hasSuffix(s[i:], substr) {
            return true
        }
    }
    return false
}

func main() {
	////////
	// Unicode

	// US-ASCII uses 7 bits to represent 128 "characters": the upper and lower c ase letters of English, digits and
	// a variety of punctuation and device control characters. This leaves much of the world unable to use computer systems in their language.
	// Unicode collects all the characters in all the world's writing systems, plus accents and other diacritical marks,
	// control codes like tab and carriage return, and plenty of esoterica, and assigns each one a standard number called a Unicode code point,
	// or in Go terminology, a rune.

	// Unicode version 8 defines code points for over 120,000 characters in well over 100 languages and scripts.
	// The natural data type for a single rune is int32, which is what Go uses, having the synonym rune for this very reason.
	// Strings could be represented with a sequence of int32 values. In this representation, called UTF-32 or UCS-4, the encoding of
	// each unicode code point has the same size, 32 bits. This is simple and uniform, but uses much more space than necessary
	// since most computer-readable text is in ASCII, which requires only 8 bits/1 byte per character.
	// Still, all characters in widespread use number less than 65535, which would fit in 16 bits.

	// UTF-8 is a variable length encoding of Unicode code points as bytes. UTF-8 uses between 1 and 4 bytes to represent each rune,
	// but only one byte for ASCII characters, and only 2 or 3 bytes for most runes in common use.
	// The high-order bits of the first byte of the encoding for a rune indicate how many bytes follow.

	// A high-order of 0 indicates 7 bit ASCII, where each rune takes only 1 byte, so it is identical to conventional ASCII.
	// A high-order of 110 indicates that the rune takes 2 bytes, the second byte beginning with 10.
	// A high-order of 1110 indicates that the rune takes 3 bytes, all other bytes beginning with 10.o.
	// A high-order of 11110 indicates that the rune takes 4 bytes, all other bytes beginning with 10.

	// 0xxxxxxx								   0 - 127		  (ASCII)
	// 110xxxxx 10xxxxxx					 128 - 2047		  (values  <128 unused)
	// 1110xxxx 10xxxxxx 10xxxxxx           2048 - 65535      (values <2048 unused)
	// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 65536 - 0x10FFFFFF (other values unused)

	// A variable length encoding precludes direct indexing to access the n-th character of a string, but UTF-8 has many desirable properties
	// to compensate. The encoding is compact, compatible with ASCII, and self-synchronizing: it's possible to find the beginning of a character
	// by backing up by no more than 3 bytes. It's also a prefix code, so it can be decoded left to right without any ambiguity or lookahead.
	// No rune's encoding is a substring of any other, or even of a sequence of others, so you can search for a rune by just searching for its bytes.
	// The lexicographic byte order equals the Unicode code point number, so sorting UTF-8 works naturally. There are no embedded NUL (zero) bytes
	// which is convenient for programming languages that use NUL to terminate strings.

	// Unicode escapes in Go string literals allow us to specify them by their numeric code point value. There are two forms:
	// \uhhhh for a 16-bit value and \Uhhhhhhhh for a 32-bit value, where each h is a hexadecimal digit; the need for 32-bits arises very infrequently.
	// Each denotes the UTF-8 encoding of the specified code point. For example, the following literals all represent the same 6 bytes:
	fmt.Println("Hello, 世界")
	fmt.Println("Hello, \xE4\xB8\x96\xE7\x95\x8C")
	fmt.Println("Hello, \u4E16\u754C")
	fmt.Println("Hello, \U00004E16\U0000754C")

    // If we really care about the individual Unicode characters, we have to use other mechanisms.
    // Take for instance the string for the first example:
    s := "Hello, 世界"
    fmt.Println(len(s))
    fmt.Println(utf8.RuneCountInString(s))
    for i := 0; i < len(s); {
        r, size := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%d\t%c\n", i, r)
        i += size
    }

    // Notice that the iterator i is increased by the size of the rune to account for the
    // number of bytes occupied by the rune.

    for i, r := range s {
        fmt.Printf("%d\t%q\t%d\n", i, r, r)
    }

    // or, we could count the number of runes in a string:
    n := 0
    for _, _ = range s {
        n++
    }
    fmt.Println(n)
    // and of course, RuneCountInString does exactly this.
    fmt.Println(utf8.RuneCountInString(s))
    
    s = "プログラム"
    fmt.Printf("% x\n", s) // small note: % x places a space between each hex character
    r := []rune(s)
    fmt.Printf("%x\n", r)

    // if a slice of runes is converted to a string, it produces a concatenation of the UTF-8 encoding of each rune:
    fmt.Println(string(r))
    // likewise:
    fmt.Println(string(65)) // "A"
    fmt.Println(string(0x4eac)) // "京"
    // but if the rune is invalid, the replacement character is substituted:
    fmt.Println(string(1234567))
}

