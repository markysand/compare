/*Package compare has a function for comparing unsorted slices*/
package compare

import (
	"errors"
	"fmt"
	"reflect"
)

/*SliceUnsorted will check whether two slices have the
same length and the same set of members. This can be useful
for writing tests/validation. An error will be returned if
either argument is ont a slice*/
func SliceUnsorted(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a).Kind() != reflect.Slice {
		return false, errors.New(fmt.Sprint("Arg <a> is not a slice", a))
	}
	if reflect.TypeOf(b).Kind() != reflect.Slice {
		return false, errors.New(fmt.Sprint("Arg <b> is not a slice", b))
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
