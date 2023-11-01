package main

import "testing"

func Test_cd(t *testing.T) {
	type args struct {
		requst []string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cd(tt.args.requst)
		})
	}
}

func Test_echo(t *testing.T) {
	type args struct {
		request []string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			echo(tt.args.request)
		})
	}
}

func Test_kill(t *testing.T) {
	type args struct {
		request []string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kill(tt.args.request)
		})
	}
}

func Test_ps(t *testing.T) {
	type args struct {
		request []string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps(tt.args.request)
		})
	}
}

func Test_pwd(t *testing.T) {
	type args struct {
		request []string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwd(tt.args.request)
		})
	}
}
