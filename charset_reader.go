package langdetect

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"unicode/utf8"
)

// Sniffs the correct character set and returns a utf8 reader,
// the original character set and a possible error.
func SniffXmlToUtf8(rd io.Reader) (utf8reader io.Reader, charset string, err error) {
	bfr := bufio.NewReader(rd)
	b, err := bfr.Peek(64)
	if len(b) < 4 {
		if err == nil {
			err = io.ErrUnexpectedEOF
		}
		return nil, "", err
	}
	var skip [4]byte
	// BOM
	switch {
	case b[0] == 0xFF && b[1] == 0xFE: // UTF16-LE BOM
		return nil, "UTF16LE", errors.New("FIXME UTF16-LE BOM")
	case b[0] == 0xFE && b[1] == 0xFF: // UTF16-BE BOM
		return nil, "UTF16BE", errors.New("FIXME UTF16-BE BOM")
	case b[0] == 0xEF && b[1] == 0xBB && b[2] == 0xBF: // UTF8-BOM
		bfr.Read(skip[0:2])
		return bfr, "UTF8", nil
	case b[0] == 0x3C && b[1] == 0x00: // UTF16 LE <
		return nil, "UTF16LE", errors.New("FIXME UTF16-LE <")
	case b[0] == 0x00 && b[1] == 0x3C: // UTF16 BE <
		return nil, "UTF16BE", errors.New("FIXME UTF16-BE <")
	case bytes.HasPrefix(b, []byte(`<?xml version="1.0" encoding="`)):
		c := b[30:]
		i := bytes.IndexByte(c, '"')
		if i < 1 {
			// This is quite invalid, treat is as utf8 however
			return bfr, "UTF8?", nil
		}
		return charsetToUtf8(string(c[0:i]), bfr)
		// HTML THINGS FIXME
	}
	return bfr, "UTF8?", nil
}

//
func FixXmlCharsetHandler(ignored string, rd io.Reader) (io.Reader, error) {
	return rd, nil
}

// Create a reader from the specified character set to utf8.
func charsetToUtf8(name string, rd io.Reader) (utf8reader io.Reader, charset string, err error) {
	name = NormalizeCharsetName(name)
	switch name {
	case "UTF8":
		return rd, name, nil
	// Translate these all as CP1252
	case "ISO88591", "ISO885915", "WINDOWS1252", "CP1252":
		return &sherd{she: &cp1252, rd: bufio.NewReader(rd)}, name, nil
	default:
		return nil, name, errors.New("Unsupported character set: " + name)
	}
}

// Normalize a character set string (removes all special characters and converts to upper case).
func NormalizeCharsetName(csn string) string {
	res := make([]byte, 0, len(csn))
	for i := 0; i < len(csn); i++ {
		c := csn[i]
		switch {
		case c >= 'a' && c <= 'z':
			c &= 223
			fallthrough
		case (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'):
			res = append(res, c)
		}
	}
	return string(res)
}

type smallhalfencoding [128]uint16

type sherd struct {
	nbuf uint32
	buf  [4]byte
	she  *smallhalfencoding
	rd   *bufio.Reader
}

func (s *sherd) Read(bs []byte) (int, error) {
	n := s.nbuf
	if n > 0 {
		copy(bs, s.buf[:])
		if n >= uint32(len(bs)) {
			s.nbuf -= uint32(len(bs))
			return len(bs), nil
		}
		s.nbuf = 0
		bs = bs[n:]
	}
	for len(bs) > 0 {
		c, e := s.rd.ReadByte()
		if e != nil {
			return int(n), e
		}
		if c < 0x80 {
			bs[0] = c
			bs = bs[1:]
			n++
		} else {
			r := rune(s.she[c&0x7F])
			nr := utf8.RuneLen(r)
			if nr < len(bs) {
				utf8.EncodeRune(s.buf[:], r)
				s.nbuf = uint32(nr)
				return int(n), nil
			}
			utf8.EncodeRune(bs, r)
			n += uint32(nr)
			bs = bs[nr:]
		}
	}
	return int(n), nil
}

/*
type sbuf struct {
	buf bytes.Buffer
	rd  io.Reader
	f   func(data []byte, ateof bool) (ndone int, writeme []byte, err error)
}

func (sb *sbuf) WriteTo(w io.Writer) (int64, error) {
	if sb.f == nil {
		if sb.buf.Len() > 0 {
			w.Write(sb.buf.Bytes())
			sb.buf = nil
		}
		return io.Copy(w, sb.rd)
	}
	total := int64(0)
	for {
		n, writeme, err := sb.f(sb.buf, false)
		if len(writeme) > 0 {
			n, err := w.Write(writeme)
			total += int64(n)
			if err != nil {
				return total, err
			}
		}
		sb.buf = append(sb.buf, make([]byte, (32*1024)*len(sb.buf)))
		n, err := sb.rd.Read(sb.buf)
		sb.buf = sb.buf[0:n]
		if err != nil {
			if err == io.EOF {

			}
			return total, err
		}
	}
}
*/
