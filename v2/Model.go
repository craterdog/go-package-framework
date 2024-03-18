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

/*
Package packages provides the ability to generate Go class files based on a
Go Model.go file that follows the format shown in the following code template:
  - https://github.com/craterdog/go-package-framework/blob/main/packages/Model.go

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-package-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package packages

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// TYPES

// Specializations

/*
PrefixType is a specialized type representing a prefix type.
*/
type PrefixType uint8

const (
	ErrorPrefix PrefixType = iota
	AliasPrefix
	ArrayPrefix
	ChannelPrefix
	MapPrefix
)

/*
TokenType is a specialized type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	EOFToken
	EOLToken
	IdentifierToken
	NoteToken
	SpaceToken
	TextToken
)

// INTERFACES

// Classes

/*
AbstractionClassLike defines the set of class constants, constructors and
functions that must be supported by all abstraction-class-like classes.
*/
type AbstractionClassLike interface {
	// Constructors
	MakeWithAttributes(
		prefix PrefixLike,
		identifier string,
		arguments ArgumentsLike,
	) AbstractionLike
}

/*
AbstractionsClassLike defines the set of class constants, constructors and
functions that must be supported by all abstractions-class-like classes.
*/
type AbstractionsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[AbstractionLike]) AbstractionsLike
}

/*
ArgumentsClassLike defines the set of class constants, constructors and
functions that must be supported by all arguments-class-like classes.
*/
type ArgumentsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[AbstractionLike]) ArgumentsLike
}

/*
AspectClassLike defines the set of class constants, constructors and
functions that must be supported by all aspect-class-like classes.
*/
type AspectClassLike interface {
	// Constructors
	MakeWithAttributes(declaration DeclarationLike, methods MethodsLike) AspectLike
}

/*
AspectsClassLike defines the set of class constants, constructors and
functions that must be supported by all aspects-class-like classes.
*/
type AspectsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[AspectLike]) AspectsLike
}

/*
AttributeClassLike defines the set of class constants, constructors and
functions that must be supported by all attribute-class-like classes.
*/
type AttributeClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameter ParameterLike,
		abstraction AbstractionLike,
	) AttributeLike
}

/*
AttributesClassLike defines the set of class constants, constructors and
functions that must be supported by all constants-class-like classes.
*/
type AttributesClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[AttributeLike]) AttributesLike
}

/*
ClassClassLike defines the set of class constants, constructors and
functions that must be supported by all class-class-like classes.
*/
type ClassClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		constants ConstantsLike,
		constructors ConstructorsLike,
		functions FunctionsLike,
	) ClassLike
}

/*
ClassesClassLike defines the set of class constants, constructors and
functions that must be supported by all classes-class-like classes.
*/
type ClassesClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[ClassLike]) ClassesLike
}

/*
ConstantClassLike defines the set of class constants, constructors and
functions that must be supported by all constant-class-like classes.
*/
type ConstantClassLike interface {
	// Constructors
	MakeWithAttributes(identifier string, abstraction AbstractionLike) ConstantLike
}

/*
ConstantsClassLike defines the set of class constants, constructors and
functions that must be supported by all constants-class-like classes.
*/
type ConstantsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[ConstantLike]) ConstantsLike
}

/*
ConstructorClassLike defines the set of class constants, constructors and
functions that must be supported by all constructor-class-like classes.
*/
type ConstructorClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		abstraction AbstractionLike,
	) ConstructorLike
}

/*
ConstructorsClassLike defines the set of class constants, constructors and
functions that must be supported by all constructors-class-like classes.
*/
type ConstructorsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[ConstructorLike]) ConstructorsLike
}

/*
DeclarationClassLike defines the set of class constants, constructors and
functions that must be supported by all declaration-class-like classes.
*/
type DeclarationClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
		parameters ParametersLike,
	) DeclarationLike
}

