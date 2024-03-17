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

const classTemplate_ = `<Copyright><Header><Imports><Access><Class><Instance>`

const headerTemplate_ = `package <PackageName>`

const importsTemplate_ = `
import (<Modules>)`

const classAccessTemplate_ = `
// CLASS ACCESS

// Reference
<Reference>

// Function
<Function>`

const classReferenceTemplate_ = `
var <TargetName>Class = &<TargetName>Class_{
	// TBA - Assign constant values.
}`

const genericReferenceTemplate_ = `
var <TargetName>Class = map[string]any{}
var <TargetName>Mutex syn.Mutex`

const classFunctionTemplate_ = `
func <ClassName>() <ClassName>ClassLike {
	return <TargetName>Class
}`

const genericFunctionTemplate_ = `
func <ClassName>[<Parameters>]() <ClassName>ClassLike[<Arguments>] {
	// Generate the name of the bound class type.
	var result_ <ClassName>ClassLike[<Arguments>]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	<TargetName>Mutex.Lock()
	var value = <TargetName>Class[name]
	switch actual := value.(type) {
	case *<TargetName>Class_[<Arguments>]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &<TargetName>Class_[<Arguments>]{
			// TBA - Assign constant values.
		}
		<TargetName>Class[name] = result_
	}
	<TargetName>Mutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}`

const classMethodsTemplate_ = `
// CLASS METHODS

// Target
<Target>
// Constants
<Constants>
// Constructors
<Constructors>
// Functions
<Functions>`

const classTargetTemplate_ = `
type <TargetName>Class_[<Parameters>] struct {<Constants>}`

const classConstantTemplate_ = `
	<ConstantName>_ <ConstantType>`

const classMethodTemplate_ = `
func (c *<TargetName>Class_[<Arguments>]) <MethodName>(<Parameters>)<ResultType> {<Body>}`

const constantBodyTemplate_ = `
	return c.<ConstantName>_
`

const constructorBodyTemplate_ = `
	return &<TargetName>_[<Arguments>]{<Assignments>}
`

const attributeAssignmentTemplate_ = `
		<AttributeName>_: <ParameterName>,`

const functionBodyTemplate_ = `
	var result_<ResultType>
	// TBA - Implement the function.
	return result_
`

const instanceMethodsTemplate_ = `
// INSTANCE METHODS

// Target
<Target>
// Attributes
<Attributes><Abstractions>
// Public
<Methods>
// Private
`

const instanceTargetTemplate_ = `
type <TargetName>_[<Parameters>] struct {<Attributes>}`

const instanceAttributeTemplate_ = `
	<AttributeName>_ <AttributeType>`

const instanceAspectTemplate_ = `
// <AspectName>
<Methods>`

const instanceMethodTemplate_ = `
func (v *<TargetName>_[<Arguments>]) <MethodName>(<Parameters>)<ResultType> {<Body>}`

const methodBodyTemplate_ = `
	// TBA - Implement the method.
`

const resultBodyTemplate_ = `
	var result_<ResultType>
	// TBA - Implement the method.
	return result_
`

const returnBodyTemplate_ = `
	// TBA - Implement the method.
	return
`

const getterBodyTemplate_ = `
	return v.<AttributeName>_
`

const setterBodyTemplate_ = `
	v.<AttributeName>_ = <ParameterName>
`

const modelTemplate_ = `
/*
................................................................................
<Notice>
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

/*
Package <packagename> provides...

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/<repository-name>/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-package-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package <packagename>

// TYPES

// Specializations

/*
<SpecializedType> is a specialized type representing...
*/
type <SpecializedType> <primitiveType>

const (
	<1stValue> <SpecializedType> = iota
	<2ndValue>
	<3rdValue>
	...
)
...

// Functionals

/*
<FunctionName>Function is a functional type that defines the signature for any
function that...
*/
type <FunctionName>Function func(<Parameters>) <AbstractType>
...

// INTERFACES

// Aspects

/*
<AspectName> is an aspect interface that defines the set of method signatures
that must be supported by each instance of a <aspect name> concrete class.
*/
type <AspectName> interface {
	// Methods
	<MethodName>(<Parameters>) <AbstractType>
	...
}
...

// Classes

/*
<ClassName>ClassLike is a class interface that defines the set of class
constants, constructors and functions that must be supported by each
<class-name>-like concrete class.
*/
type <ClassName>ClassLike interface {
	// Constants
	<ConstantName>() <AbstractType>
	...

	// Constructors
	Make<FromContext>(<Parameters>) <ClassName>Like
	...

	// Functions
	<FunctionName>(<Parameters>) <AbstractType>
	...
}
...

// Instances

/*
<ClassName>Like is an instance interface that defines the complete set of
abstractions and methods that must be supported by each instance of a
<class-name>-like concrete class.
*/
type <ClassName>Like interface {
	// Attributes
	Get<AttributeName>() <AttributeType>
	Set<AttributeName>(<attributeName> <AttributeType>)
	...

	// Abstractions
	<AspectName>
	...

	// Methods
	<MethodName>(<Parameters>) <AbstractType>
	...
}
...
`
