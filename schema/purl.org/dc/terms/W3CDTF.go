// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import (
	"encoding/xml"
	"fmt"
)

type W3CDTF struct {
}

func NewW3CDTF() *W3CDTF {
	ret := &W3CDTF{}
	return ret
}

func (m *W3CDTF) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "W3CDTF"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *W3CDTF) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing W3CDTF: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the W3CDTF and its children
func (m *W3CDTF) Validate() error {
	return m.ValidateWithPath("W3CDTF")
}

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (m *W3CDTF) ValidateWithPath(path string) error {
	return nil
}
