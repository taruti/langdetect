package langdetect

type commonWordMatcherState struct {
	total        int
	totalMatches int
	counts       [0xFF]byte
}

func (c *commonWordMatcherState) addWord(word []byte) {
	vector, found := commonWordMap[string(word)]
	c.total++
	if found {
		for _, lc := range vector {
			count := c.counts[lc]
			if count < 0xFF {
				c.totalMatches++
				c.counts[lc] = count + 1
			}
		}
	}
}

// Score for language with number "num", lower is better, 0..1.
func (c *commonWordMatcherState) score(num uint8) float64 {
	if c.totalMatches < 5 {
		return 1.0
	}
	return 1 - (float64(c.counts[num]) / float64(c.totalMatches))
}

func lowerkillpunctuation(bs []byte) {
	skip := 0
	for i, c := range bs {
		if skip > 0 {
			skip--
			continue
		}
		switch {
		case c == '\'':
		case c < 'A':
			bs[i] = ' '
		case c < 'Z':
			bs[i] |= 32
		case c < 'a':
			bs[i] = ' '
		case c < 'z':
		case c < 0x80:
			bs[i] = ' '
		case c < 0xC0:
			skip = 1
		case c < 0xE0:
			skip = 2
		case c < 0xF0:
			skip = 3
		default:
			skip = 4
		}
	}
}

func (cms *commonWordMatcherState) processText(bs []byte) {
	tgt := make([]byte, len(bs))
	copy(tgt, bs)
	lowerkillpunctuation(tgt)
	ws := 0
	for i, c := range tgt {
		if c == ' ' {
			if ws != i {
				cms.addWord(tgt[ws:i])
			}
			ws = i + 1
		}
	}
}
