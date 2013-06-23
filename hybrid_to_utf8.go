package langdetect

import (
	"bufio"
	"errors"
	"io"
	"unicode/utf8"
)

// Create a reader from a hybrid of the specified character set and utf8 to utf8.
func HybridToUtf8(name string, rd io.Reader) (utf8reader ReadSetter, charset string, err error) {
	name = NormalizeCharsetName(name)
	switch name {
	case "UTF8":
		return &brdwrap{bufio.NewReader(rd)}, name, nil
	// Translate these all as CP1252
	case "ISO88591", "ISO885915", "WINDOWS1252", "CP1252":
		return &hybrid{sherd{she: &cp1252, rd: bufio.NewReader(rd)}}, name, nil
	default:
		return nil, name, errors.New("Unsupported character set: " + name)
	}
}

type hybrid struct {
	sherd
}

func (s *hybrid) Read(bs []byte) (int, error) {
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
		r, nr, e := s.rd.ReadRune()
		if e != nil {
			return int(n), e
		}
		if r == 0xFFFD && nr == 1 {
			s.rd.UnreadByte()
			c, e := s.rd.ReadByte()
			if e != nil {
				return int(n), e
			}
			if c < 0x80 {
				r = rune(c)
			} else {
				r = rune(s.she[c&0x7F])
				nr = utf8.RuneLen(r)
			}
		}
		if r < 0x80 {
			bs[0] = byte(r)
			bs = bs[1:]
			n++
		} else {
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
