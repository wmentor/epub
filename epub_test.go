package epub

import (
	"bytes"
	"testing"
)

func TestReader(t *testing.T) {
	bk, err := Open("./data/test.epub")
	if err != nil {
		t.Fatal(err)
	}

	if len(bk.Files()) != 211 {
		t.Fatal("invalid files counter")
	}

	bk.Close()

	i := 0

	Reader("./data/test.epub", func(n string, data []byte) bool {
		i++
		if data == nil {
			t.Fatal("reader failed")
		}
		return true
	})

	if i != 182 {
		t.Fatal("Invalid chapter numbers")
	}

	buf := bytes.NewBuffer(nil)

	ToTxt("./data/test.epub", buf)

	if buf.Len() == 0 {
		t.Fatal("ToTxt failed")
	}
}
