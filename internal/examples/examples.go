package examples

import (
	"log"

	somecolor "github.com/nikolaydubina/go-enum-example/internal/color"
)

// MutateExportedVar shows that Go compiler does not error on mutating exporting structs
func MutateExportedVar() {
	somecolor.Blue = somecolor.Red
	somecolor.Blue = somecolor.Undefined
	log.Println(somecolor.Blue)
}

// ImplicitAssignmentToNonExportedStruct shows that Go compiler errors on implicit assigment of non exported fields
func ImplicitAssignmentToNonExportedStruct() {
	//somecolor.Blue = somecolor.Color{5}
}
