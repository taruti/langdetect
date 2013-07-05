package main

import (
	"bytes"
	"fmt"
	"github.com/taruti/langdetect"
	"io/ioutil"
	"log"
)

func main() {
	bs, e := ioutil.ReadFile("all.words")
	if e != nil {
		log.Fatal(e)
	}
	bss := bytes.Split(bs, []byte{'\n'})
	m := map[string][]byte{}
	maxl := 0
	for _, x := range bss {
		if len(x) == 0 {
			continue
		}
		ws := bytes.Split(x, []byte{'\t'})
		l := langdetect.ReadLanguageString(string(ws[0]))
		if l.IsZero() {
			log.Fatal("Invalid language:", string(ws[0]))
		}
		w := string(ws[1])
		s, _ := m[w]
		if len(s)+1 >= maxl {
			maxl = len(s) + 1
		}
		m[w] = append(s, l.Number())
	}
	t := "[]byte"
	if maxl <= 8 {
		t = fmt.Sprintf("[%d]byte", maxl)
	}
	fmt.Println("package langdetect\n")
	fmt.Printf("var commonWordMap = map[string]%s{\n", t)
	for k, v := range m {
		var iv interface{} = v
		switch maxl {
		case 1:
			var a [1]byte
			copy(a[:], v)
			iv = a
		case 2:
			var a [2]byte
			copy(a[:], v)
			iv = a
		case 3:
			var a [3]byte
			copy(a[:], v)
			iv = a
		case 4:
			var a [4]byte
			copy(a[:], v)
			iv = a
		case 5:
			var a [5]byte
			copy(a[:], v)
			iv = a
		case 6:
			var a [6]byte
			copy(a[:], v)
			iv = a
		case 7:
			var a [7]byte
			copy(a[:], v)
			iv = a
		case 8:
			var a [8]byte
			copy(a[:], v)
			iv = a
		}
		fmt.Printf("%#v: %#v,\n", k, iv)
	}
	fmt.Println("}")
}
