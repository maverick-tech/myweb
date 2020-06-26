package main

import "testing"

func TestGetPort(test *testing.T){
	if GetPort() == "" {
		test.Error("Empty port")
	}
}
