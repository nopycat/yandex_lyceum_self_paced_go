package main

import (
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name  string
		value []byte
		want1 int
		want2 error
	}{
		{
			name:  "negative",
			value: []byte{65},
			want1: 1,
			want2: nil,
		},
		{
			name:  "xyu",
			value: []byte{225},
			want1: 0,
			want2: ErrInvalidUTF8,
		},
		// {
		// 	name:  "short",
		// 	value: "i",
		// 	want:  "",
		// },
		// {
		// 	name:  "long",
		// 	value: "o",
		// 	want:  "",
		// },
		// {
		// 	name:  "very long",
		// 	value: "u",
		// 	want:  "",
		// },
		// {
		// 	name:  "xyu",
		// 	value: "f",
		// 	want:  "f",
		// },
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			value, got := GetUTFLength(tc.value)
			if value != tc.want1 && got != nil {
				t.Fail()
			}
		})
	}
}
