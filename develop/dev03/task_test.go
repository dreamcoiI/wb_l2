package main

import (
	"reflect"
	"testing"
)

func Test_getFile(t *testing.T) {
	type args struct {
		path   string
		reader bool
	}
	var tests []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFile(tt.args.path, tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortFile(t *testing.T) {
	type args struct {
		data []string
		flag bool
	}
	var tests []struct {
		name string
		args args
		want []string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortFile(tt.args.data, tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