/*
EnumerationClassLike defines the set of class constants, constructors and
functions that must be supported by all enumeration-class-like classes.
*/
type EnumerationClassLike interface {
	// Constructors
	MakeWithAttributes(values ValuesLike) EnumerationLike
}

/*
FormatterClassLike defines the set of class constants, constructors and
functions that must be supported by all formatter-class-like classes.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
FunctionClassLike defines the set of class constants, constructors and
functions that must be supported by all function-class-like classes.
*/
type FunctionClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) FunctionLike
}

/*
FunctionalClassLike defines the set of class constants, constructors and
functions that must be supported by all functional-class-like classes.
*/
type FunctionalClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		parameters ParametersLike,
		result ResultLike,
	) FunctionalLike
}

/*
FunctionalsClassLike defines the set of class constants, constructors and
functions that must be supported by all functionals-class-like classes.
*/
type FunctionalsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[FunctionalLike]) FunctionalsLike
}

/*
FunctionsClassLike defines the set of class constants, constructors and
functions that must be supported by all functions-class-like classes.
*/
type FunctionsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[FunctionLike]) FunctionsLike
}

/*
GeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all generator-class-like classes.
*/
type GeneratorClassLike interface {
	// Constructors
	Make() GeneratorLike
}

/*
PackageClassLike defines the set of class constants, constructors and functions
that must be supported by all package-class-like classes.
*/
type PackageClassLike interface {
	// Constructors
	MakeWithAttributes(
		notice NoticeLike,
		header HeaderLike,
		imports ImportsLike,
		types TypesLike,
		interfaces InterfacesLike,
	) PackageLike
}

/*
HeaderClassLike defines the set of class constants, constructors and functions
that must be supported by all header-class-like classes.
*/
type HeaderClassLike interface {
	// Constructors
	MakeWithAttributes(comment string, identifier string) HeaderLike
}

/*
ImportsClassLike defines the set of class constants, constructors and functions
that must be supported by all imports-class-like classes.
*/
type ImportsClassLike interface {
	// Constructors
	MakeWithAttributes(modules ModulesLike) ImportsLike
}

/*
InstanceClassLike defines the set of class constants, constructors and functions
that must be supported by all instance-class-like classes.
*/
type InstanceClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		attributes AttributesLike,
		abstractions AbstractionsLike,
		methods MethodsLike,
	) InstanceLike
}

/*
InstancesClassLike defines the set of class constants, constructors and
functions that must be supported by all instances-class-like classes.
*/
type InstancesClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[InstanceLike]) InstancesLike
}

/*
InterfacesClassLike defines the set of class constants, constructors and
functions that must be supported by all interfaces-class-like classes.
*/
type InterfacesClassLike interface {
	// Constructors
	MakeWithAttributes(
		aspects AspectsLike,
		classes ClassesLike,
		instances InstancesLike,
	) InterfacesLike
}

/*
MethodClassLike defines the set of class constants, constructors and functions
that must be supported by all method-class-like classes.
*/
type MethodClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) MethodLike
}

/*
MethodsClassLike defines the set of class constants, constructors and functions
that must be supported by all methods-class-like classes.
*/
type MethodsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[MethodLike]) MethodsLike
}

/*
ModuleClassLike defines the set of class constants, constructors and
functions that must be supported by all module-class-like classes.
*/
type ModuleClassLike interface {
	// Constructors
	MakeWithAttributes(identifier string, repository string) ModuleLike
}

/*
ModulesClassLike defines the set of class constants, constructors and functions
that must be supported by all modules-class-like classes.
*/
type ModulesClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[ModuleLike]) ModulesLike
}

/*
NoticeClassLike defines the set of class constants, constructors and
functions that must be supported by all notice-class-like classes.
*/
type NoticeClassLike interface {
	// Constructors
	MakeWithAttributes(comment string) NoticeLike
}

