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

var aspectClass = &aspectClass_{
	// TBA - Assign constant values.
}

// Function

func Aspect() AspectClassLike {
	return aspectClass
}

// CLASS METHODS

// Target

type aspectClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *aspectClass_) MakeWithAttributes(declaration DeclarationLike, methods MethodsLike) AspectLike {
	return &aspect_{
		declaration_: declaration,
		methods_:     methods,
	}
}

// Functions

// INSTANCE METHODS

// Target

type aspect_ struct {
	declaration_ DeclarationLike
	methods_     MethodsLike
}

// Attributes

func (v *aspect_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *aspect_) GetMethods() MethodsLike {
	return v.methods_
}

// Public

// Private
