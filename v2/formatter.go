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

import (
	reg "regexp"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// This class does not initialize any class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// This class does not define any class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	depth  int
	result sts.Builder
}

// Public

func (v *formatter_) FormatAbstraction(abstraction AbstractionLike) string {
	v.formatAbstraction(abstraction)
	return v.getResult()
}

func (v *formatter_) FormatArguments(arguments ArgumentsLike) string {
	v.formatArguments(arguments)
	return v.getResult()
}

func (v *formatter_) FormatGoPN(gopn GoPNLike) string {
	v.formatGoPN(gopn)
	return v.getResult()
}

func (v *formatter_) FormatMethod(method MethodLike) string {
	v.formatMethod(method)
	return v.getResult()
}

func (v *formatter_) FormatParameter(parameter ParameterLike) string {
	v.formatParameter(parameter)
	return v.getResult()
}

func (v *formatter_) FormatParameterNames(parameters ParametersLike) string {
	v.formatParameterNames(parameters)
	return v.getResult()
}

func (v *formatter_) FormatParameters(parameters ParametersLike) string {
	v.formatParameters(parameters)
	return v.getResult()
}

func (v *formatter_) FormatResult(result ResultLike) string {
	v.formatResult(result)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var separator = "\n"
	for level := 0; level < v.depth; level++ {
		separator += "\t"
	}
	v.appendString(separator)
}

func (v *formatter_) appendString(s string) {
	v.result.WriteString(s)
}

func (v *formatter_) fixComment(comment string) string {
	var matcher = reg.MustCompile("\n    ")
	comment = matcher.ReplaceAllString(comment, "\n\t")
	return comment
}

func (v *formatter_) formatAbstraction(abstraction AbstractionLike) {
	var prefix = abstraction.GetPrefix()
	if prefix != nil {
		v.formatPrefix(prefix)
	}
	var identifier = abstraction.GetIdentifier()
	v.appendString(identifier)
	var arguments = abstraction.GetArguments()
	if arguments != nil {
		v.appendString("[")
		v.formatArguments(arguments)
		v.appendString("]")
	}
}

func (v *formatter_) formatAbstractions(abstractions AbstractionsLike) {
	v.appendNewline()
	v.appendString("// Abstractions")
	var iterator = abstractions.GetSequence().GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.appendNewline()
		v.formatAbstraction(abstraction)
	}
}

func (v *formatter_) formatArguments(arguments ArgumentsLike) {
	var sequence = arguments.GetSequence()
	var size = sequence.GetSize()
	if size > 2 {
		v.depth++
		v.appendNewline()
	}
	var iterator = sequence.GetIterator()
	var abstraction = iterator.GetNext()
	v.formatAbstraction(abstraction)
	for iterator.HasNext() {
		abstraction = iterator.GetNext()
		v.appendString(",")
		if size > 2 {
			v.appendNewline()
		} else {
			v.appendString(" ")
		}
		v.formatAbstraction(abstraction)
	}
	if size > 2 {
		v.appendString(",")
		v.depth--
		v.appendNewline()
	}
}

func (v *formatter_) formatAspect(aspect AspectLike) {
	var declaration = aspect.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth++
	var methods = aspect.GetMethods()
	if methods != nil {
		v.formatMethods(methods)
	}
	v.depth--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatAspects(aspects AspectsLike) {
	v.appendNewline()
	v.appendString("// Aspects")
	v.appendNewline()
	var iterator = aspects.GetSequence().GetIterator()
	for iterator.HasNext() {
		var aspect = iterator.GetNext()
		v.formatAspect(aspect)
		v.appendNewline()
	}
}

func (v *formatter_) formatAttribute(attribute AttributeLike) {
	var identifier = attribute.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameter = attribute.GetParameter()
	if parameter != nil {
		v.formatParameter(parameter)
	}
	v.appendString(")")
	var abstraction = attribute.GetAbstraction()
	if abstraction != nil {
		v.appendString(" ")
		v.formatAbstraction(abstraction)
	}
}

func (v *formatter_) formatAttributes(attributes AttributesLike) {
	v.appendNewline()
	v.appendString("// Attributes")
	var iterator = attributes.GetSequence().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		v.appendNewline()
		v.formatAttribute(attribute)
	}
}

func (v *formatter_) formatClass(class ClassLike) {
	var declaration = class.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth++
	var hasContent bool
	var constants = class.GetConstants()
	if constants != nil {
		v.formatConstants(constants)
		hasContent = true
	}
	var constructors = class.GetConstructors()
	if constructors != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatConstructors(constructors)
		hasContent = true
	}
	var functions = class.GetFunctions()
	if functions != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatFunctions(functions)
	}
	v.depth--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatClasses(classes ClassesLike) {
	v.appendNewline()
	v.appendString("// Classes")
	v.appendNewline()
	var iterator = classes.GetSequence().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		v.formatClass(class)
		v.appendNewline()
	}
}

