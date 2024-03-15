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

import ()

// CLASS ACCESS

// Reference

var constructorClass = &constructorClass_{
	// TBA - Assign constant values.
}

// Function

func Constructor() ConstructorClassLike {
	return constructorClass
}

// CLASS METHODS

// Target

type constructorClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *constructorClass_) MakeWithAttributes(
	identifier string,
	parameters ParametersLike,
	abstraction AbstractionLike,
) ConstructorLike {
	return &constructor_{
		identifier_:  identifier,
		parameters_:  parameters,
		abstraction_: abstraction,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constructor_ struct {
	abstraction_ AbstractionLike
	identifier_  string
	parameters_  ParametersLike
}

// Attributes

func (v *constructor_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *constructor_) GetIdentifier() string {
	return v.identifier_
}

func (v *constructor_) GetParameters() ParametersLike {
	return v.parameters_
}

// Public

// Private
