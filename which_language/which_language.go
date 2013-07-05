package main

import "bytes"
import "fmt"
import "github.com/taruti/langdetect"
import "io"
import "os"

func main() {
	buf := new(bytes.Buffer)
	io.CopyN(buf, os.Stdin, 32*1024)
	if buf.Len() < 10 {
		fmt.Println("Too short input")
		os.Exit(1)
	}
	lang := langdetect.DetectLanguage(buf.Bytes(), "")
	fmt.Printf("%s\t%s\t%s\n", lang, lang.EnglishName(), lang.NativeName())
}
