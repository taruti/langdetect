package langdetect

import (
	"unicode/utf8"

	//	"bytes"
	//	"encoding/xml"
	//	"io"
	//"fmt"
	//"fmt"
)

/*
func Debug(s string, vals ...interface{}) {
	fmt.Printf(s, vals...)
}

type Charset byte

const (
	Utf8 = Charset(iota)
	Ascii
	Utf16Be
	Utf16Le
	Other
)

type DetectFlags uint32

const (
	DetectBOM = DetectFlags(1 << iota)
	DetectLang
	DetectXML
	DetectQP
	DetectBase64
	DetectUrlEncoding
	DetectAll = ^DetectFlags(0)
)
*/

// Detect language in a slice of utf8 data.
func DetectLanguage(utf8slice []byte, originalCharset string) Language {
	var total int
	var sp [maxscript]uint16
	raw := utf8slice
	if len(raw) > 32*1024 {
		raw = raw[:32*1024]
	}
	var i int
	for {
		c, s := utf8.DecodeRune(raw[i:])
		if s == 0 {
			break
		}
		//		if sp[special] > 1000 { // FIXME
		//			Debug("C='%c' %X\n", c, c)
		//		}
		i += s
		switch {
		case c < 0x0009: // special, skip
			sp[special]++
		case c < 0x000E: // space
			sp[space]++
		case c < 0x0020: // special, skip
			sp[special]++
		case c == 0x0020: // space
			sp[space]++
		case c < 0x0041: // special/digit, skip
			sp[special]++
		case c < 0x005B: // Uppercase ascii letter
			sp[ascii]++
		case c < 0x0061: // special, skip
			sp[special]++
		case c < 0x007B: // Lowercase ascii letter
			sp[ascii]++
		case c < 0x00C0: // special, skip
			sp[special]++
		case c < 0x0250: // Latin accented
			sp[latin]++
		case c < 0x0374: // modifier/combining
			sp[special]++
		case c < 0x0400: // Greek/Coptic
			sp[greek]++
		case c < 0x0530: // Cyrillic
			sp[cyrillic]++
		case c < 0x05BE: // Armenian
			sp[armenian]++
		case c < 0x0600: // Hebrew
			sp[hebrew]++
			//		case c < 0x0780: // Arabic/Syrian
			//		case c < 0x07C0: // Thaana
			//		case c < 0x0800: // N'Ko
			//		case c < 0x0840: // Samaritan
			//		case c < 0x08A0: // Mandaic
		case c < 0x0900: // Arabic (extended-A)
			sp[arabicetc]++
		case c < 0x0E00: // Indian scripts (multiple! FIXME?)
			sp[indian]++
		case c < 0x0E80: // Thai
			sp[thai]++
		case c < 0x0F00: // Lao
			sp[lao]++
		case c < 0x1000: // Tibetan
			sp[tibetan]++
		case c < 0x10A0: // Burmese
			sp[burmese]++
		case c < 0x1100: // Georgian
			sp[georgian]++
		case c < 0x1200: // Hangul (Korean)
			sp[korean]++
			//		case c < 0x13A0: // Cherokee, Canadian native, Tagalog, Hanuhoo, Buhid, ...
			//		case c < 0x1680: // Canadian native
		case c < 0x1E00: // Too many to count, ignore, FIXME
			sp[special]++
		case c < 0x1F00: // Latin extended
			sp[latin]++
		case c < 0x2000: // Greek extended
			sp[greek]++
		case c < 0x2C80: // Misc symbols, braille, latin extended C, IGNORE
		case c < 0x2D00: // Coptic
			sp[greek]++
		case c < 0x2D30: // Georgian extra
			sp[georgian]++
		case c < 0x2DE0: // Tifinagh
			sp[tifinagh]++
		case c < 0x2E00: // Cyrillic extended
			sp[cyrillic]++
		case c < 0x2E80: // punctuation, ignore
			sp[special]++
		case c < 0x3040: // CJK
			sp[cjk]++
		case c < 0x3100: // Japanese
			sp[japanese]++
		case c < 0x3130: // Chinese
			sp[cjk]++
		case c < 0x3190: // Korean
			sp[korean]++
		case c < 0xA000: // CJK
			sp[cjk]++
		case c < 0xA720: // Yi
			sp[yi]++
		case c < 0xA800: // Latin extended
			sp[latin]++
		case c < 0xAC00: // FIXME
			sp[special]++
		case c < 0xD7B0:
			sp[korean]++
		default: // higher, ignore
			sp[special]++
		}
		total++
	}
	letters := float64(total - (int(sp[space]) + int(sp[special])))
	half := (uint16(total) - (sp[space] + sp[special])) / 2
	//	Debug("total=%d letters=%d half=%d arr: %v\n", total, half*2, half, sp)
	if sp[greek] > half {
		return Unknown // FIXME
	}
	if sp[cyrillic] > half {
		return Unknown // FIXME
	}
	if sp[armenian] > half {
		return Hy
	}
	if sp[hebrew] > half {
		return Unknown // FIXME
	}
	if sp[arabicetc] > half {
		return Unknown // FIXME
	}
	if sp[indian] > half {
		return Unknown // FIXME
	}
	if sp[thai] > half {
		return Th
	}
	if sp[lao] > half {
		return Lo
	}
	if sp[tibetan] > half {
		return Unknown // FIXME
	}
	if sp[burmese] > half {
		return My
	}
	if sp[georgian] > half {
		return Ka
	}
	if sp[tifinagh] > half {
		return Unknown // FIXME
	}
	// CJK
	totalcjk := float64(sp[cjk])
	if ko := float64(sp[korean]); (ko+totalcjk)/letters > 0.5 && ko/letters > 0.2 {
		return Ko
	}
	if ja := float64(sp[japanese]); (ja+totalcjk)/letters > 0.5 && ja/letters > 0.1 {
		return Ja
	}
	if totalcjk/letters > 0.5 {
		return Zh
	}
	// Yi
	if yi := float64(sp[yi]); yi/letters > 0.5 {
		return Ii
	}
	// Latin+Ascii
	if sp[latin]+sp[ascii] > half {
		return detectLatinByteTrigram(raw)
	}
	return Unknown
}

