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

var valuesClass = &valuesClass_{
	// TBA - Assign constant values.
}

// Function

func Values() ValuesClassLike {
	return valuesClass
}

// CLASS METHODS

// Target

type valuesClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *valuesClass_) MakeWithAttributes(parameter ParameterLike, sequence col.Sequential[string]) ValuesLike {
	return &values_{
		parameter_: parameter,
		sequence_:  sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type values_ struct {
	parameter_ ParameterLike
	sequence_  col.Sequential[string]
}

// Attributes

func (v *values_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *values_) GetSequence() col.Sequential[string] {
	return v.sequence_
}

// Public

// Private
