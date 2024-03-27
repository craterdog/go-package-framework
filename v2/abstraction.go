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

import ()

// CLASS ACCESS

// Reference

var abstractionClass = &abstractionClass_{
	// TBA - Assign constant values.
}

// Function

func Abstraction() AbstractionClassLike {
	return abstractionClass
}

// CLASS METHODS

// Target

type abstractionClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *abstractionClass_) MakeWithAttributes(
	prefix PrefixLike,
	identifier string,
	arguments ArgumentsLike,
) AbstractionLike {
	return &abstraction_{
		prefix_:     prefix,
		identifier_: identifier,
		arguments_:  arguments,
	}
}

// Functions

// INSTANCE METHODS

// Target

type abstraction_ struct {
	prefix_     PrefixLike
	identifier_ string
	arguments_  ArgumentsLike
}

// Attributes

func (v *abstraction_) GetPrefix() PrefixLike {
	return v.prefix_
}

func (v *abstraction_) GetIdentifier() string {
	return v.identifier_
}

func (v *abstraction_) GetArguments() ArgumentsLike {
	return v.arguments_
}

// Public

// Private
