package struct2interface

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var (
	src = []byte(`// Code generated by struct2interface; DO NOT EDIT.

package testdata

// MethodInterface ...
// Method describes the code and documentation
// tied into a method
type MethodInterface interface {
	// Lines return a []string consisting of
	// the documentation and code appended
	// in chronological order
	Lines() []string
}

// Method1Interface ...
// Method1 describes the code and documentation
// tied into a method
type Method1Interface interface {
	// Lines1 return a []string consisting of
	// the documentation and code appended
	// in chronological order
	Lines1() []string
}
`)
)

func TestMaker(t *testing.T) {
	files := []string{"./testdata/testdata.go"}
	output, err := Make(files)
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != string(src) {
		fmt.Println(len(output), len(src))
		fmt.Println("-------------------------")
		fmt.Println(strings.ReplaceAll(string(output), "\n", " "))
		fmt.Println("-------------------------")
		fmt.Println(strings.ReplaceAll(string(src), "\n", " "))
		fmt.Println("-------------------------")
		t.Fail()
	}
}

func TestDir(t *testing.T) {
	err := MakeDir("./testdata")
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile("./testdata/interface_testdata.go")
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != string(src) {
		t.Fail()
	}
}
