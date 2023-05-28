// Package goarabic contains utility functions for working with strings.
package goarabic

import (
	"strings"
	"testing"
)

// Reverse returns its argument string reversed rune-wise left to right.
func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Crowdbotics", "scitobdworC"},
		{"Hello, 世界", "界世 ,olleH"},
		{"نص عربي", "يبرع صن"},
		{"نَصٌ عَربِيٌّ", "ٌّيِبرَع ٌصَن"},
		{"نَصٌ عَربِيٌّ!", "!ٌّيِبرَع ٌصَن"},
		{"نَصٌ example, عَربِيٌّ!", "!ٌّيِبرَع ,elpmaxe ٌصَن"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	input := strings.Repeat("النص", 100)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Reverse(input)
	}
}

func TestRemoveTashkeel(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"نَصٌ عَربِيٌّ", "نص عربي"},
		{"نص عربي", "نص عربي"},
		{"", ""},
	}
	for _, c := range cases {
		got := RemoveTashkeel(c.in)
		if got != c.want {
			t.Errorf("RemoveTashkeel(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkRemoveTashkeel(b *testing.B) {
	input := strings.Repeat("نَصٌ عَربِيٌّ", 100)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = RemoveTashkeel(input)
	}
}

func TestSmartLength(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"نَصٌ عَربِيٌّ", 7},
		{"نص عربي", 7},
		{"Hello, world", 12},
		{"Hello, 世界", 9},
		{"", 0},
	}
	for _, c := range cases {
		got := SmartLength(&c.in)
		if got != c.want {
			t.Errorf("SmartLength(...) got %d, want %d", got, c.want)
		}
	}
}

func TestToGlyph(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"تجربة النص العربي", "\ufe97\ufea0\ufeae\ufe91\ufe94 \u0627\ufedf\ufee8\ufeba \u0627\ufedf\ufecc\ufeae\ufe91\ufef2"},
		{"", ""},
	}
	for _, c := range cases {
		got := ToGlyph(c.in)
		if got != c.want {
			t.Errorf("ToGlyph(...) got %q, want %+q", got, c.want)
		}
	}
}

func TestSubstringsAllahReplacedToUnicodeCharacterToGlyph(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"الله", "\ufdf2"},
		{"عبد الله أبو عبد الرحمن", "\ufecb\ufe92\ufeaa \ufdf2 \u0623\ufe91\ufeee \ufecb\ufe92\ufeaa \u0627\ufedf\ufeae\ufea3\ufee4\ufee6"},
	}
	for _, c := range cases {
		got := ToGlyph(c.in)
		if got != c.want {
			t.Errorf("ToGlyph(...) got %q, want %+q", got, c.want)
		}
	}
}

func TestLamAlefToGlyph(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"لا", "\ufefb"},
		{"لاي", "\ufefb\ufef1"},
	}
	for _, c := range cases {
		got := ToGlyph(c.in)
		if got != c.want {
			t.Errorf("ToGlyph(...) got %q, want %+q", got, c.want)
		}
	}
}

func BenchmarkToGlyph(b *testing.B) {
	input := strings.Repeat("النص", 100)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ToGlyph(input)
	}
}

func TestRemoveTatweel(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"نـــص عــربــي", "نص عربي"},
		{"نـــَصٌ عَـربي", "نَصٌ عَربي"},
		{"", ""},
	}
	for _, c := range cases {
		got := RemoveTatweel(c.in)
		if got != c.want {
			t.Errorf("RemoveTatweel(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRemoveAllNonArabicChars(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"عــربـabcـينـــص", "عــربــينـــص"},
		{"عــربــينwo%%rd_ـــصa", "عــربــينـــص"},
		{"", ""},
	}
	for _, c := range cases {
		got := RemoveAllNonArabicChars(c.in)
		if got != c.want {
			t.Errorf("RemoveAllNonArabicChars(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkRemoveAllNonArabicChars(b *testing.B) {
	input := strings.Repeat("عــربــينwo%%rd_ـــصa", 100)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = RemoveAllNonArabicChars(input)
	}
}
