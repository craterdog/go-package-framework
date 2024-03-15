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

var declarationClass = &declarationClass_{
	// TBA - Assign constant values.
}

// Function

func Declaration() DeclarationClassLike {
	return declarationClass
}

// CLASS METHODS

// Target

type declarationClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *declarationClass_) MakeWithAttributes(
	comment string,
	identifier string,
	parameters ParametersLike,
) DeclarationLike {
	return &declaration_{
		comment_:    comment,
		identifier_: identifier,
		parameters_: parameters,
	}
}

// Functions

// INSTANCE METHODS

// Target

type declaration_ struct {
	comment_    string
	identifier_ string
	parameters_ ParametersLike
}

// Attributes

func (v *declaration_) GetComment() string {
	return v.comment_
}

func (v *declaration_) GetIdentifier() string {
	return v.identifier_
}

func (v *declaration_) GetParameters() ParametersLike {
	return v.parameters_
}

// Public

// Private
