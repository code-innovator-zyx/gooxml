// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"

	"gooxml"
)

type OMath struct {
	CT_OMath
}

func NewOMath() *OMath {
	ret := &OMath{}
	ret.CT_OMath = *NewCT_OMath()
	return ret
}

func (m *OMath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "m:oMath"
	return m.CT_OMath.MarshalXML(e, start)
}

func (m *OMath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_OMath = *NewCT_OMath()
lOMath:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "acc"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Acc = NewCT_Acc()
				if err := d.DecodeElement(tmpomathmathelements.Acc, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "bar"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Bar = NewCT_Bar()
				if err := d.DecodeElement(tmpomathmathelements.Bar, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "box"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Box = NewCT_Box()
				if err := d.DecodeElement(tmpomathmathelements.Box, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "borderBox"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.BorderBox = NewCT_BorderBox()
				if err := d.DecodeElement(tmpomathmathelements.BorderBox, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "d"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.D = NewCT_D()
				if err := d.DecodeElement(tmpomathmathelements.D, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "eqArr"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.EqArr = NewCT_EqArr()
				if err := d.DecodeElement(tmpomathmathelements.EqArr, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "f"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.F = NewCT_F()
				if err := d.DecodeElement(tmpomathmathelements.F, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "func"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Func = NewCT_Func()
				if err := d.DecodeElement(tmpomathmathelements.Func, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "groupChr"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.GroupChr = NewCT_GroupChr()
				if err := d.DecodeElement(tmpomathmathelements.GroupChr, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limLow"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.LimLow = NewCT_LimLow()
				if err := d.DecodeElement(tmpomathmathelements.LimLow, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limUpp"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.LimUpp = NewCT_LimUpp()
				if err := d.DecodeElement(tmpomathmathelements.LimUpp, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "m"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.M = NewCT_M()
				if err := d.DecodeElement(tmpomathmathelements.M, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "nary"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Nary = NewCT_Nary()
				if err := d.DecodeElement(tmpomathmathelements.Nary, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "phant"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Phant = NewCT_Phant()
				if err := d.DecodeElement(tmpomathmathelements.Phant, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "rad"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.Rad = NewCT_Rad()
				if err := d.DecodeElement(tmpomathmathelements.Rad, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sPre"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SPre = NewCT_SPre()
				if err := d.DecodeElement(tmpomathmathelements.SPre, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSub"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSub = NewCT_SSub()
				if err := d.DecodeElement(tmpomathmathelements.SSub, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSubSup"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSubSup = NewCT_SSubSup()
				if err := d.DecodeElement(tmpomathmathelements.SSubSup, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSup"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.SSup = NewCT_SSup()
				if err := d.DecodeElement(tmpomathmathelements.SSup, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "r"}:
				tmpomathmathelements := NewEG_OMathMathElements()
				tmpomathmathelements.R = NewCT_R()
				if err := d.DecodeElement(tmpomathmathelements.R, &el); err != nil {
					return err
				}
				m.EG_OMathMathElements = append(m.EG_OMathMathElements, tmpomathmathelements)
			default:
				gooxml.Log("skipping unsupported element on OMath %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOMath
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OMath and its children
func (m *OMath) Validate() error {
	return m.ValidateWithPath("OMath")
}

// ValidateWithPath validates the OMath and its children, prefixing error messages with path
func (m *OMath) ValidateWithPath(path string) error {
	if err := m.CT_OMath.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
