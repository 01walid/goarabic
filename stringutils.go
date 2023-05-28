// Package goarabic contains utility functions for working with Arabic strings.
package goarabic

import (
	"strings"
	"unicode/utf8"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	var b strings.Builder
	b.Grow(len(s) * 2)

	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		b.WriteRune(r)

		s = s[:len(s)-size]
	}

	return b.String()
}

// SmartLength returns the length of the given string
// without considering the Arabic Vowels (Tashkeel).
func SmartLength(s *string) int {
	// len() use int as return value, so we'd better follow for compatibility
	length := 0

	for _, value := range *s {
		if _, ok := tashkeel[value]; ok {
			continue
		}
		length++
	}

	return length
}

// RemoveTashkeel returns its argument as rune-wise string without Arabic vowels (Tashkeel).
func RemoveTashkeel(s string) string {
	var b strings.Builder
	b.Grow(len(s) * 2)

	for _, value := range s {
		if _, ok := tashkeel[value]; ok {
			continue
		}
		b.WriteRune(value)
	}

	return b.String()
}

// RemoveTatweel returns its argument as rune-wise string without Arabic Tatweel character.
func RemoveTatweel(s string) string {
	var b strings.Builder
	b.Grow(len(s) * 2)

	for _, value := range s {
		if tatweelRune == value {
			continue
		}
		b.WriteRune(value)
	}

	return b.String()
}

func getCharGlyph(previousChar, currentChar, nextChar rune) rune {
	glyph := currentChar
	previousIn := false // in the Arabic Alphabet or not
	nextIn := false     // in the Arabic Alphabet or not

	if _, ok := runeAlphabetMap[previousChar]; ok {
		// previousChar in the Arabic Alphabet ?
		previousIn = true
	}

	if _, ok := runeAlphabetMap[nextChar]; ok {
		// nextChar in the Arabic Alphabet ?
		nextIn = true
	}

	if _, ok := runeAlphabetMap[currentChar]; ok {
		// currentChar in the Arabic Alphabet ?

		if previousIn && nextIn { // between two Arabic Alphabet, return the Medium glyph
			if _, ok := runeBeginningAfter[previousChar]; ok {
				return getHarf(currentChar).Beginning
			}
			return getHarf(currentChar).Medium
		}

		if nextIn { // Beginning (because the previous is not in the Arabic Alphabet)
			return getHarf(currentChar).Beginning
		}

		if previousIn { // Final (because the next is not in the Arabic Alphabet)
			if _, ok := runeBeginningAfter[previousChar]; ok {
				return getHarf(currentChar).Isolated
			}
			return getHarf(currentChar).Final
		}

		if !previousIn && !nextIn {
			return getHarf(currentChar).Isolated
		}
	}

	return glyph
}

// getHarf gets the correspondent Harf for the given Arabic char
func getHarf(char rune) Harf {
	if h, ok := runeAlphabetMap[char]; ok {
		return h
	}
	return Harf{Unicode: char, Isolated: char, Medium: char, Final: char}
}

// RemoveAllNonArabicChars deletes all characters which are not included in Arabic Alphabet
func RemoveAllNonArabicChars(text string) string {
	var b strings.Builder
	b.Grow(len(text) * 2)

	for len(text) > 0 {
		current, size := utf8.DecodeRuneInString(text)
		if _, ok := runeAlphabetMap[current]; ok {
			b.WriteRune(current)
		}
		text = text[size:]
	}

	return b.String()
}

// ToGlyph returns the glyph representation of the given text
func ToGlyph(text string) string {
	var b strings.Builder
	b.Grow(len(text) * 2)

	var previousChar, currentChar, nextChar rune
	var size int

	for len(text) > 0 {
		currentChar, size = utf8.DecodeRuneInString(text)

		// currentChar == alif
		if currentChar == '\u0627' {
			// change الله to unicode character ﷲ
			if len(text) >= 8 && text[:8] == "\u0627\u0644\u0644\u0647" {
				b.WriteRune('\ufdf2')
				previousChar = '\u0647'
				//totalSize += 8
				text = text[8:]
				continue
			}
		}

		text = text[size:]

		nextChar, _ = utf8.DecodeRuneInString(text)

		curHarf := runeAlphabetMap[currentChar]
		nextHarf := runeAlphabetMap[nextChar]
		if curHarf == LAM && nextHarf == ALEF {
			// replace lam and alef to lam alef rune
			b.WriteRune(LAM_ALEF.Unicode)
			previousChar = LAM_ALEF.Unicode
			text = text[2:]
			continue
		}

		// get the current char representation or return the same if unnecessary
		glyph := getCharGlyph(previousChar, currentChar, nextChar)

		// write the new char representation
		b.WriteRune(glyph)

		previousChar = currentChar
	}

	return b.String()
}
