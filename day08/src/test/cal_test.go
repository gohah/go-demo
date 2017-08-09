package main

import "testing"

func TestAdd(t *testing.T) {
	r := add(10,20)

	if r != 30 {
		t.Fatal("error")
	}

	t.Logf("test add success")
}