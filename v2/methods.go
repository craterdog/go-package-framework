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

var methodsClass = &methodsClass_{
	// TBA - Assign constant values.
}

// Function

func Methods() MethodsClassLike {
	return methodsClass
}

// CLASS METHODS

// Target

type methodsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *methodsClass_) MakeWithAttributes(sequence col.Sequential[MethodLike]) MethodsLike {
	return &methods_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type methods_ struct {
	sequence_ col.Sequential[MethodLike]
}

// Attributes

func (v *methods_) GetSequence() col.Sequential[MethodLike] {
	return v.sequence_
}

// Public

// Private
