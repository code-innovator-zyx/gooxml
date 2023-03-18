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
	"fmt"

	"gooxml"
	"gooxml/schema/soo/dml"
)

type CT_BarSer struct {
	Idx              *CT_UnsignedInt
	Order            *CT_UnsignedInt
	Tx               *CT_SerTx
	SpPr             *dml.CT_ShapeProperties
	InvertIfNegative *CT_Boolean
	PictureOptions   *CT_PictureOptions
	DPt              []*CT_DPt
	DLbls            *CT_DLbls
	Trendline        []*CT_Trendline
	ErrBars          *CT_ErrBars
	Cat              *CT_AxDataSource
	Val              *CT_NumDataSource
	Shape            *CT_Shape
	ExtLst           *CT_ExtensionList
}

func NewCT_BarSer() *CT_BarSer {
	ret := &CT_BarSer{}
	ret.Idx = NewCT_UnsignedInt()
	ret.Order = NewCT_UnsignedInt()
	return ret
}

func (m *CT_BarSer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "c:idx"}}
	e.EncodeElement(m.Idx, seidx)
	seorder := xml.StartElement{Name: xml.Name{Local: "c:order"}}
	e.EncodeElement(m.Order, seorder)
	if m.Tx != nil {
		setx := xml.StartElement{Name: xml.Name{Local: "c:tx"}}
		e.EncodeElement(m.Tx, setx)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.InvertIfNegative != nil {
		seinvertIfNegative := xml.StartElement{Name: xml.Name{Local: "c:invertIfNegative"}}
		e.EncodeElement(m.InvertIfNegative, seinvertIfNegative)
	}
	if m.PictureOptions != nil {
		sepictureOptions := xml.StartElement{Name: xml.Name{Local: "c:pictureOptions"}}
		e.EncodeElement(m.PictureOptions, sepictureOptions)
	}
	if m.DPt != nil {
		sedPt := xml.StartElement{Name: xml.Name{Local: "c:dPt"}}
		for _, c := range m.DPt {
			e.EncodeElement(c, sedPt)
		}
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "c:dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.Trendline != nil {
		setrendline := xml.StartElement{Name: xml.Name{Local: "c:trendline"}}
		for _, c := range m.Trendline {
			e.EncodeElement(c, setrendline)
		}
	}
	if m.ErrBars != nil {
		seerrBars := xml.StartElement{Name: xml.Name{Local: "c:errBars"}}
		e.EncodeElement(m.ErrBars, seerrBars)
	}
	if m.Cat != nil {
		secat := xml.StartElement{Name: xml.Name{Local: "c:cat"}}
		e.EncodeElement(m.Cat, secat)
	}
	if m.Val != nil {
		seval := xml.StartElement{Name: xml.Name{Local: "c:val"}}
		e.EncodeElement(m.Val, seval)
	}
	if m.Shape != nil {
		seshape := xml.StartElement{Name: xml.Name{Local: "c:shape"}}
		e.EncodeElement(m.Shape, seshape)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BarSer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
	m.Order = NewCT_UnsignedInt()
lCT_BarSer:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "idx"}:
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "order"}:
				if err := d.DecodeElement(m.Order, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "tx"}:
				m.Tx = NewCT_SerTx()
				if err := d.DecodeElement(m.Tx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "invertIfNegative"}:
				m.InvertIfNegative = NewCT_Boolean()
				if err := d.DecodeElement(m.InvertIfNegative, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pictureOptions"}:
				m.PictureOptions = NewCT_PictureOptions()
				if err := d.DecodeElement(m.PictureOptions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dPt"}:
				tmp := NewCT_DPt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DPt = append(m.DPt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbls"}:
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "trendline"}:
				tmp := NewCT_Trendline()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Trendline = append(m.Trendline, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "errBars"}:
				m.ErrBars = NewCT_ErrBars()
				if err := d.DecodeElement(m.ErrBars, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "cat"}:
				m.Cat = NewCT_AxDataSource()
				if err := d.DecodeElement(m.Cat, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "val"}:
				m.Val = NewCT_NumDataSource()
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "shape"}:
				m.Shape = NewCT_Shape()
				if err := d.DecodeElement(m.Shape, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_BarSer %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BarSer
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BarSer and its children
func (m *CT_BarSer) Validate() error {
	return m.ValidateWithPath("CT_BarSer")
}

// ValidateWithPath validates the CT_BarSer and its children, prefixing error messages with path
func (m *CT_BarSer) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if err := m.Order.ValidateWithPath(path + "/Order"); err != nil {
		return err
	}
	if m.Tx != nil {
		if err := m.Tx.ValidateWithPath(path + "/Tx"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.InvertIfNegative != nil {
		if err := m.InvertIfNegative.ValidateWithPath(path + "/InvertIfNegative"); err != nil {
			return err
		}
	}
	if m.PictureOptions != nil {
		if err := m.PictureOptions.ValidateWithPath(path + "/PictureOptions"); err != nil {
			return err
		}
	}
	for i, v := range m.DPt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DPt[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	for i, v := range m.Trendline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Trendline[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ErrBars != nil {
		if err := m.ErrBars.ValidateWithPath(path + "/ErrBars"); err != nil {
			return err
		}
	}
	if m.Cat != nil {
		if err := m.Cat.ValidateWithPath(path + "/Cat"); err != nil {
			return err
		}
	}
	if m.Val != nil {
		if err := m.Val.ValidateWithPath(path + "/Val"); err != nil {
			return err
		}
	}
	if m.Shape != nil {
		if err := m.Shape.ValidateWithPath(path + "/Shape"); err != nil {
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
