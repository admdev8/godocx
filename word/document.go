package word

import (
	"encoding/xml"
	"os"
	"path"
)

type Document struct {
	XMLName  xml.Name `xml:"w:document"`
	XmlnsO   string   `xml:"xmlns:o,attr,omitempty"`
	XmlnsR   string   `xml:"xmlns:r,attr,omitempty"`
	XmlnsM   string   `xml:"xmlns:m,attr,omitempty"`
	XmlnsV   string   `xml:"xmlns:v,attr,omitempty"`
	XmlnsW   string   `xml:"xmlns:w,attr,omitempty"`
	XmlnsVe  string   `xml:"xmlns:ve,attr,omitempty"`
	XmlnsWp  string   `xml:"xmlns:wp,attr,omitempty"`
	XmlnsW10 string   `xml:"xmlns:w10,attr,omitempty"`
	XmlnsWne string   `xml:"xmlns:wne,attr,omitempty"`
	Body     Body
}

func NewDocument() *Document {
	d := &Document{}

	return d
}

func (d *Document) AddParagraph() *Paragraph {
	paragh := NewParagraph()
	d.Body.Content = append(d.Body.Content, paragh)
	return paragh
}

func (d *Document) AddTable() *Table {
	tbl := NewTable()
	d.Body.Content = append(d.Body.Content, tbl)
	return tbl
}

func (d *Document) Save(dirpath string) error {
	fpath := path.Join(dirpath, "word")
	os.Mkdir(fpath, os.ModePerm)

	output, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(fpath, "document.xml"))
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(xml.Header)
	f.Write(output)

	return nil
}
