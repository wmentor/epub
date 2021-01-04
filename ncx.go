package epub

type Ncx struct {
	Points []NavPoint `xml:"navMap>navPoint"`
}

type NavPoint struct {
	Text    string     `xml:"navLabel>text"`
	Content Content    `xml:"content"`
	Points  []NavPoint `xml:"navPoint"`
}

type Content struct {
	Src string `xml:"src,attr"`
}
