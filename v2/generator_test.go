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
	osx "os"
	sts "strings"
	tes "testing"
)

const generatedDirectory = "./generated/"

func TestGeneration(t *tes.T) {
	var generator = pac.Generator().Make()

	var files, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var fileSuffix = ".gopn"
		var fileName = sts.TrimSuffix(file.Name(), fileSuffix)
		var bytes, err = osx.ReadFile(testDirectory + file.Name())
		if err != nil {
			panic(err)
		}
		var directoryName = generatedDirectory + fileName + "/"
		err = osx.RemoveAll(directoryName)
		if err != nil {
			panic(err)
		}
		err = osx.MkdirAll(directoryName, 0755)
		if err != nil {
			panic(err)
		}
		err = osx.WriteFile(directoryName+"Model.go", bytes, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println(fileName)
		generator.GeneratePackage(directoryName)
	}
}
