package langdetect

import "bytes"
import "io"
import "strings"
import "testing"

func TestHybridToUtf8(t *testing.T) {
	raw := "Testing äää and in iso \xE4\xE4\xE4 and Go."
	fin := "Testing äää and in iso äää and Go."
	res := new(bytes.Buffer)
	h, _, _ := HybridToUtf8("ISO-8859-15", strings.NewReader(raw))
	io.Copy(res, h)
	if fin != res.String() {
		t.Fatal("Hybrid conversion failure!")
	}
}
