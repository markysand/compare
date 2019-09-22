package compare_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"github.com/markysand/compare"
)

func TestSliceUnsorted(t *testing.T) {
	const util = "Fails At: %v, Got: %v, Want: %v"
	tt := []struct {
		name string
		a    interface{}
		b    interface{}
		ok   bool
		err  bool
	}{
		{
			"String slices",
			[]string{
				"alpha", "beta",
			},
			[]string{
				"beta", "alpha",
			},
			true,
			false,
		},
		{
			"A not slice",
			"alpha beta",
			[]string{"gamma", "delta"},
			false,
			true,
		},
		{
			"B not slice",
			[]string{"gamma", "delta"},
			"alpha beta",
			false,
			true,
		},
		{
			"Different type",
			[]struct {
				foo string
				bar int
			}{
				{
					"hej",
					22,
				},
			},
			[]struct {
				yoda string
				bar  int
			}{
				{
					"hej",
					22,
				},
			},
			false,
			false,
		},
		{
			"Same type different value",
			[]struct {
				foo string
				bar int
			}{
				{
					"hej",
					22,
				},
			},
			[]struct {
				foo string
				bar int
			}{
				{
					"hej",
					21,
				},
			},
			false,
			false,
		},
		{
			"Same type same value",
			[]struct {
				foo string
				bar int
			}{
				{
					"hej",
					22,
				},
			},
			[]struct {
				foo string
				bar int
			}{
				{
					"hej",
					22,
				},
			},
			true,
			false,
		}, {
			"Same length and members but with different doubles will return false",
			[]string{
				"foo", "bar", "jedi", "yoda", "jedi",
			},
			[]string{
				"foo", "bar", "yoda", "bar", "jedi",
			},
			false,
			false,
		}, {
			"Same length and members with same doubles will return true",
			[]string{
				"foo", "bar", "jedi", "yoda", "jedi",
			},
			[]string{
				"foo", "jedi", "yoda", "bar", "jedi",
			},
			true,
			false,
		}, {
			"Different length",
			[]string{
				"foo", "bar", "jedi", "yoda", "jedi",
			},
			[]string{
				"foo", "bar", "yoda", "jedi",
			},
			false,
			false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := compare.SliceUnsorted(tc.a, tc.b)
			if ok != tc.ok {
				t.Errorf(util, "RESULT", ok, tc.ok)
			}
			if (err != nil) != tc.err {
				t.Errorf(util, "ERROR", err, tc.err)
			}
		})
	}
}

// With SliceUnsorted there is no need to construct
// a sort function for the data type in question
func ExampleSliceUnsorted() {
	sliceA := []struct {
		name    string
		surname string
	}{
		{"Donald", "Duck"},
		{"Mickey", "Mouse"},
	}
	sliceB := []struct {
		name    string
		surname string
	}{
		{"Mickey", "Mouse"},
		{"Donald", "Duck"},
	}
	result1, _ := compare.SliceUnsorted(sliceA, sliceB)
	result2 := reflect.DeepEqual(sliceA, sliceB)
	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// true
	// false
}

type test struct {
	foo string
	bar int
}

func getSliceUtil(n int) []test {
	res := make([]test, n)
	for i, t := range res {
		t.bar = i
		t.foo = strconv.Itoa(i)
	}
	rand.Shuffle(n, func(i int, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}

func BenchmarkSliceUnsorted(b *testing.B) {
	const n = 1000
	arg1, arg2 := getSliceUtil(1000), getSliceUtil(1000)
	for i := 0; i < b.N; i++ {
		_, err := compare.SliceUnsorted(arg1, arg2)
		if err != nil {
			b.Error(err.Error())
		}
	}
}
