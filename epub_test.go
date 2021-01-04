package epub

import (
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
}
