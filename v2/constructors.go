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

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var constructorsClass = &constructorsClass_{
	// TBA - Assign constant values.
}

// Function

func Constructors() ConstructorsClassLike {
	return constructorsClass
}

// CLASS METHODS

// Target

type constructorsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *constructorsClass_) MakeWithAttributes(sequence col.Sequential[ConstructorLike]) ConstructorsLike {
	return &constructors_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constructors_ struct {
	sequence_ col.Sequential[ConstructorLike]
}

// Attributes

func (v *constructors_) GetSequence() col.Sequential[ConstructorLike] {
	return v.sequence_
}

// Public

// Private
