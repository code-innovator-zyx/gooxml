// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"gooxml"
)

type EG_TextUnderlineLine struct {
	ULnTx *CT_TextUnderlineLineFollowText
	ULn   *CT_LineProperties
}

func NewEG_TextUnderlineLine() *EG_TextUnderlineLine {
	ret := &EG_TextUnderlineLine{}
	return ret
}

func (m *EG_TextUnderlineLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ULnTx != nil {
		seuLnTx := xml.StartElement{Name: xml.Name{Local: "a:uLnTx"}}
		e.EncodeElement(m.ULnTx, seuLnTx)
	}
	if m.ULn != nil {
		seuLn := xml.StartElement{Name: xml.Name{Local: "a:uLn"}}
		e.EncodeElement(m.ULn, seuLn)
	}
	return nil
}

func (m *EG_TextUnderlineLine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextUnderlineLine:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "uLnTx"}:
				m.ULnTx = NewCT_TextUnderlineLineFollowText()
				if err := d.DecodeElement(m.ULnTx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "uLn"}:
				m.ULn = NewCT_LineProperties()
				if err := d.DecodeElement(m.ULn, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_TextUnderlineLine %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextUnderlineLine
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TextUnderlineLine and its children
func (m *EG_TextUnderlineLine) Validate() error {
	return m.ValidateWithPath("EG_TextUnderlineLine")
}

// ValidateWithPath validates the EG_TextUnderlineLine and its children, prefixing error messages with path
func (m *EG_TextUnderlineLine) ValidateWithPath(path string) error {
	if m.ULnTx != nil {
		if err := m.ULnTx.ValidateWithPath(path + "/ULnTx"); err != nil {
			return err
		}
	}
	if m.ULn != nil {
		if err := m.ULn.ValidateWithPath(path + "/ULn"); err != nil {
			return err
		}
	}
	return nil
}
