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

var constantClass = &constantClass_{
	// TBA - Assign constant values.
}

// Function

func Constant() ConstantClassLike {
	return constantClass
}

// CLASS METHODS

// Target

type constantClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *constantClass_) MakeWithAttributes(identifier string, abstraction AbstractionLike) ConstantLike {
	return &constant_{
		identifier_:  identifier,
		abstraction_: abstraction,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constant_ struct {
	abstraction_ AbstractionLike
	identifier_  string
}

// Attributes

func (v *constant_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *constant_) GetIdentifier() string {
	return v.identifier_
}

// Public

// Private