func (v *formatter_) formatConstant(constant ConstantLike) {
	var identifier = constant.GetIdentifier()
	v.appendString(identifier)
	v.appendString("() ")
	var abstraction = constant.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatConstants(constants ConstantsLike) {
	v.appendNewline()
	v.appendString("// Constants")
	var iterator = constants.GetSequence().GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.appendNewline()
		v.formatConstant(constant)
	}
}

func (v *formatter_) formatConstructor(constructor ConstructorLike) {
	var identifier = constructor.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameters = constructor.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(") ")
	var abstraction = constructor.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatConstructors(constructors ConstructorsLike) {
	v.appendNewline()
	v.appendString("// Constructors")
	var iterator = constructors.GetSequence().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.appendNewline()
		v.formatConstructor(constructor)
	}
}

func (v *formatter_) formatCopyright(copyright CopyrightLike) {
	var comment = copyright.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
}

func (v *formatter_) formatDeclaration(declaration DeclarationLike) {
	v.appendNewline()
	var comment = declaration.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
	v.appendString("type ")
	var identifier = declaration.GetIdentifier()
	v.appendString(identifier)
	var parameters = declaration.GetParameters()
	if parameters != nil {
		v.appendString("[")
		v.formatParameters(parameters)
		v.appendString("]")
	}
}

func (v *formatter_) formatEnumeration(enumeration EnumerationLike) {
	v.appendNewline()
	v.appendString("const (")
	v.depth++
	v.appendNewline()
	var values = enumeration.GetValues()
	v.formatValues(values)
	v.depth--
	v.appendNewline()
	v.appendString(")")
}

func (v *formatter_) formatFunction(function FunctionLike) {
	var identifier = function.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameters = function.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(") ")
	var result = function.GetResult()
	v.formatResult(result)
}

func (v *formatter_) formatFunctions(functions FunctionsLike) {
	v.appendNewline()
	v.appendString("// Functions")
	var iterator = functions.GetSequence().GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.appendNewline()
		v.formatFunction(function)
	}
}

func (v *formatter_) formatFunctional(functional FunctionalLike) {
	var declaration = functional.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" func(")
	var parameters = functional.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(") ")
	var result = functional.GetResult()
	v.formatResult(result)
}

func (v *formatter_) formatFunctionals(functionals FunctionalsLike) {
	v.appendNewline()
	v.appendString("// Functionals")
	v.appendNewline()
	var iterator = functionals.GetSequence().GetIterator()
	for iterator.HasNext() {
		var functional = iterator.GetNext()
		v.formatFunctional(functional)
		v.appendNewline()
	}
}

func (v *formatter_) formatGoPN(gopn GoPNLike) {
	var copyright = gopn.GetCopyright()
	v.formatCopyright(copyright)
	var header = gopn.GetHeader()
	v.formatHeader(header)
	v.appendNewline()
	var imports = gopn.GetImports()
	if imports != nil {
		v.formatImports(imports)
	}
	var types = gopn.GetTypes()
	if types != nil {
		v.formatTypes(types)
	}
	var interfaces = gopn.GetInterfaces()
	if interfaces != nil {
		v.formatInterfaces(interfaces)
	}
}

func (v *formatter_) formatHeader(header HeaderLike) {
	var comment = header.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
	v.appendString("package ")
	var identifier = header.GetIdentifier()
	v.appendString(identifier)
}

func (v *formatter_) formatImports(imports ImportsLike) {
	v.appendNewline()
	v.appendString("import (")
	var modules = imports.GetModules()
	if modules != nil {
		v.formatModules(modules)
	}
	v.appendString(")")
	v.appendNewline()
}

func (v *formatter_) formatInstance(instance InstanceLike) {
	var declaration = instance.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth++
	var hasContent bool
	var attributes = instance.GetAttributes()
	if attributes != nil {
		v.formatAttributes(attributes)
		hasContent = true
	}
	var abstractions = instance.GetAbstractions()
	if abstractions != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatAbstractions(abstractions)
		hasContent = true
	}
	var methods = instance.GetMethods()
	if methods != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatMethods(methods)
	}
	v.depth--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatInstances(instances InstancesLike) {
	v.appendNewline()
	v.appendString("// Instances")
	v.appendNewline()
	var iterator = instances.GetSequence().GetIterator()
	for iterator.HasNext() {
		var instance = iterator.GetNext()
		v.formatInstance(instance)
		v.appendNewline()
	}
}

func (v *formatter_) formatInterfaces(interfaces InterfacesLike) {
	v.appendNewline()
	v.appendString("// INTERFACES")
	v.appendNewline()
	var aspects = interfaces.GetAspects()
	if aspects != nil {
		v.formatAspects(aspects)
	}
	var classes = interfaces.GetClasses()
	if classes != nil {
		v.formatClasses(classes)
	}
	var instances = interfaces.GetInstances()
	if instances != nil {
		v.formatInstances(instances)
	}
}

func (v *formatter_) formatMethod(method MethodLike) {
	var identifier = method.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameters = method.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(")")
	var result = method.GetResult()
	if result != nil {
		v.appendString(" ")
		v.formatResult(result)
	}
}

