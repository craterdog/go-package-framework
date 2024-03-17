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

var abstractionsClass = &abstractionsClass_{
	// TBA - Assign constant values.
}

// Function

func Abstractions() AbstractionsClassLike {
	return abstractionsClass
}

// CLASS METHODS

// Target

type abstractionsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *abstractionsClass_) MakeWithAttributes(sequence col.Sequential[AbstractionLike]) AbstractionsLike {
	return &abstractions_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type abstractions_ struct {
	sequence_ col.Sequential[AbstractionLike]
}

// Attributes

func (v *abstractions_) GetSequence() col.Sequential[AbstractionLike] {
	return v.sequence_
}

// Public

// Private
