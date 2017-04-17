package rst

import (
	"io"

	"github.com/mrsinham/rst/other"
)

type blu struct {
	tutu int
}

type customFiles string

type test struct {
	diane                           [4]customFiles
	moineau                         [5]other.MyOtherType
	blublu                          []int
	blu                             *blu
	fieldWithCustomTypeFromOtherPkg other.MyOtherType
	fieldWithCustomType             customFiles
	teststr                         string
	testMap                         map[string]int
	// testint is here
	testint int
	io.Writer
}
