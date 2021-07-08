package main

import (
	"reflect"
	"testing"
)

var testcases = map[string]struct {
	in  [][]int
	out [][]int
}{
	"2 pixeles eliminados": {
		in: [][]int{
			{1, 0, 1, 0, 0},
			{0, 1, 0, 1, 1},
			{0, 0, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 0, 1},
		},
		out: [][]int{
			{1, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		},
	},
	"0 pixeles eliminados": {
		in: [][]int{
			{1, 0, 1, 0, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 0, 1},
		},
		out: [][]int{
			{1, 0, 1, 0, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 0, 1},
		},
	},
}

func TestRemove(t *testing.T) {
	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			x := removeBlackPixels(tt.in)
			if !reflect.DeepEqual(tt.out, x) {
				t.Errorf("got %v, want %v", x, tt.out)
			}
		})
	}
}