/*
ParameterClassLike defines the set of class constants, constructors and
functions that must be supported by all parameter-class-like classes.
*/
type ParameterClassLike interface {
	// Constructors
	MakeWithAttributes(identifier string, abstraction AbstractionLike) ParameterLike
}

/*
ParametersClassLike defines the set of class constants, constructors and
functions that must be supported by all parameters-class-like classes.
*/
type ParametersClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[ParameterLike]) ParametersLike
}

/*
ParserClassLike defines the set of class constants, constructors and functions
that must be supported by all parser-class-like classes.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
PrefixClassLike defines the set of class constants, constructors and
functions that must be supported by all prefix-class-like classes.
*/
type PrefixClassLike interface {
	// Constructors
	MakeWithAttributes(identifier string, type_ PrefixType) PrefixLike
}

/*
ResultClassLike defines the set of class constants, constructors and functions
that must be supported by all result-class-like classes.
*/
type ResultClassLike interface {
	// Constructors
	MakeWithAbstraction(abstraction AbstractionLike) ResultLike
	MakeWithParameters(parameters ParametersLike) ResultLike
}

/*
ScannerClassLike defines the set of class constants, constructors and functions
that must be supported by all scanner-class-like classes.
*/
type ScannerClassLike interface {
	// Constructors
	MakeFromSource(source string, tokens col.QueueLike[TokenLike]) ScannerLike

	// Functions
	MatchToken(type_ TokenType, text string) col.ListLike[string]
}

/*
SpecializationClassLike defines the set of class constants, constructors and
functions that must be supported by all specialization-class-like classes.
*/
type SpecializationClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		enumeration EnumerationLike,
	) SpecializationLike
}

/*
SpecializationsClassLike defines the set of class constants, constructors and
functions that must be supported by all specializations-class-like classes.
*/
type SpecializationsClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[SpecializationLike]) SpecializationsLike
}

/*
TokenClassLike defines the set of class constants, constructors and functions
that must be supported by all token-class-like classes.
*/
type TokenClassLike interface {
	// Constructors
	MakeWithAttributes(
		line int,
		position int,
		type_ TokenType,
		value string,
	) TokenLike

	// Functions
	AsString(type_ TokenType) string
}

/*
TypesClassLike defines the set of class constants, constructors and functions
that must be supported by all types-class-like classes.
*/
type TypesClassLike interface {
	// Constructors
	MakeWithAttributes(specializations SpecializationsLike, functionals FunctionalsLike) TypesLike
}

/*
ValidatorClassLike defines the set of class constants, constructors and
functions that must be supported by all validator-class-like classes.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

/*
ValuesClassLike defines the set of class constants, constructors and functions
that must be supported by all values-class-like classes.
*/
type ValuesClassLike interface {
	// Constructors
	MakeWithAttributes(sequence col.Sequential[string], abstraction AbstractionLike) ValuesLike
}

// Instances

/*
AbstractionLike defines the set of abstractions and methods that must be
supported by all abstraction-like instances.
*/
type AbstractionLike interface {
	// Attributes
	GetArguments() ArgumentsLike
	GetIdentifier() string
	GetPrefix() PrefixLike
}

/*
AbstractionsLike defines the set of abstractions and methods that must be
supported by all abstractions-like instances.
*/
type AbstractionsLike interface {
	// Attributes
	GetSequence() col.Sequential[AbstractionLike]
}

/*
ArgumentsLike defines the set of abstractions and methods that must be supported
by all arguments-like instances.
*/
type ArgumentsLike interface {
	// Attributes
	GetSequence() col.Sequential[AbstractionLike]
}

/*
AspectLike defines the set of abstractions and methods that must be supported by
all aspect-like instances.
*/
type AspectLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetMethods() MethodsLike
}

/*
AspectsLike defines the set of abstractions and methods that must be supported
by all aspects-like instances.
*/
type AspectsLike interface {
	// Attributes
	GetSequence() col.Sequential[AspectLike]
}

