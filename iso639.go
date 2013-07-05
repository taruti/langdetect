package langdetect

type Language struct {
	code [2]byte
	transient uint8
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
