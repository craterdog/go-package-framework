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

var constantsClass = &constantsClass_{
	// TBA - Assign constant values.
}

// Function

func Constants() ConstantsClassLike {
	return constantsClass
}

// CLASS METHODS

// Target

type constantsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *constantsClass_) MakeWithAttributes(sequence col.Sequential[ConstantLike]) ConstantsLike {
	return &constants_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constants_ struct {
	sequence_ col.Sequential[ConstantLike]
}

// Attributes

func (v *constants_) GetSequence() col.Sequential[ConstantLike] {
	return v.sequence_
}

// Public

// Private
