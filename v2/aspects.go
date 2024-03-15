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

var aspectsClass = &aspectsClass_{
	// TBA - Assign constant values.
}

// Function

func Aspects() AspectsClassLike {
	return aspectsClass
}

// CLASS METHODS

// Target

type aspectsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *aspectsClass_) MakeWithAttributes(sequence col.Sequential[AspectLike]) AspectsLike {
	return &aspects_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type aspects_ struct {
	sequence_ col.Sequential[AspectLike]
}

// Attributes

func (v *aspects_) GetSequence() col.Sequential[AspectLike] {
	return v.sequence_
}

// Public

// Private