func (v *formatter_) formatMethods(methods MethodsLike) {
	v.appendNewline()
	v.appendString("// Methods")
	var iterator = methods.GetSequence().GetIterator()
	for iterator.HasNext() {
		var method = iterator.GetNext()
		v.appendNewline()
		v.formatMethod(method)
	}
}

func (v *formatter_) formatModule(module ModuleLike) {
	var identifier = module.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var repository = module.GetRepository()
	v.appendString(repository)
}

func (v *formatter_) formatModules(modules ModulesLike) {
	v.depth++
	v.appendNewline()
	var sequence = modules.GetSequence()
	var iterator = sequence.GetIterator()
	var module = iterator.GetNext()
	v.formatModule(module)
	for iterator.HasNext() {
		module = iterator.GetNext()
		v.appendNewline()
		v.formatModule(module)
	}
	v.depth--
	v.appendNewline()
}

func (v *formatter_) formatParameter(parameter ParameterLike) {
	var identifier = parameter.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var abstraction = parameter.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatParameterName(parameter ParameterLike) {
	var identifier = parameter.GetIdentifier()
	v.appendString(identifier)
}

func (v *formatter_) formatParameterNames(parameters ParametersLike) {
	var sequence = parameters.GetSequence()
	var size = sequence.GetSize()
	if size > 2 {
		v.depth++
		v.appendNewline()
	}
	var iterator = sequence.GetIterator()
	var parameter = iterator.GetNext()
	v.formatParameterName(parameter)
	for iterator.HasNext() {
		parameter = iterator.GetNext()
		v.appendString(",")
		if size > 2 {
			v.appendNewline()
		} else {
			v.appendString(" ")
		}
		v.formatParameterName(parameter)
	}
	if size > 2 {
		v.appendString(",")
		v.depth--
		v.appendNewline()
	}
}

func (v *formatter_) formatParameters(parameters ParametersLike) {
	var sequence = parameters.GetSequence()
	var size = sequence.GetSize()
	if size > 2 {
		v.depth++
		v.appendNewline()
	}
	var iterator = sequence.GetIterator()
	var parameter = iterator.GetNext()
	v.formatParameter(parameter)
	for iterator.HasNext() {
		parameter = iterator.GetNext()
		v.appendString(",")
		if size > 2 {
			v.appendNewline()
		} else {
			v.appendString(" ")
		}
		v.formatParameter(parameter)
	}
	if size > 2 {
		v.appendString(",")
		v.depth--
		v.appendNewline()
	}
}

func (v *formatter_) formatPrefix(prefix PrefixLike) {
	var identifier = prefix.GetIdentifier()
	switch prefix.GetType() {
	case AliasPrefix:
		v.appendString(identifier)
		v.appendString(".")
	case ArrayPrefix:
		v.appendString("[]")
	case ChannelPrefix:
		v.appendString("chan ")
	case MapPrefix:
		v.appendString("map[")
		v.appendString(identifier)
		v.appendString("]")
	}
}

func (v *formatter_) formatResult(result ResultLike) {
	var abstraction = result.GetAbstraction()
	if abstraction != nil {
		v.formatAbstraction(abstraction)
	} else {
		v.appendString("(")
		var parameters = result.GetParameters()
		v.formatParameters(parameters)
		v.appendString(")")
	}
}

func (v *formatter_) formatSpecialization(specialization SpecializationLike) {
	var declaration = specialization.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" ")
	var abstraction = specialization.GetAbstraction()
	v.formatAbstraction(abstraction)
	var enumeration = specialization.GetEnumeration()
	if enumeration != nil {
		v.appendNewline()
		v.formatEnumeration(enumeration)
	}
}

func (v *formatter_) formatSpecializations(specializations SpecializationsLike) {
	v.appendNewline()
	v.appendString("// Specializations")
	v.appendNewline()
	var iterator = specializations.GetSequence().GetIterator()
	for iterator.HasNext() {
		var specialization = iterator.GetNext()
		v.formatSpecialization(specialization)
		v.appendNewline()
	}
}

func (v *formatter_) formatTypes(types TypesLike) {
	v.appendNewline()
	v.appendString("// TYPES")
	v.appendNewline()
	var specializations = types.GetSpecializations()
	if specializations != nil {
		v.formatSpecializations(specializations)
	}
	var functionals = types.GetFunctionals()
	if functionals != nil {
		v.formatFunctionals(functionals)
	}
}

func (v *formatter_) formatValues(values ValuesLike) {
	var iterator = values.GetSequence().GetIterator()
	var identifier = iterator.GetNext()
	v.appendString(identifier)
	v.appendString(" ")
	var abstraction = values.GetAbstraction()
	v.formatAbstraction(abstraction)
	v.appendString(" = iota")
	for iterator.HasNext() {
		identifier = iterator.GetNext()
		v.appendNewline()
		v.appendString(identifier)
	}
}

func (v *formatter_) getResult() string {
	var result = v.result.String()
	v.result.Reset()
	return result
}
