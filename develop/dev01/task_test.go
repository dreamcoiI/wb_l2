package main

import "testing"

func TestGetTime(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTime()
		})
	}
}
