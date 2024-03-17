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

var importsClass = &importsClass_{
	// TBA - Assign constant values.
}

// Function

func Imports() ImportsClassLike {
	return importsClass
}

// CLASS METHODS

// Target

type importsClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *importsClass_) MakeWithAttributes(modules ModulesLike) ImportsLike {
	return &imports_{
		modules_: modules,
	}
}

// Functions

// INSTANCE METHODS

// Target

type imports_ struct {
	modules_ ModulesLike
}

// Attributes

func (v *imports_) GetModules() ModulesLike {
	return v.modules_
}

// Public

// Private
