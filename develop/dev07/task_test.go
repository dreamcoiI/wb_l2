package main

import (
	"reflect"
	"testing"
)

func Test_or(t *testing.T) {
	type args struct {
		ch []<-chan interface{}
	}
	var tests []struct {
		name string
		args args
		want <-chan interface{}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := or(tt.args.ch...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("or() = %v, want %v", got, tt.want)
			}
		})
	}
}
