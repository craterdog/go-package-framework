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

var functionsClass = &functionsClass_{
	// TBA - Assign constant values.
}

// Function

func Functions() FunctionsClassLike {
	return functionsClass
}

// CLASS METHODS

// Target

type functionsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *functionsClass_) MakeWithAttributes(sequence col.Sequential[FunctionLike]) FunctionsLike {
	return &functions_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type functions_ struct {
	sequence_ col.Sequential[FunctionLike]
}

// Attributes

func (v *functions_) GetSequence() col.Sequential[FunctionLike] {
	return v.sequence_
}

// Public

// Private
