// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"gooxml"
)

type Numbering struct {
	CT_Numbering
}

func NewNumbering() *Numbering {
	ret := &Numbering{}
	ret.CT_Numbering = *NewCT_Numbering()
	return ret
}

func (m *Numbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:numbering"
	return m.CT_Numbering.MarshalXML(e, start)
}

func (m *Numbering) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Numbering = *NewCT_Numbering()
lNumbering:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numPicBullet"}:
				tmp := NewCT_NumPicBullet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.NumPicBullet = append(m.NumPicBullet, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "abstractNum"}:
				tmp := NewCT_AbstractNum()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AbstractNum = append(m.AbstractNum, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "num"}:
				tmp := NewCT_Num()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Num = append(m.Num, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numIdMacAtCleanup"}:
				m.NumIdMacAtCleanup = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumIdMacAtCleanup, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on Numbering %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lNumbering
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Numbering and its children
func (m *Numbering) Validate() error {
	return m.ValidateWithPath("Numbering")
}

// ValidateWithPath validates the Numbering and its children, prefixing error messages with path
func (m *Numbering) ValidateWithPath(path string) error {
	if err := m.CT_Numbering.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
