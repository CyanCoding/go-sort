package gosort

import "fmt"

var alphabet = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

type SequenceKind int

const (
	Alphabetical SequenceKind = iota
	ReverseAlphabetical
	Reverse
)

var sequences = map[SequenceKind]int{
	Alphabetical:        1,
	ReverseAlphabetical: 2,
	Reverse:             3,
}

func SortString(word string) string {
	var byteArray = []byte(word)

	// If we sort the string at least once,
	// we repeat the loop. If we go an entire
	// loop without sorting it's alphabetical
	var sorted = false

	for {
		for i := 1; i < len(byteArray); i++ {
			// If the current char is lower in the alphabet than previous
			if alphabet[string(byteArray[i])] <
				alphabet[string(byteArray[i-1])] {

				// Swap last char with current char and vice versa
				char := byteArray[i-1]
				byteArray[i-1] = byteArray[i]
				byteArray[i] = char
				sorted = true
			}
		}

		if !sorted { // No sorts, so we're done!
			break
		} else {
			sorted = false

		}
	}

	return string(byteArray)

}

func Sort(word string, sort SequenceKind) {
	fmt.Println(sequences[sort])
}
