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

var noticeClass = &noticeClass_{
	// TBA - Assign constant values.
}

// Function

func Notice() NoticeClassLike {
	return noticeClass
}

// CLASS METHODS

// Target

type noticeClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *noticeClass_) MakeWithAttributes(comment string) NoticeLike {
	return &notice_{
		comment_: comment,
	}
}

// Functions

// INSTANCE METHODS

// Target

type notice_ struct {
	comment_ string
}

// Attributes

func (v *notice_) GetComment() string {
	return v.comment_
}

// Public

// Private
