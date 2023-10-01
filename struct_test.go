package intro_test

import (
	"reflect"
	"testing"
)

type profile struct {
	Name string
	Age  int
}

func TestStruct(t *testing.T) {
	var p0 *profile
	p1 := new(profile)

	if reflect.DeepEqual(p0, p1) {
		t.Errorf("compare value zero and ref to an object")
	}
}

func TestStructPointZero(t *testing.T) {
	var p0 *profile
	p2 := &profile{
		Name: "Tani",
		Age:  21,
	}

	if reflect.DeepEqual(p0, p2) {
		t.Errorf("error")
	}
}

func TestStructZero(t *testing.T) {
	p1 := new(profile)
	p2 := &profile{
		Name: "Tani",
		Age:  21,
	}

	if reflect.DeepEqual(p1, p2) {
		t.Errorf("error")
	}
}

func TestStructEquality(t *testing.T) {
	var p1 *profile
	p2 := &profile{
		Name: "Tani",
		Age:  21,
	}
	p1 = p2

	if p1 != p2 {
		t.Errorf("Equality")
	}
	if !reflect.DeepEqual(p1, p2) {
		t.Errorf("Equivalent")
	}
}

func TestStructEquivalent(t *testing.T) {
	p1 := &profile{
		Name: "Tani",
		Age:  21,
	}
	p2 := &profile{
		Name: "Tani",
		Age:  21,
	}

	if p1 == p2 {
		t.Errorf("Equality")
	}
	if !reflect.DeepEqual(p1, p2) {
		t.Errorf("Equivalent")
	}
}

func TestAaa(t *testing.T) {
	p := new(profile)
	p.Name = "Senda"
	p.Age = 22

	if p.Name == "" {
		t.Errorf("")
	}
}
