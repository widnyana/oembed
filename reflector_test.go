package oembed

import (
	"testing"
	"fmt"
)

func TestGetType(t *testing.T)  {
	var result string
	var strVal string = `string`
	var intval int = 100
	var uintVal uint = 100
	var fl64Val float64 = 100.2
	// too lazy to add more. pardon me

	result = getType(strVal)
	if result != "string" {
		t.Fail()
	}

	result = getType(intval)
	if result != "int" {
		t.Fail()
	}

	result = getType(uintVal)
	if result != "uint" {
		t.Fail()
	}

	result = getType(fl64Val)
	if result != "float64" {
		t.Fail()
	}
}

func TestToUInt64(t *testing.T) {
	var result uint64
	var one string = "1"
	var twopointone float64 = 2.1

	result = toUInt64(one)
	if result != 1 {
		t.Fail()
	}

	result = toUInt64(twopointone)
	if result != 2 {
		t.Fatal(fmt.Sprintf("Expecting %v got %v", 2, result))
	}
}