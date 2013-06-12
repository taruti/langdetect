package langdetect_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/taruti/langdetect"
	"testing"
)

type MyStruct struct {
	X []byte
}

var someXml = []byte(`<?xml version="1.0" encoding="ISO-8859-15"?><X><X>
Wikipedian kieliversiot toimivat hyvin itsenäisesti. Niillä on omat käytännöt, ja artikkelien sisällön annetaan kehittyä riippumattomasti. Artikkelien ei tarvitse olla käännöksiä toisistaan, ja eri kieliversioihin kirjoitetaankin runsaasti etenkin paikallisluonteista materiaalia, jota ei muunkielisissä Wikipedioissa vielä ole. Muutamat perusperiaatteet, kuten neutraali näkökulma, sitovat kuitenkin kaikkia versioita. Eri kieliversiot linkittävät toisiinsa sivun vasemmassa reunassa sijaitsevien interwiki-linkkien avulla. Lisäksi eri kieliversioilla on runsaasti yhteistä kuva- ja muuta mediamateriaalia.

Alexan mukaan arvioilta 60 prosenttia kaikista Wikipedian käynneistä kohdistuu englanninkieliseen versioon. Loput jakaantuvat muiden kieliversioiden kesken.

Suurin Wikipedia on englanninkielinen, jossa on yli 3 600 000 artikkelia. Saksankielisessä on yli 1 263 000, ranskankielisessä noin 1 130 000.[12] Puolankielisessä Wikipediassa on yli 800 000 artikkelia, japanin-, italian- ja hollanninkielisessä yli 500 000, portugalin- ja espanjankielisessä yli 400 000.Suomen-, Ruotsin- ja venäjänkielisessä on yli 300 000 artikkelia. kiinan- ja norjankielisessä Wikipediassa on yli 200 000 artikkelia, ja volapükin- sekä esperantonkielisessä Wikipediassa on yli 100 000 artikkelia. Yli 10 000 artikkelia on 75 Wikipediassa. Suomenkielinen Wikipedia on suhteellisen suuri verrattuna kieltä puhuvien määrään.[13] Virallisesti Wikipedia on 264 kielellä.
</X></X>`)
var someXmlInIso885915 = bytes.Replace(bytes.Replace(someXml, []byte("ä"), []byte{0xE4}, -1), []byte("ö"), []byte{0xF6}, -1)

func ExampleSniffXmlToUtf8() {
	rd, cs, e := langdetect.SniffXmlToUtf8(bytes.NewReader(someXmlInIso885915))
	if e != nil {
		return
	}
	var v MyStruct
	d := xml.NewDecoder(rd)
	d.CharsetReader = langdetect.FixXmlCharsetReader
	e = d.Decode(&v)
	if e != nil {
		return
	}
	lang := langdetect.DetectLanguage(v.X, cs)
	fmt.Println("ExampleXml:", cs, lang)
	// Output: ExampleXml: ISO885915 fi
}

func TestEmpty(*testing.T) {}
