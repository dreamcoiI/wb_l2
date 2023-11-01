package main

import (
	"reflect"
	"testing"
)

func Test_findAnagramSets(t *testing.T) {
	type args struct {
		words []string
	}
	var tests []struct {
		name string
		args args
		want map[string][]string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagramSets(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagramSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortString(t *testing.T) {
	type args struct {
		s string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