/*
AttributeLike defines the set of abstractions and methods that must be supported
by all attribute-like instances.
*/
type AttributeLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetIdentifier() string
	GetParameter() ParameterLike
}

/*
AttributesLike defines the set of abstractions and methods that must be supported
by all attributes-like instances.
*/
type AttributesLike interface {
	// Attributes
	GetSequence() col.Sequential[AttributeLike]
}

/*
ClassLike defines the set of abstractions and methods that must be supported by
all class-like instances.
*/
type ClassLike interface {
	// Attributes
	GetConstants() ConstantsLike
	GetConstructors() ConstructorsLike
	GetDeclaration() DeclarationLike
	GetFunctions() FunctionsLike
}

/*
ClassesLike defines the set of abstractions and methods that must be supported
by all classes-like instances.
*/
type ClassesLike interface {
	// Attributes
	GetSequence() col.Sequential[ClassLike]
}

/*
ConstantLike defines the set of abstractions and methods that must be supported
by all constant-like instances.
*/
type ConstantLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetIdentifier() string
}

/*
ConstantsLike defines the set of abstractions and methods that must be supported
by all constants-like instances.
*/
type ConstantsLike interface {
	// Attributes
	GetSequence() col.Sequential[ConstantLike]
}

/*
ConstructorLike defines the set of abstractions and methods that must be
supported by all constructor-like instances.
*/
type ConstructorLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetIdentifier() string
	GetParameters() ParametersLike
}

/*
ConstructorsLike defines the set of abstractions and methods that must be
supported by all constructors-like instances.
*/
type ConstructorsLike interface {
	// Attributes
	GetSequence() col.Sequential[ConstructorLike]
}

/*
DeclarationLike defines the set of abstractions and methods that must be
supported by all declaration-like instances.
*/
type DeclarationLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
	GetParameters() ParametersLike
}

/*
EnumerationLike defines the set of abstractions and methods that must be
supported by all enumeration-like instances.
*/
type EnumerationLike interface {
	// Attributes
	GetValues() ValuesLike
}

/*
FormatterLike defines the set of abstractions and methods that must be
supported by all formatter-like instances.
*/
type FormatterLike interface {
	// Methods
	FormatAbstraction(abstraction AbstractionLike) string
	FormatArguments(arguments ArgumentsLike) string
	FormatPackage(package_ PackageLike) string
	FormatMethod(method MethodLike) string
	FormatParameter(parameter ParameterLike) string
	FormatParameterNames(parameters ParametersLike) string
	FormatParameters(parameters ParametersLike) string
	FormatResult(result ResultLike) string
}

/*
FunctionLike defines the set of abstractions and methods that must be supported
by all function-like instances.
*/
type FunctionLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
FunctionalLike defines the set of abstractions and methods that must be
supported by all functional-like instances.
*/
type FunctionalLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
FunctionalsLike defines the set of abstractions and methods that must be
supported by all functionals-like instances.
*/
type FunctionalsLike interface {
	// Attributes
	GetSequence() col.Sequential[FunctionalLike]
}

/*
FunctionsLike defines the set of abstractions and methods that must be supported
by all functions-like instances.
*/
type FunctionsLike interface {
	// Attributes
	GetSequence() col.Sequential[FunctionLike]
}

/*
GeneratorLike defines the set of abstractions and methods that must be
supported by all generator-like instances.
*/
type GeneratorLike interface {
	// Methods
	CreatePackage(directory string, copyright string)
	GeneratePackage(directory string)
}

/*
PackageLike defines the set of abstractions and methods that must be supported by
all package-like instances.
*/
type PackageLike interface {
	// Attributes
	GetNotice() NoticeLike
	GetHeader() HeaderLike
	GetImports() ImportsLike
	GetInterfaces() InterfacesLike
	GetTypes() TypesLike
}

/*
HeaderLike defines the set of abstractions and methods that must be supported by
all header-like instances.
*/
type HeaderLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
}

