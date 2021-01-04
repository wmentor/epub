package epub

type Opf struct {
	Metadata Metadata   `xml:"metadata"`
	Manifest []Manifest `xml:"manifest>item"`
	Spine    Spine      `xml:"spine"`
}

type Metadata struct {
	Title       []string     `xml:"title"`
	Language    []string     `xml:"language"`
	Identifier  []Identifier `xml:"identifier"`
	Creator     []Author     `xml:"creator"`
	Subject     []string     `xml:"subject"`
	Description []string     `xml:"description"`
	Publisher   []string     `xml:"publisher"`
	Contributor []Author     `xml:"contributor"`
	Date        []Date       `xml:"date"`
	Type        []string     `xml:"type"`
	Format      []string     `xml:"format"`
	Source      []string     `xml:"source"`
	Relation    []string     `xml:"relation"`
	Coverage    []string     `xml:"coverage"`
	Rights      []string     `xml:"rights"`
	Meta        []Metafield  `xml:"meta"`
}

type Identifier struct {
	Data   string `xml:",chardata"`
	ID     string `xml:"id,attr"`
	Scheme string `xml:"scheme,attr"`
}

type Author struct {
	Data   string `xml:",chardata"`
	FileAs string `xml:"file-as,attr"`
	Role   string `xml:"role,attr"`
}

type Date struct {
	Data  string `xml:",chardata"`
	Event string `xml:"event,attr"`
}

type Metafield struct {
	Name    string `xml:"name,attr"`
	Content string `xml:"content,attr"`
}

type Manifest struct {
	ID           string `xml:"id,attr"`
	Href         string `xml:"href,attr"`
	MediaType    string `xml:"media-type,attr"`
	Fallback     string `xml:"media-fallback,attr"`
	Properties   string `xml:"properties,attr"`
	MediaOverlay string `xml:"media-overlay,attr"`
}

type Spine struct {
	ID              string      `xml:"id,attr"`
	Toc             string      `xml:"toc,attr"`
	PageProgression string      `xml:"page-progression-direction,attr"`
	Items           []SpineItem `xml:"itemref"`
}

type SpineItem struct {
	IDref      string `xml:"idref,attr"`
	Linear     string `xml:"linear,attr"`
	ID         string `xml:"id,attr"`
	Properties string `xml:"properties,attr"`
}