var Unknown = Language{}

//ko:                                   asc  space spec
//total=6696 letters=1104 half=552 arr: [919 2117 3474 31 8 69 0 5 24 6 3 0 0 0 0 0 0 32 0 0]
//

const (
	ascii = iota
	space
	special
	latin
	greek
	cyrillic
	armenian
	hebrew
	arabicetc
	indian
	thai
	lao
	tibetan
	burmese
	georgian
	korean
	tifinagh
	cjk
	japanese
	yi
	maxscript
)

/*
func Detect(data []byte, flags DetectFlags) Charset {
	orig := data
	cs := Any
	// BOM
	if len(data) >= 3 && (flags&DetectBOM) != 0 {
		switch {
		case data[0] == 0xFF && data[1] == 0xFE:
			return Utf16Be
		case data[0] == 0xFE && data[1] == 0xFF:
			return Utf16Le
		case data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF:
			// wanna check special ascii codings here?
			cs = Utf8
			data = data[3:]
		}
	}

	// Ascii-loop first
	maybeBase64Std := true
	maybeBase64Url := true
	var wordlengths [32]int32
	var curwordlen int32
	for i, c := range data {
		wordchar := false
		switch {
		case (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'):
			c &= 223 // to upper case
			wordchar = true
		case c >= '0' && c <= '9':
		case c == ' ' || c == '\r' || c == '\n':
		case c == '+' || c == '/':
			maybeBase64Url = false
		case c == '-' || c == '_':
			maybeBase64Std = false
		case c <= 0x7F:
			maybeBase64Std = false
			maybeBase64Url = false
		default:
			data = data[i:]
			goto eightbitloop
		}
		if wordchar {
			curwordlen++
		} else if curwordlen > 0 {
			if curwordlen > 31 {
				curwordlen = 31
			}
			wordlengths[curwordlen]++
			curwordlen = 0
		}
	}
	_ = maybeBase64Std
	_ = maybeBase64Url
	return Ascii
eightbitloop:
	var rem int
	for _, c := range data {
		switch {
		case rem > 0:
			if c < 0x80 || c > 0xC0 {
				goto other
			}
			rem--
		case c < 0x7F:
		case c >= 0xC0 && c < 0xDF:
			rem = 1
		case c >= 0xE0 && c < 0xEF:
			rem = 2
		case c >= 0xF0 && c < 0xF4:
			rem = 3
		default:
			goto other
		}
	}
	return Utf8
	_ = orig

	// Xml handling (+html)
	if flags&DetectXML != 0 && len(data) > 8 && data[0] == '<' {
		var buf [8]byte
		copy(buf[:], data[1:])
		for i := range buf {
			buf[i] &= 223
		}

		switch {
		case bytes.HasPrefix(buf[:], []byte("?XML ")) ||
			bytes.HasPrefix(buf[:], []byte("!DOCTYPE")) ||
			bytes.HasPrefix(buf[:], []byte("HTML")):
			data = xmlextract(data)
		}
	}

other:
	return Other
}

func xmlextract(data []byte) []byte {
	res := make([]byte, 0, len(data))
	var d xml.Decoder
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		tok, err := d.RawToken()
		if err != nil {
			return res
		}
		ch, ok := tok.(xml.CharData)
		if ok {
			res = append(res, []byte(ch)...)
		}
	}
}

type Translator interface {
	Translate(data []byte, eof bool) (n int, cdata []byte, err error)
}

type genTransLate struct {
	b []byte
	f func(gt *genTransLate, data []byte, eof bool) (n int, cdata []byte, err error)
}

func utf16BeToUtf8(gt *genTransLate, data []byte, eof bool) (n int, cdata []byte, err error) {

}

type T interface {
	String() string
	ScrapeWordsTo(wordWriter io.Writer, source io.Reader) (bytes int, words int, err error)
	Utf8(io.Reader) io.Reader
	Language() Language
}
*/
