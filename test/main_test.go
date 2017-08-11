package main

import "testing"

func TestSuccessDemo(t *testing.T) {
	if result := successFunc(); result != "success" {
		t.Error(result)
	}
}

func successFunc() string {
	return "success"
}
