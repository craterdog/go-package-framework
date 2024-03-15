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

var attributesClass = &attributesClass_{
	// TBA - Assign constant values.
}

// Function

func Attributes() AttributesClassLike {
	return attributesClass
}

// CLASS METHODS

// Target

type attributesClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *attributesClass_) MakeWithAttributes(sequence col.Sequential[AttributeLike]) AttributesLike {
	return &attributes_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type attributes_ struct {
	sequence_ col.Sequential[AttributeLike]
}

// Attributes

func (v *attributes_) GetSequence() col.Sequential[AttributeLike] {
	return v.sequence_
}

// Public

// Private
