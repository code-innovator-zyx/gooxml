// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"gooxml"
)

type CT_RevisionFormatting struct {
	// Sheet Id
	SheetIdAttr uint32
	// Row or Column Formatting Change
	XfDxfAttr *bool
	// Style
	SAttr *bool
	// Sequence Of References
	SqrefAttr ST_Sqref
	// Start index
	StartAttr *uint32
	// Length
	LengthAttr *uint32
	// Formatting
	Dxf    *CT_Dxf
	ExtLst *CT_ExtensionList
}

func NewCT_RevisionFormatting() *CT_RevisionFormatting {
	ret := &CT_RevisionFormatting{}
	return ret
}

func (m *CT_RevisionFormatting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	if m.XfDxfAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xfDxf"},
			Value: fmt.Sprintf("%d", b2i(*m.XfDxfAttr))})
	}
	if m.SAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
			Value: fmt.Sprintf("%d", b2i(*m.SAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
		Value: fmt.Sprintf("%v", m.SqrefAttr)})
	if m.StartAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "start"},
			Value: fmt.Sprintf("%v", *m.StartAttr)})
	}
	if m.LengthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "length"},
			Value: fmt.Sprintf("%v", *m.LengthAttr)})
	}
	e.EncodeToken(start)
	if m.Dxf != nil {
		sedxf := xml.StartElement{Name: xml.Name{Local: "ma:dxf"}}
		e.EncodeElement(m.Dxf, sedxf)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionFormatting) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "xfDxf" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.XfDxfAttr = &parsed
			continue
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SAttr = &parsed
			continue
		}
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = parsed
			continue
		}
		if attr.Name.Local == "start" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.StartAttr = &pt
			continue
		}
		if attr.Name.Local == "length" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LengthAttr = &pt
			continue
		}
	}
lCT_RevisionFormatting:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dxf"}:
				m.Dxf = NewCT_Dxf()
				if err := d.DecodeElement(m.Dxf, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_RevisionFormatting %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionFormatting
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RevisionFormatting and its children
func (m *CT_RevisionFormatting) Validate() error {
	return m.ValidateWithPath("CT_RevisionFormatting")
}

// ValidateWithPath validates the CT_RevisionFormatting and its children, prefixing error messages with path
func (m *CT_RevisionFormatting) ValidateWithPath(path string) error {
	if m.Dxf != nil {
		if err := m.Dxf.ValidateWithPath(path + "/Dxf"); err != nil {
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
