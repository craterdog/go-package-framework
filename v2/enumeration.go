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

var enumerationClass = &enumerationClass_{
	// TBA - Assign constant values.
}

// Function

func Enumeration() EnumerationClassLike {
	return enumerationClass
}

// CLASS METHODS

// Target

type enumerationClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *enumerationClass_) MakeWithAttributes(values ValuesLike) EnumerationLike {
	return &enumeration_{
		values_: values,
	}
}

// Functions

// INSTANCE METHODS

// Target

type enumeration_ struct {
	values_ ValuesLike
}

// Attributes

func (v *enumeration_) GetValues() ValuesLike {
	return v.values_
}

// Public

// Private
