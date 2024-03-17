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

var copyrightClass = &copyrightClass_{
	// TBA - Assign constant values.
}

// Function

func Copyright() CopyrightClassLike {
	return copyrightClass
}

// CLASS METHODS

// Target

type copyrightClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *copyrightClass_) MakeWithAttributes(comment string) CopyrightLike {
	return &copyright_{
		comment_: comment,
	}
}

// Functions

// INSTANCE METHODS

// Target

type copyright_ struct {
	comment_ string
}

// Attributes

func (v *copyright_) GetComment() string {
	return v.comment_
}

// Public

// Private
