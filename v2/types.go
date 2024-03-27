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

var typesClass = &typesClass_{
	// TBA - Assign constant values.
}

// Function

func Types() TypesClassLike {
	return typesClass
}

// CLASS METHODS

// Target

type typesClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *typesClass_) MakeWithAttributes(specializations SpecializationsLike, functionals FunctionalsLike) TypesLike {
	return &types_{
		specializations_: specializations,
		functionals_:     functionals,
	}
}

// Functions

// INSTANCE METHODS

// Target

type types_ struct {
	specializations_ SpecializationsLike
	functionals_     FunctionalsLike
}

// Attributes

func (v *types_) GetSpecializations() SpecializationsLike {
	return v.specializations_
}

func (v *types_) GetFunctionals() FunctionalsLike {
	return v.functionals_
}

// Public

// Private
