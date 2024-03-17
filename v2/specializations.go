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

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var specializationsClass = &specializationsClass_{
	// TBA - Assign constant values.
}

// Function

func Specializations() SpecializationsClassLike {
	return specializationsClass
}

// CLASS METHODS

// Target

type specializationsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *specializationsClass_) MakeWithAttributes(sequence col.Sequential[SpecializationLike]) SpecializationsLike {
	return &specializations_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type specializations_ struct {
	sequence_ col.Sequential[SpecializationLike]
}

// Attributes

func (v *specializations_) GetSequence() col.Sequential[SpecializationLike] {
	return v.sequence_
}

// Public

// Private
