package compare_test

import (
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
			"Not slice",
			"alpha beta",
			[]string{
				"alpha", "beta",
			},
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
			"NOTE! Same length and members but with different doubles will return true",
			[]string{
				"foo", "bar", "jedi", "yoda", "jedi",
			},
			[]string{
				"foo", "bar", "yoda", "bar", "jedi",
			},
			true,
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
