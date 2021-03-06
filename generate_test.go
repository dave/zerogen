package main

import (
	"bytes"
	"testing"
)

func TestSimple(t *testing.T) {

	pkg := "github.com/mrsinham/zerogen/test/simple"
	structure := "simple"
	write := false
	b := &bytes.Buffer{}
	if err := parsePackage(&pkg, &structure, &write, b); err != nil {
		t.Fatal(err)
	}

	result := `// Code generated by github.com/mrsinham/zerogen. DO NOT EDIT.
package simple

func (s *simple) Reset() {
	s.field1 = 0
	s.field2 = ""
	s.field3 = nil
	s.field4 = [3]int{}
	s.field5 = nil
}
`
	if b.String() != result {
		t.Fatalf("result is not the same as expected : got \n%q \n expected  \n%q	", b.String(), result)
	}

}

func TestComposition(t *testing.T) {

	pkg := "github.com/mrsinham/zerogen/test/composition"
	structure := "composited"
	write := false
	b := &bytes.Buffer{}
	if err := parsePackage(&pkg, &structure, &write, b); err != nil {
		t.Fatal(err)
	}

	result := `// Code generated by github.com/mrsinham/zerogen. DO NOT EDIT.
package composition

func (c *composited) Reset() {
	c.subfield1 = 0
	c.subfield2 = ""
	c.fieldsimple = ""
}
`
	if b.String() != result {
		t.Fatalf("result is not the same as expected : got \n%q \n expected  \n%q", b.String(), result)
	}

}

func TestCompositionExternal(t *testing.T) {

	pkg := "github.com/mrsinham/zerogen/test/compositionexternal"
	structure := "base"
	write := false
	b := &bytes.Buffer{}
	if err := parsePackage(&pkg, &structure, &write, b); err != nil {
		t.Fatal(err)
	}

	result := `// Code generated by github.com/mrsinham/zerogen. DO NOT EDIT.
package compositionexternal

func (b *base) Reset() {
	b.field1 = ""
	b.Field1 = ""
}
`
	if b.String() != result {
		t.Fatalf("result is not the same as expected : got \n%q \n expected  \n%q", b.String(), result)
	}

}

func TestCompositionExternalNotExported(t *testing.T) {

	pkg := "github.com/mrsinham/zerogen/test/compositionexternal"
	structure := "compositionNotExported"
	write := false
	b := &bytes.Buffer{}
	if err := parsePackage(&pkg, &structure, &write, b); err != nil {
		t.Fatal(err)
	}

	result := `// Code generated by github.com/mrsinham/zerogen. DO NOT EDIT.
package compositionexternal

import sub "github.com/mrsinham/zerogen/test/compositionexternal/sub"

func (c *compositionNotExported) Reset() {
	c.field1 = ""
	c.Sub2 = sub.Sub2{}
}
`
	if b.String() != result {
		t.Fatalf("result is not the same as expected : got \n%q \n expected  \n%q", b.String(), result)
	}

}

func TestNoNil(t *testing.T) {

	pkg := "github.com/mrsinham/zerogen/test/nonil"
	structure := "nonil"
	write := false
	b := &bytes.Buffer{}
	if err := parsePackage(&pkg, &structure, &write, b); err != nil {
		t.Fatal(err)
	}

	result := `// Code generated by github.com/mrsinham/zerogen. DO NOT EDIT.
package nonil

import http "net/http"

func (n *nonil) Reset() {
	n.field1 = [4]int{}
	n.field2 = &http.Request{}
}
`
	if b.String() != result {
		t.Fatalf("result is not the same as expected : got \n%q \n expected  \n%q", b.String(), result)
	}

}

func TestNoNilExternal(t *testing.T) {

	pkg := "github.com/mrsinham/zerogen/test/nonilexternal"
	structure := "nonilexternal"
	write := false
	b := &bytes.Buffer{}
	if err := parsePackage(&pkg, &structure, &write, b); err != nil {
		t.Fatal(err)
	}

	result := `// Code generated by github.com/mrsinham/zerogen. DO NOT EDIT.
package nonilexternal

func (n *nonilexternal) Reset() {
	n.field1 = 0
	n.Field2 = 0
	n.Field3 = 0
}
`
	if b.String() != result {
		t.Fatalf("result is not the same as expected : got \n%q \n expected  \n%q", b.String(), result)
	}

}

func TestWithCompositionWithReset(t *testing.T) {
	pkg := "github.com/mrsinham/zerogen/test/nonilexternalreset"
	structure := "nonilExternalWithReset"
	write := false
	b := &bytes.Buffer{}
	if err := parsePackage(&pkg, &structure, &write, b); err != nil {
		t.Fatal(err)
	}

	result := `// Code generated by github.com/mrsinham/zerogen. DO NOT EDIT.
package nonilexternalreset

func (n *nonilExternalWithReset) Reset() {
	n.Field1 = 0
	n.ExternalReset.Reset()
	if n.interfaceReset != nil {
		n.interfaceReset.Reset()
	} else {
		n.interfaceReset = nil
	}
}
`
	if b.String() != result {
		t.Fatalf("result is not the same as expected : got \n%q \n expected  \n%q", b.String(), result)
	}
}
