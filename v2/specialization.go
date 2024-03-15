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

var specializationClass = &specializationClass_{
	// TBA - Assign constant values.
}

// Function

func Specialization() SpecializationClassLike {
	return specializationClass
}

// CLASS METHODS

// Target

type specializationClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *specializationClass_) MakeWithAttributes(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	enumeration EnumerationLike,
) SpecializationLike {
	return &specialization_{
		declaration_: declaration,
		abstraction_: abstraction,
		enumeration_: enumeration,
	}
}

// Functions

// INSTANCE METHODS

// Target

type specialization_ struct {
	abstraction_ AbstractionLike
	declaration_ DeclarationLike
	enumeration_ EnumerationLike
}

// Attributes

func (v *specialization_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *specialization_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *specialization_) GetEnumeration() EnumerationLike {
	return v.enumeration_
}

// Public

// Private
