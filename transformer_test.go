package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHappy(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestNotHappy(t *testing.T) {
	assert.Equal(t, 1, 2)
}

func TestVery____Happy(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestTheDNSCheck(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestUnwanted(t *testing.T) {
	t.Skip("I don't want to run this test")
}

func ExampleItem_PrettyName() {
	i := Item{Test: "TestThisIsATest"}
	fmt.Println(i.PrettyName())
	// Output: This is a test
}

func ExampleItem_PrettyName_simple() {
	i := Item{Test: "TestCamelCaseConversion/YesItWorks"}
	fmt.Println(i.PrettyName())
	// Output: Camel case conversion Yes it works
}

func Example_convertCamelCase() {
	fmt.Println(convertCamelCase("CamelCaseConversion"))
	// Output: Camel case conversion
}

func Example_convertCamelCase_underscore() {
	fmt.Println(convertCamelCase("CamelCase_Conversion"))
	// Output: Camel case conversion
}
func Example_convertCamelCase_abbr() {
	fmt.Println(convertCamelCase("CamelCaseDNSCheck"))
	// Output: Camel case dns check
}
func Example_convertCamelCase_abbrOnTheEnd() {
	fmt.Println(convertCamelCase("CamelCaseDNS"))
	// Output: Camel case dns
}

func Example_convertCamelCase_abbrOnTheStart() {
	fmt.Println(convertCamelCase("DNSCheckCamelCase"))
	// Output: Dns check camel case
}

func Example_convertCamelCase_abbrOnTheStartAndEnd() {
	fmt.Println(convertCamelCase("DNSCamelCaseDNS"))
	// Output: Dns camel case dns
}

func Example_convertCamelCase_multi_underscores() {
	fmt.Println(convertCamelCase("CamelCase_____Conversion"))
	// Output: Camel case conversion
}
