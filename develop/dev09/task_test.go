package main

import (
	"net/http"
	"os"
	"reflect"
	"testing"
)

func Test_createFile(t *testing.T) {
	type args struct {
		filename string
	}
	var tests []struct {
		name string
		args args
		want *os.File
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getData(t *testing.T) {
	type args struct {
		urlPath string
		client  *http.Client
		file    *os.File
	}
	var tests []struct {
		name string
		args args
		want int64
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getData(tt.args.urlPath, tt.args.client, tt.args.file); got != tt.want {
				t.Errorf("getData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getURLandFilename(t *testing.T) {
	var tests []struct {
		name  string
		want  string
		want1 string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getURLandFilename()
			if got != tt.want {
				t.Errorf("getURLandFilename() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getURLandFilename() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
