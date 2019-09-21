package compare

import (
	"errors"
	"fmt"
	"reflect"
)

// SliceUnsorted will check whether two slices have the
// same length and the same set of members. This can be useful
// for writing tests, comparing search results of unwieldy objects that
// are difficult to sort.
func SliceUnsorted(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a).Kind() != reflect.Slice {
		return false, errors.New(fmt.Sprint("Arg a not slice", a))
	}
	if reflect.TypeOf(b).Kind() != reflect.Slice {
		return false, errors.New(fmt.Sprint("Arg b not slice", b))
	}
	aV, bV := reflect.ValueOf(a), reflect.ValueOf(b)
	if aV.Len() != bV.Len() {
		return false, nil
	}
	for i := 0; i < aV.Len(); i++ {
		var found bool
		for j := 0; j < bV.Len(); j++ {
			if reflect.DeepEqual(aV.Index(i).Interface(), bV.Index(j).Interface()) {
				found = true
				break
			}
		}
		if found == false {
			return false, nil
		}
	}
	return true, nil
}
