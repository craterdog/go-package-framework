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

var classesClass = &classesClass_{
	// TBA - Assign constant values.
}

// Function

func Classes() ClassesClassLike {
	return classesClass
}

// CLASS METHODS

// Target

type classesClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *classesClass_) MakeWithAttributes(sequence col.Sequential[ClassLike]) ClassesLike {
	return &classes_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type classes_ struct {
	sequence_ col.Sequential[ClassLike]
}

// Attributes

func (v *classes_) GetSequence() col.Sequential[ClassLike] {
	return v.sequence_
}

// Public

// Private
