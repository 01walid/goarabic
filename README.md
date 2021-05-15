
# GoArabic
[![Build Status](https://travis-ci.org/01walid/goarabic.svg)](https://travis-ci.org/01walid/goarabic)  [![GoDoc](https://godoc.org/github.com/01walid/goarabic?status.svg)](https://godoc.org/github.com/01walid/goarabic)

A Go Lang package for dealing with Arabic text.

This is an initial work on a set of Go functions developed to enhance Arabic web applications.
It all started when I wanted to develop a self-hosted no-dependency and standalone Arabic Captcha in Go (for fun, as I'm discovering Go), and as expected, the Arabic text wasn't rendered as it should:

![Before processing](https://res.cloudinary.com/walid/image/upload/v1429186546/before_pcyoha.png)

So I started playing with Glyphs, unicode, and the special Arabic rules for joining letters. After a bit of work, I got this:

![Before processing](https://res.cloudinary.com/walid/image/upload/v1429186546/after_cmkukt.png)

shouldn't this be a separated package, along with other functionalities? I guess it should!

## Current functionalities
- Glyph representation of the given Arabic text for images/pdf .. etc generation.
- Strip Tashkeel (Arabic Vowels).
- SmartLengh: return the length of the given Arabic String without considering Tashkeel (Arabic Vowels).
- Strip Tatweel
- rune-wise (UTF-8) reverse of the Arabic text, leaving out the Latin one.

## Todo
- Arabic text normalization (Unshaping)
- Add support for some special Quran text characters in Glyphs
- Richer SmartLengh, ignoring some special characters.
- Slugify: a simple Arabic slug generation. Striping Tashkeel and other special chars.
- Spell numbers in Arabic idiom.
- Present dates in Arabic or Hijri and convert Hijri date into Unix timestamp
And maybe more...

## Usage

### Importing
```go
go get github.com/01walid/goarabic
```
##### Example Usage
```go
package main

import (
    "fmt"
    "github.com/01walid/goarabic"
)

func main() {
    fmt.Println(goarabic.RemoveTashkeel("نًصٌ عَربيُّ"))
    fmt.Println(goarabic.ToGlyph("تجربة النص العربي"))
}

```
## Documentation
Package docs available on [gopkgdoc](https://godoc.org/github.com/01walid/goarabic).

## Contributing
Contributions are greatly appreciated. Please fork this repository, make your changes, and open a pull request. More test cases and considerations might be needed, you can run tests using `go test` for the existing functionalities.

This a [SemVer](http://semver.org/)sioned package.
## License
MIT License.
