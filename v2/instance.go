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

var instanceClass = &instanceClass_{
	// TBA - Assign constant values.
}

// Function

func Instance() InstanceClassLike {
	return instanceClass
}

// CLASS METHODS

// Target

type instanceClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *instanceClass_) MakeWithAttributes(
	declaration DeclarationLike,
	attributes AttributesLike,
	abstractions AbstractionsLike,
	methods MethodsLike,
) InstanceLike {
	return &instance_{
		declaration_:  declaration,
		attributes_:   attributes,
		abstractions_: abstractions,
		methods_:      methods,
	}
}

// Functions

// INSTANCE METHODS

// Target

type instance_ struct {
	declaration_  DeclarationLike
	attributes_   AttributesLike
	abstractions_ AbstractionsLike
	methods_      MethodsLike
}

// Attributes

func (v *instance_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instance_) GetAttributes() AttributesLike {
	return v.attributes_
}

func (v *instance_) GetAbstractions() AbstractionsLike {
	return v.abstractions_
}

func (v *instance_) GetMethods() MethodsLike {
	return v.methods_
}

// Public

// Private
