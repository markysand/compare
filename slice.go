/*Package compare has a function for comparing unsorted slices*/
package compare

import (
	"errors"
	"fmt"
	"reflect"
)

/*SliceUnsorted will check whether two slices are equal when
their sort order should not matter. Useful for testing, comparing
results. Will return error if input is not slice*/
func SliceUnsorted(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a).Kind() != reflect.Slice {
		return false, errors.New(fmt.Sprint("Arg <a> is not a slice", a))
	}
	if reflect.TypeOf(b).Kind() != reflect.Slice {
		return false, errors.New(fmt.Sprint("Arg <b> is not a slice", b))
	}
	aV, bV := reflect.ValueOf(a), reflect.ValueOf(b)
	aLen, bLen := aV.Len(), bV.Len()
	if aLen != bLen {
		return false, nil
	}
	usedB := make([]bool, aLen)
	var found bool
	for i := 0; i < aV.Len(); i++ {
		found = false
		for j := 0; j < bV.Len(); j++ {
			if usedB[j] {
				continue
			}
			if reflect.DeepEqual(aV.Index(i).Interface(), bV.Index(j).Interface()) {
				found = true
				usedB[j] = true
				break
			}
		}
		if found == false {
			return false, nil
		}
	}
	return true, nil
}
