package epub

import (
	"fmt"
	"testing"
)

func TestEpub(t *testing.T) {
	bk, err := open(t, "./data/test.epub")
	if err != nil {
		t.Fatal(err)
	}
	defer bk.Close()
}

func open(t *testing.T, f string) (*Book, error) {
	bk, err := Open(f)
	if err != nil {
		return nil, err
	}
	defer bk.Close()

	if len(bk.Files()) != 211 {
		t.Fatal("invalid files counter")
	}

	for _, pt := range bk.Ncx.Points {
		for _, np := range pt.Points {
			fmt.Println(np.Text + " . " + np.Content.Src)
		}
	}

	return bk, nil
}
