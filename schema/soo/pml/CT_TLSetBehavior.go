// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"

	"gooxml"
)

type CT_TLSetBehavior struct {
	// Common Behavior
	CBhvr *CT_TLCommonBehaviorData
	// To
	To *CT_TLAnimVariant
}

func NewCT_TLSetBehavior() *CT_TLSetBehavior {
	ret := &CT_TLSetBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}

func (m *CT_TLSetBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	if m.To != nil {
		seto := xml.StartElement{Name: xml.Name{Local: "p:to"}}
		e.EncodeElement(m.To, seto)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLSetBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
lCT_TLSetBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cBhvr"}:
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "to"}:
				m.To = NewCT_TLAnimVariant()
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TLSetBehavior %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLSetBehavior
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLSetBehavior and its children
func (m *CT_TLSetBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLSetBehavior")
}

// ValidateWithPath validates the CT_TLSetBehavior and its children, prefixing error messages with path
func (m *CT_TLSetBehavior) ValidateWithPath(path string) error {
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	if m.To != nil {
		if err := m.To.ValidateWithPath(path + "/To"); err != nil {
			return err
		}
	}
	return nil
}
