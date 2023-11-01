package main

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_echo(t *testing.T) {
	type args struct {
		file   []string
		i      int
		after  int
		before int
		number bool
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			echo(tt.args.file, tt.args.i, tt.args.after, tt.args.before, tt.args.number)
		})
	}
}

func Test_getExpression(t *testing.T) {
	type args struct {
		pattern string
		ignore  bool
	}
	var tests []struct {
		name    string
		args    args
		want    *regexp.Regexp
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getExpression(tt.args.pattern, tt.args.ignore)
			if (err != nil) != tt.wantErr {
				t.Errorf("getExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumberOfIntersections(t *testing.T) {
	type args struct {
		file       []string
		expression *regexp.Regexp
	}
	var tests []struct {
		name string
		args args
		want int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfIntersections(tt.args.file, tt.args.expression); got != tt.want {
				t.Errorf("getNumberOfIntersections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openFile(t *testing.T) {
	type args struct {
		path string
	}
	var tests []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reg(t *testing.T) {
	type args struct {
		file       []string
		expression *regexp.Regexp
		after      int
		before     int
		number     bool
		invert     bool
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg(tt.args.file, tt.args.expression, tt.args.after, tt.args.before, tt.args.number, tt.args.invert)
		})
	}
}
