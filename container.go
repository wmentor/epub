package epub

type Container struct {
	Rootfile Rootfile `xml:"rootfiles>rootfile" json:"rootfile"`
}

type Rootfile struct {
	Path string `xml:"full-path,attr"`
	Type string `xml:"media-type,attr"`
}
