/*
................................................................................
.    Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package packages_test

import (
	fmt "fmt"
	pac "github.com/craterdog/go-package-framework/v2"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestRoundtrips(t *tes.T) {
	var files, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var parser = pac.Parser().Make()
		var validator = pac.Validator().Make()
		var formatter = pac.Formatter().Make()
		var filename = testDirectory + file.Name()
		if sts.HasSuffix(filename, ".gopn") {
			fmt.Println(filename)
			var bytes, err = osx.ReadFile(filename)
			if err != nil {
				panic(err)
			}
			var expected = string(bytes)
			var gopn = parser.ParseSource(expected)
			validator.ValidatePackage(gopn)
			var actual = formatter.FormatGoPN(gopn)
			ass.Equal(t, expected, actual)
		}
	}
}
