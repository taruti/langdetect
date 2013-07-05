package langdetect

type Language struct {
	code         [2]byte
	transient    uint8
	_placeholder byte
}

// The language code as a string.
func (l Language) String() string {
	return string(l.code[:])
}

// Language name in English (may change between versions).
func (l Language) EnglishName() string {
	if l.code == [2]byte{} {
		return ""
	}
	return langEnglish[l.Number()]
}

// Language native name (may change between versions).
func (l Language) NativeName() string {
	if l.code == [2]byte{} {
		return ""
	}
	return langNative[l.Number()]
}

// A number identifying the language
func (l Language) Number() uint8 {
	return l.transient
}

func (l Language) IsZero() bool {
	return l.code == [2]byte{}
}

// Read a language code from a string to a Language, returns a zero Language on error.
func ReadLanguageString(code string) Language {
	if len(code) == 2 {
		for i := 0; i < len(lcodesString); i += 2 {
			if code == lcodesString[i:i+2] {
				return Language{[2]byte{code[0], code[1]}, byte(i / 2), 0}
			}
		}
	}
	return Language{}
}
