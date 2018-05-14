// Package goarabic contains utility functions for working with Arabic strings.
package goarabic

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// SmartLength returns the length of the given string
// without considering the Arabic Vowels (Tashkeel).
func SmartLength(s *string) int {
	// len() use int as return value, so we'd better follow for compatibility
	length := 0

	for _, value := range *s {
		if tashkeel[value] {
			continue
		}
		length++
	}

	return length
}

// RemoveTashkeel returns its argument as rune-wise string without Arabic vowels (Tashkeel).
func RemoveTashkeel(s string) string {
	// var r []rune
	// the capcity of the slice wont be greater than the length of the string itself
	// hence, cap = len(s)
	r := make([]rune, 0, len(s))

	for _, value := range s {
		if tashkeel[value] {
			continue
		}
		r = append(r, value)
	}

	return string(r)
}

// RemoveTatweel returns its argument as rune-wise string without Arabic Tatweel character.
func RemoveTatweel(s string) string {
	r := make([]rune, 0, len(s))

	for _, value := range s {
		if TATWEEL.equals(value) {
			continue
		}
		r = append(r, value)
	}

	return string(r)
}

func getCharGlyph(previousChar, currentChar, nextChar rune) rune {
	glyph := currentChar
	previousIn := false // in the Arabic Alphabet or not
	nextIn := false     // in the Arabic Alphabet or not

	for _, s := range alphabet {
		if s.equals(previousChar) { // previousChar in the Arabic Alphabet ?
			previousIn = true
		}

		if s.equals(nextChar) { // nextChar in the Arabic Alphabet ?
			nextIn = true
		}
	}

	for _, s := range alphabet {

		if !s.equals(currentChar) { // currentChar in the Arabic Alphabet ?
			continue
		}

		if previousIn && nextIn { // between two Arabic Alphabet, return the medium glyph
			for s := range beggining_after {
				if s.equals(previousChar) {
					return getHarf(currentChar).Beggining
				}
			}

			return getHarf(currentChar).Medium
		}

		if nextIn { // beginning (because the previous is not in the Arabic Alphabet)
			return getHarf(currentChar).Beggining
		}

		if previousIn { // final (because the next is not in the Arabic Alphabet)
			for s := range beggining_after {
				if s.equals(previousChar) {
					return getHarf(currentChar).Isolated
				}
			}
			return getHarf(currentChar).Final
		}

		if !previousIn && !nextIn {
			return getHarf(currentChar).Isolated
		}

	}
	return glyph
}

// equals() return if true if the given Arabic char is alphabetically equal to
// the current Harf regardless its shape (Glyph)
func (c *Harf) equals(char rune) bool {
	switch char {
	case c.Unicode:
		return true
	case c.Beggining:
		return true
	case c.Isolated:
		return true
	case c.Medium:
		return true
	case c.Final:
		return true
	default:
		return false
	}
}

// getHarf gets the correspondent Harf for the given Arabic char
func getHarf(char rune) Harf {
	for _, s := range alphabet {
		if s.equals(char) {
			return s
		}
	}

	return Harf{Unicode: char, Isolated: char, Medium: char, Final: char}
}

//RemoveAllNonAlphabetChars deletes all characters which are not included in Arabic Alphabet
func RemoveAllNonArabicChars(text string) string {
	runes := []rune(text)
	newText := []rune{}
	for _, current := range runes {
		inAlphabet := false
		for _, s := range alphabet {
			if s.equals(current) {
				inAlphabet = true
			}
		}
		if inAlphabet {
			newText = append(newText, current)
		}
	}
	return string(newText)
}

// ToGlyph returns the glyph representation of the given text
func ToGlyph(text string) string {
	var prev, next rune

	runes := []rune(text)
	length := len(runes)
	newText := make([]rune, 0, length)

	for i, current := range runes {
		// get the previous char
		if (i - 1) < 0 {
			prev = 0
		} else {
			prev = runes[i-1]
		}

		// get the next char
		if (i + 1) <= length-1 {
			next = runes[i+1]
		} else {
			next = 0
		}

		// get the current char representation or return the same if unnecessary
		glyph := getCharGlyph(prev, current, next)

		// append the new char representation to the newText
		newText = append(newText, glyph)
	}

	return string(newText)
}

//testing first commit by adding new comment//

// RemoveTashkeel returns its argument as rune-wise string without Arabic vowels (Tashkeel).
/*
func RemoveTashkeelExtended(s string) string {
	r := []rune(s)

	m := map[string]bool{"\u064e": true, "\u064b": true, "\u064f": true,
		"\u064c": true, "\u0650": true, "\u064d": true,
		"\u0651": true, "\u0652": true}

	for key, value := range s {
		if m[value] {
			continue
		}
		r[key] = value
	}

	return string(r)
}
*/
