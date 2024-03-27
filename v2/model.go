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

var modelClass = &modelClass_{
	// TBA - Assign constant values.
}

// Function

func Model() ModelClassLike {
	return modelClass
}

// CLASS METHODS

// Target

type modelClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *modelClass_) MakeWithAttributes(
	notice NoticeLike,
	header HeaderLike,
	imports ImportsLike,
	types TypesLike,
	interfaces InterfacesLike,
) ModelLike {
	return &model_{
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

type model_ struct {
	notice_     NoticeLike
	header_     HeaderLike
	imports_    ImportsLike
	types_      TypesLike
	interfaces_ InterfacesLike
}

// Attributes

func (v *model_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *model_) GetHeader() HeaderLike {
	return v.header_
}

func (v *model_) GetImports() ImportsLike {
	return v.imports_
}

func (v *model_) GetTypes() TypesLike {
	return v.types_
}

func (v *model_) GetInterfaces() InterfacesLike {
	return v.interfaces_
}

// Public

// Private
