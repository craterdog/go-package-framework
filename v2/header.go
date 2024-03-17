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

var headerClass = &headerClass_{
	// TBA - Assign constant values.
}

// Function

func Header() HeaderClassLike {
	return headerClass
}

// CLASS METHODS

// Target

type headerClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *headerClass_) MakeWithAttributes(comment string, identifier string) HeaderLike {
	return &header_{
		comment_:    comment,
		identifier_: identifier,
	}
}

// Functions

// INSTANCE METHODS

// Target

type header_ struct {
	comment_    string
	identifier_ string
}

// Attributes

func (v *header_) GetComment() string {
	return v.comment_
}

func (v *header_) GetIdentifier() string {
	return v.identifier_
}

// Public

// Private
