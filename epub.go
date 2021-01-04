package epub

import (
	"archive/zip"
	"io/ioutil"
)

func Open(fn string) (*Book, error) {
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}

	bk := Book{fd: fd}

	mt, err := bk.readBytes("mimetype")
	if err == nil {
		bk.Mimetype = string(mt)
		err = bk.readXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)
	}

	for _, mf := range bk.Opf.Manifest {
		if mf.ID == bk.Opf.Spine.Toc {
			err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
			break
		}
	}

	if err != nil {
		fd.Close()
		return nil, err
	}

	return &bk, nil
}

func Reader(filename string, onChapter func(chapter string, data []byte) bool) error {
	bk, err := Open(filename)
	if err != nil {
		return err
	}
	defer bk.Close()

	if onChapter == nil {
		return nil
	}

	readerF := func(filename string) ([]byte, error) {
		fd, err := bk.Open(filename)
		if err != nil {
			return nil, err
		}
		defer fd.Close()

		return ioutil.ReadAll(fd)
	}

	for _, pt := range bk.Ncx.Points {
		for _, np := range pt.Points {

			name := np.Text

			data, err := readerF(np.Content.Src)
			if err != nil {
				return err
			}

			if !onChapter(name, data) {
				return nil
			}

		}
	}

	return nil
}
