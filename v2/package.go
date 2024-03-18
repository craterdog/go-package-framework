/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
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

var packageClass = &packageClass_{
	// TBA - Assign constant values.
}

// Function

func Package() PackageClassLike {
	return packageClass
}

// CLASS METHODS

// Target

type packageClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *packageClass_) MakeWithAttributes(
	notice NoticeLike,
	header HeaderLike,
	imports ImportsLike,
	types TypesLike,
	interfaces InterfacesLike,
) PackageLike {
	return &package_{
		notice_:     notice,
		header_:     header,
		imports_:    imports,
		types_:      types,
		interfaces_: interfaces,
	}
}

// Functions

// INSTANCE METHODS

// Target

type package_ struct {
	notice_     NoticeLike
	header_     HeaderLike
	imports_    ImportsLike
	interfaces_ InterfacesLike
	types_      TypesLike
}

// Attributes

func (v *package_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *package_) GetHeader() HeaderLike {
	return v.header_
}

func (v *package_) GetImports() ImportsLike {
	return v.imports_
}

func (v *package_) GetInterfaces() InterfacesLike {
	return v.interfaces_
}

func (v *package_) GetTypes() TypesLike {
	return v.types_
}

// Public

// Private
