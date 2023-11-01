package main

import "testing"

func TestUnpackString(t *testing.T) {
	type args struct {
		str string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnpackString(tt.args.str); got != tt.want {
				t.Errorf("UnpackString() = %v, want %v", got, tt.want)
			}
		})
	}
}
