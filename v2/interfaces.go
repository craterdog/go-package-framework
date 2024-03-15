/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
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

var interfacesClass = &interfacesClass_{
	// TBA - Assign constant values.
}

// Function

func Interfaces() InterfacesClassLike {
	return interfacesClass
}

// CLASS METHODS

// Target

type interfacesClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *interfacesClass_) MakeWithAttributes(
	aspects AspectsLike,
	classes ClassesLike,
	instances InstancesLike,
) InterfacesLike {
	return &interfaces_{
		aspects_:   aspects,
		classes_:   classes,
		instances_: instances,
	}
}

// Functions

// INSTANCE METHODS

// Target

type interfaces_ struct {
	aspects_   AspectsLike
	classes_   ClassesLike
	instances_ InstancesLike
}

// Attributes

func (v *interfaces_) GetAspects() AspectsLike {
	return v.aspects_
}

func (v *interfaces_) GetClasses() ClassesLike {
	return v.classes_
}

func (v *interfaces_) GetInstances() InstancesLike {
	return v.instances_
}

// Public

// Private
