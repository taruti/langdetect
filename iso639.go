package langdetect

type Language [2]byte

// The language code as a string.
func (l Language) String() string {
	return string(l[:])
}

// Language name in English (may change between versions).
func (l Language) EnglishName() string {
	s, _ := langEnglish[l]
	return s
}

// Language native name (may change between versions).
func (l Language) NativeName() string {
	s, _ := langNative[l]
	return s
}
