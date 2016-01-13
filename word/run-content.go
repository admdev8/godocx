package word

import (
	"encoding/xml"
)

type RunContent struct {
	XMLName xml.Name       `xml:"w:r"`
	RsidRPr string         `xml:"w:rsidRPr,attr,omitempty"`
	RPr     *RunProperties `xml:"w:rPr"`
	Content []interface{}
	T       string `xml:"w:t,omitempty"`
}

func NewRunContent() *RunContent {
	return &RunContent{Content: make([]interface{}, 0)}
}

func (r *RunContent) AddRunProperties() *RunProperties {
	obj := NewRunProperties()
	r.Content = append(r.Content, obj)
	return obj
}

func (r *RunContent) Text(val string) {
	r.T = val
}