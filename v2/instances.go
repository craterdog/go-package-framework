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

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var instancesClass = &instancesClass_{
	// TBA - Assign constant values.
}

// Function

func Instances() InstancesClassLike {
	return instancesClass
}

// CLASS METHODS

// Target

type instancesClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *instancesClass_) MakeWithAttributes(sequence col.Sequential[InstanceLike]) InstancesLike {
	return &instances_{
		sequence_: sequence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type instances_ struct {
	sequence_ col.Sequential[InstanceLike]
}

// Attributes

func (v *instances_) GetSequence() col.Sequential[InstanceLike] {
	return v.sequence_
}

// Public

// Private
