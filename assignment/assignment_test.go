package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	cases := []struct {
		testX, testY, resultInteger uint32
		resultOverflow              bool
	}{
		{math.MaxUint32, 1, 0, true},
		{1, 1, 2, false},
		{42, 2701, 2743, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, false},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
		{4294967290, 10, 4, true},
	}

	for _, c := range cases {
		t.Run("should run TestAddUint32 testing", func(t *testing.T) {
			sum, overflow := AddUint32(c.testX, c.testY)
			assert.Equal(t, c.resultInteger, sum)
			assert.Equal(t, c.resultOverflow, overflow)
		})
	}
}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/

	cases := []struct {
		value, expectedNumber float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}

	t.Run("should test run CeilNumber testing", func(t *testing.T) {
		for _, c := range cases {
			point := CeilNumber(c.value)
			assert.Equal(t, c.expectedNumber, point)
		}
	})

}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/
	cases := []struct {
		value, expectedValue string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	t.Run("should test run AlphabetSoup testing", func(t *testing.T) {
		for _, c := range cases {
			result := AlphabetSoup(c.value)
			assert.Equal(t, c.expectedValue, result)
		}
	})
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/

	cases := []struct {
		value          string
		number         uint
		expectedResult string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 3, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	t.Run("should run test StringMask testing", func(t *testing.T) {
		for _, c := range cases {
			result := StringMask(c.value, c.number)
			assert.Equal(t, c.expectedResult, result)
		}
	})
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/

	cases := []struct {
		arr           [2]string
		expectedValue string
	}{
		{[2]string{"hellocat", words}, "hello,cat"},
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks", words}, "not possible"},
	}

	t.Run("should test WordSplit testing", func(t *testing.T) {
		for _, c := range cases {
			result := WordSplit(c.arr)
			assert.Equal(t, c.expectedValue, result)
		}
	})

}

type ReallyDance []interface{}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks!" => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
	cases := []struct {
		set    ReallyDance
		result []interface{}
	}{
		{ReallyDance{4, 2, 5, 4, 2, 4}, []interface{}{4, 2, 5}},
		{ReallyDance{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{ReallyDance{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}

	for _, c := range cases {
		set := VariadicSet(c.set)
		assert.Equal(t, c.result, set)
	}

}
