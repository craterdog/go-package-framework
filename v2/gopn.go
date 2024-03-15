/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package packages

import ()

// CLASS ACCESS

// Reference

var goPNClass = &goPNClass_{
	// TBA - Assign constant values.
}

// Function

func GoPN() GoPNClassLike {
	return goPNClass
}

// CLASS METHODS

// Target

type goPNClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *goPNClass_) MakeWithAttributes(
	copyright CopyrightLike,
	header HeaderLike,
	imports ImportsLike,
	types TypesLike,
	interfaces InterfacesLike,
) GoPNLike {
	return &goPN_{
		copyright_:  copyright,
		header_:     header,
		imports_:    imports,
		types_:      types,
		interfaces_: interfaces,
	}
}

// Functions

// INSTANCE METHODS

// Target

type goPN_ struct {
	copyright_  CopyrightLike
	header_     HeaderLike
	imports_    ImportsLike
	interfaces_ InterfacesLike
	types_      TypesLike
}

// Attributes

func (v *goPN_) GetCopyright() CopyrightLike {
	return v.copyright_
}

func (v *goPN_) GetHeader() HeaderLike {
	return v.header_
}

func (v *goPN_) GetImports() ImportsLike {
	return v.imports_
}

func (v *goPN_) GetInterfaces() InterfacesLike {
	return v.interfaces_
}

func (v *goPN_) GetTypes() TypesLike {
	return v.types_
}

// Public

// Private
