package struct2interface

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDirCompared = `// Code generated by struct2interface; DO NOT EDIT.

package case_single_file

// MethodInterface ...
//
//	Method describes the code and documentation
//	tied into a method
type MethodInterface interface {
	// Lines return a []string consisting of
	// the documentation and code appended
	// in chronological order
	Lines() []string
}

// Method1Interface ...
//
//	Method1 describes the code and documentation
//	tied into a method
type Method1Interface interface {
	// Lines return a []string consisting of
	// the documentation and code appended
	// in chronological order
	Lines() []string
}
`
	testPackageCompared = `// Code generated by struct2interface; DO NOT EDIT.

package testdata

// PackageMethodInterface ...
type PackageMethodInterface interface {
	// the //-style comment test
	Method1() string
	Method2() string
}

// PackageMethod2Interface ...
type PackageMethod2Interface interface {
	/*
	   the /*-style comment test
	*/
	Method1() string
}
`
)

func TestDir(t *testing.T) {
	err := MakeDir("./testdata/case_single_file")
	if err != nil {
		t.Fatal(err)
	}

	output, err := os.ReadFile("./testdata/case_single_file/interface_case_single_file.go")
	if err != nil {
		t.Fatal(err)
	}
	if string(output) != testDirCompared {
		t.Fail()
	}
}

func TestPackage(t *testing.T) {
	err := MakeDir("./testdata/case_package")
	if err != nil {
		t.Fatal(err)
	}

	output, err := os.ReadFile("./testdata/case_package/interface_testdata.go")
	if err != nil {
		t.Fatal(err)
	}
	if string(output) != testPackageCompared {
		t.Fail()
	}
}

func TestNil(t *testing.T) {
	t.Run("空路径", func(t *testing.T) {
		err := MakeDir("./notfind")
		assert.EqualError(t, err, "lstat ./notfind: no such file or directory")
	})
}

func Benchmark_TestDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeDir("./testdata")
	}
}
