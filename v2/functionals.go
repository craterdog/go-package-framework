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

var functionalsClass = &functionalsClass_{
	// TBA - Assign constant values.
}

// Function

func Functionals() FunctionalsClassLike {
	return functionalsClass
}

// CLASS METHODS

// Target

type functionalsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *functionalsClass_) MakeWithAttributes(sequence col.Sequential[FunctionalLike]) FunctionalsLike {
	return &functionals_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type functionals_ struct {
	sequence_ col.Sequential[FunctionalLike]
}

// Attributes

func (v *functionals_) GetSequence() col.Sequential[FunctionalLike] {
	return v.sequence_
}

// Public

// Private
