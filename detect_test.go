package langdetect

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestReadLanguageString(t *testing.T) {
	if ReadLanguageString(En.String()) != En {
		t.Fatal("ReadLanguage string failed!")
	}
}

func TestWikipedia(t *testing.T) {
	f, e := os.Open("test_data/wikipedia")
	if e != nil {
		fmt.Println("NO test_data PRESENT EXITING")
		return
	}
	ns, e := f.Readdirnames(-1)
	if e != nil {
		panic(e)
	}
	for _, n := range ns {
		bs, e := ioutil.ReadFile("test_data/wikipedia/" + n)
		l := n[5:8]
		if l[2] == '.' {
			l = l[:2]
		}
		if e != nil {
			panic(e)
		}
		lang := DetectLanguage(bs, "")
		var alert string
		if l != lang.String() {
			alert = "WRONG ("
			alert += lang.EnglishName()
			alert += ")"
		}
		if lang == Unknown {
			alert = "UNKNOWN ("
			alert += lang.EnglishName()
			alert += ")"
		}
		fmt.Printf("%20s => %v %s %s\n", n, lang, lang.EnglishName(), alert)
	}
}

func XTestCalcTrigrams(t *testing.T) {
	f, e := os.Open("test_data/calc_trigrams")
	if e != nil {
		panic(e)
	}
	ns, e := f.Readdirnames(-1)
	if e != nil {
		panic(e)
	}
	for _, n := range ns {
		bs, e := ioutil.ReadFile("test_data/calc_trigrams/" + n)
		fmt.Println("BUILDING TRIGRAM TABLE: " + n)
		if e != nil {
			panic(e)
		}
		vs := calcLatinByteTrigram(bs)
		lim := 500
		if lim > len(vs) {
			lim = len(vs)
		}
		vs = vs[0:lim]
		x := trigramfreqs(vs)
		sort.Sort(x)
		buf := new(bytes.Buffer)
		buf.WriteString("package langdetect\n\nvar trigram_" + n + " = trigramfreqs{\n")
		for i := 0; i < lim; i++ {
			buf.WriteString(strings.Replace(fmt.Sprintf("%#v,\n", x[i]), "langdetect.", "", -1))
		}
		buf.WriteString("}\n\n")
		ioutil.WriteFile("autogenerated_trigrams_"+n+".go", buf.Bytes(), 0600)
	}
}

func BenchmarkDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trigram_en.Diff(trigram_fr)
	}
}

//func BenchmarkDiffOld(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		trigram_en.DiffOld(trigram_fr)
//	}
//}

func BenchmarkDetectEn(b *testing.B) {
	bs, e := ioutil.ReadFile("test_data/wikipedia/auto_en.txt")
	b.ResetTimer()
	if e != nil {
		fmt.Println("NO test_data PRESENT EXITING")
		return
	}
	for i := 0; i < b.N; i++ {
		DetectLanguage(bs, "")
	}
}

func BenchmarkDetectZh(b *testing.B) {
	bs, e := ioutil.ReadFile("test_data/wikipedia/auto_zh.txt")
	if e != nil {
		fmt.Println("NO test_data PRESENT EXITING")
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DetectLanguage(bs, "")
	}
}
