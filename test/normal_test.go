package test

import (
	"fmt"
	"testing"
)
type testt struct{
	name string
	id int
}
func TestNormal(t *testing.T){
	a := make(map[string]testt)
	test1 := testt{"123123",1}
	a["123"] = test1
	fmt.Println(a)
	delete(a,"123")
	fmt.Println(a)
}