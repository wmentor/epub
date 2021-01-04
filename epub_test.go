package epub

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestReader(t *testing.T) {
	bk, err := Open("./data/test.epub")
	if err != nil {
		t.Fatal(err)
	}
	defer bk.Close()

	if len(bk.Files()) != 211 {
		t.Fatal("invalid files counter")
	}

	for _, pt := range bk.Ncx.Points {
		for _, np := range pt.Points {
			fmt.Println(np.Text + " . " + np.Content.Src)

			input, err := bk.Open(np.Content.Src)
			if err != nil {
				t.Fatal(err)
			}
			defer input.Close()

			data, err := ioutil.ReadAll(input)
			if err != nil {
				t.Fatal(err)
			}

			fmt.Println(string(data))

		}
	}

}
