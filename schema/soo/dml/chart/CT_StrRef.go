// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"

	"gooxml"
)

type CT_StrRef struct {
	F        string
	StrCache *CT_StrData
	ExtLst   *CT_ExtensionList
}

func NewCT_StrRef() *CT_StrRef {
	ret := &CT_StrRef{}
	return ret
}

func (m *CT_StrRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sef := xml.StartElement{Name: xml.Name{Local: "c:f"}}
	gooxml.AddPreserveSpaceAttr(&sef, m.F)
	e.EncodeElement(m.F, sef)
	if m.StrCache != nil {
		sestrCache := xml.StartElement{Name: xml.Name{Local: "c:strCache"}}
		e.EncodeElement(m.StrCache, sestrCache)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StrRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_StrRef:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "f"}:
				if err := d.DecodeElement(&m.F, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strCache"}:
				m.StrCache = NewCT_StrData()
				if err := d.DecodeElement(m.StrCache, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_StrRef %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StrRef
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StrRef and its children
func (m *CT_StrRef) Validate() error {
	return m.ValidateWithPath("CT_StrRef")
}

// ValidateWithPath validates the CT_StrRef and its children, prefixing error messages with path
func (m *CT_StrRef) ValidateWithPath(path string) error {
	if m.StrCache != nil {
		if err := m.StrCache.ValidateWithPath(path + "/StrCache"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
