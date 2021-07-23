package main

import (
	"strconv"
)

type FizzbuzzOptions struct {
	Int1 int    `json:"int1" validate:"min=1"`
	Int2 int    `json:"int2" validate:"min=1"`
	Str1 string `json:"str1"`
	Str2 string `json:"str2"`
}

var (
	DefaultFizzbuzzOptions = FizzbuzzOptions{
		Int1: 3,
		Int2: 5,
		Str1: "fizz",
		Str2: "buzz",
	}
)

func (options *FizzbuzzOptions) fillDefaults() {
	if options.Int1 == -1 {
		options.Int1 = DefaultFizzbuzzOptions.Int1
	}
	if options.Int2 == -1 {
		options.Int2 = DefaultFizzbuzzOptions.Int2
	}
	if options.Str1 == "" {
		options.Str1 = DefaultFizzbuzzOptions.Str1
	}
	if options.Str2 == "" {
		options.Str2 = DefaultFizzbuzzOptions.Str2
	}
}

func getFizzbuzzValue(i int, options FizzbuzzOptions) string {
	fizz := options.Str1
	buzz := options.Str2

	if i%options.Int1 == 0 && i%options.Int2 == 0 {
		return fizz + buzz
	} else if i%options.Int1 == 0 {
		return fizz
	} else if i%options.Int2 == 0 {
		return buzz
	}
	return strconv.Itoa(i)
}

func fizzbuzz(limit int, options FizzbuzzOptions) string {
	retval := ""
	if options.Int1 == 0 || options.Int2 == 0 {
		panic("int1 and int2 cannot be 0")
	}
	for i := 1; i <= limit; i++ {
		retval += getFizzbuzzValue(i, options) + ","
	}
	return retval
}

func parseInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		value = -1
	}
	return value
}
