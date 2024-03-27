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

var classClass = &classClass_{
	// TBA - Assign constant values.
}

// Function

func Class() ClassClassLike {
	return classClass
}

// CLASS METHODS

// Target

type classClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *classClass_) MakeWithAttributes(
	declaration DeclarationLike,
	constants ConstantsLike,
	constructors ConstructorsLike,
	functions FunctionsLike,
) ClassLike {
	return &class_{
		declaration_:  declaration,
		constants_:    constants,
		constructors_: constructors,
		functions_:    functions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type class_ struct {
	declaration_  DeclarationLike
	constants_    ConstantsLike
	constructors_ ConstructorsLike
	functions_    FunctionsLike
}

// Attributes

func (v *class_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *class_) GetConstants() ConstantsLike {
	return v.constants_
}

func (v *class_) GetConstructors() ConstructorsLike {
	return v.constructors_
}

func (v *class_) GetFunctions() FunctionsLike {
	return v.functions_
}

// Public

// Private
