package main

import "testing"

func TestGenCode(t *testing.T){
	var trimstr=""
	g:=Generator{
		trimPrefix: trimstr,
	}
	g.parsePackage([]string{"/Users/menglei/workspace/iam/internal/pkg/code"},[]string{"int"})
}