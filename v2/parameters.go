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

var parametersClass = &parametersClass_{
	// TBA - Assign constant values.
}

// Function

func Parameters() ParametersClassLike {
	return parametersClass
}

// CLASS METHODS

// Target

type parametersClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *parametersClass_) MakeWithAttributes(sequence col.Sequential[ParameterLike]) ParametersLike {
	return &parameters_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type parameters_ struct {
	sequence_ col.Sequential[ParameterLike]
}

// Attributes

func (v *parameters_) GetSequence() col.Sequential[ParameterLike] {
	return v.sequence_
}

// Public

// Private
