// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"gooxml/schema/soo/sml"
)

func TestCT_BookViewConstructor(t *testing.T) {
	v := sml.NewCT_BookView()
	if v == nil {
		t.Errorf("sml.NewCT_BookView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_BookView should validate: %s", err)
	}
}

func TestCT_BookViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_BookView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_BookView()
	xml.Unmarshal(buf, v2)
}
