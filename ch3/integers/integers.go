package main

import "fmt"

func printByte(b byte) {
	fmt.Printf("%08b\n", b)
}

func main() {
	var x uint8 = 1<<1 | 1<<5 // set the uint8 bits in position 1 and 5 "00100010"
	var y uint8 = 1<<1 | 1<<2 // set the uint8 bits in position 1 and 2 "00000110"

	printByte(x)      // "00100010", the set {1, 5}
	printByte(y)      // "00000110", the set {1, 2}
	printByte(x & y)  // "00000010", the intersection {1}
	printByte(x | y)  // "00100110", the union {1, 2, 5}
	printByte(x ^ y)  // "00100100", the symmetric difference {2, 5}
	printByte(x &^ y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i)
		}
	}

	// bit shift all bits one position to the right
	printByte(x << 1) // "01000100", the set {2, 6}

	// bit shift all bits one position to the left
	// bit shifting is equivalent to dividing each bit by 2^(num of bits to shift)
	// by dividing the bit in position 5 (32 or 2^5) by 2^1, we get 2^4, or 16, or a 1 in the 4th position
	// and by dividing the bit in position 1 (2 or 2^1) by 2^1, we get 2^1 or 1, or a 1 in 1st position
	printByte(x >> 1) // "00010001", the set {0, 4}
}