/*
ImportsLike defines the set of abstractions and methods that must be supported
by all imports-like instances.
*/
type ImportsLike interface {
	// Attributes
	GetModules() ModulesLike
}

/*
InstanceLike defines the set of abstractions and methods that must be supported
by all instance-like instances.
*/
type InstanceLike interface {
	// Attributes
	GetAbstractions() AbstractionsLike
	GetAttributes() AttributesLike
	GetDeclaration() DeclarationLike
	GetMethods() MethodsLike
}

/*
InstancesLike defines the set of abstractions and methods that must be supported
by all instances-like instances.
*/
type InstancesLike interface {
	// Attributes
	GetSequence() col.Sequential[InstanceLike]
}

/*
InterfacesLike defines the set of abstractions and methods that must be
supported by all interfaces-like instances.
*/
type InterfacesLike interface {
	// Attributes
	GetAspects() AspectsLike
	GetClasses() ClassesLike
	GetInstances() InstancesLike
}

/*
MethodLike defines the set of abstractions and methods that must be supported by
all method-like instances.
*/
type MethodLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
MethodsLike defines the set of abstractions and methods that must be supported
by all methods-like instances.
*/
type MethodsLike interface {
	// Attributes
	GetSequence() col.Sequential[MethodLike]
}

/*
ModuleLike defines the set of abstractions and methods that must be
supported by all module-like instances.
*/
type ModuleLike interface {
	// Attributes
	GetIdentifier() string
	GetRepository() string
}

/*
ModulesLike defines the set of abstractions and methods that must be supported
by all modules-like instances.
*/
type ModulesLike interface {
	// Attributes
	GetSequence() col.Sequential[ModuleLike]
}

/*
NoticeLike defines the set of abstractions and methods that must be supported
by all notice-like instances.
*/
type NoticeLike interface {
	// Attributes
	GetComment() string
}

/*
ParameterLike defines the set of abstractions and methods that must be supported
by all parameter-like instances.
*/
type ParameterLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetIdentifier() string
}

/*
ParametersLike defines the set of abstractions and methods that must be
supported by all parameters-like instances.
*/
type ParametersLike interface {
	// Attributes
	GetSequence() col.Sequential[ParameterLike]
}

/*
ParserLike defines the set of abstractions and methods that must be supported by
all parser-like instances.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) PackageLike
}

/*
PrefixLike defines the set of abstractions and methods that must be
supported by all prefix-like instances.
*/
type PrefixLike interface {
	// Attributes
	GetIdentifier() string
	GetType() PrefixType
}

/*
ResultLike defines the set of abstractions and methods that must be supported by
all result-like instances.
*/
type ResultLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetParameters() ParametersLike
}

/*
ScannerLike defines the set of abstractions and methods that must be supported
by all scanner-like instances.
*/
type ScannerLike interface {
}

/*
SpecializationLike defines the set of abstractions and methods that must be
supported by all specialization-like instances.
*/
type SpecializationLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetDeclaration() DeclarationLike
	GetEnumeration() EnumerationLike
}

/*
SpecializationsLike defines the set of abstractions and methods that must be
supported by all parameters-like instances.
*/
type SpecializationsLike interface {
	// Attributes
	GetSequence() col.Sequential[SpecializationLike]
}

/*
TokenLike defines the set of abstractions and methods that must be supported by
all token-like instances.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
TypesLike defines the set of abstractions and methods that must be supported by
all types-like instances.
*/
type TypesLike interface {
	// Attributes
	GetFunctionals() FunctionalsLike
	GetSpecializations() SpecializationsLike
}

/*
ValidatorLike defines the set of abstractions and methods that must be
supported by all validator-like instances.
*/
type ValidatorLike interface {
	// Methods
	ValidatePackage(package_ PackageLike)
}

/*
ValuesLike defines the set of abstractions and methods that must be supported by
all values-like instances.
*/
type ValuesLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetSequence() col.Sequential[string]
}
