// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/soo/ofc/docPropsVTypes"
)

func TestVstreamConstructor(t *testing.T) {
	v := docPropsVTypes.NewVstream()
	if v == nil {
		t.Errorf("docPropsVTypes.NewVstream must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Vstream should validate: %s", err)
	}
}

func TestVstreamMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewVstream()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewVstream()
	xml.Unmarshal(buf, v2)
}
