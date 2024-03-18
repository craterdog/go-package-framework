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
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
	osx "os"
	sts "strings"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// This class does not initialize any class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// This class does not define any class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		// This class does not initialize any instance attributes.
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	// This class does not define any instance attributes.
}

// Public

func (v *generator_) GeneratePackage(directory string, copyright string) {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var package_ = v.parseModel(directory, copyright)
	if package_ == nil {
		return
	}
	v.generateModel(directory, package_)
	v.generateClasses(directory, package_)
}

// Private

func (v *generator_) createDirectory(directory string) {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) extractConstructorAttributes(
	class ClassLike,
	catalog col.CatalogLike[string, string],
) {
	var constructors = class.GetConstructors()
	if constructors != nil {
		var iterator = constructors.GetSequence().GetIterator()
		for iterator.HasNext() {
			var constructor = iterator.GetNext()
			var methodName = constructor.GetIdentifier()
			if sts.HasPrefix(methodName, "MakeWith") {
				var parameters = constructor.GetParameters()
				v.extractParameterAttributes(parameters, catalog)
			}
		}
	}
}

func (v *generator_) extractInstanceAttributes(
	instance InstanceLike,
	catalog col.CatalogLike[string, string],
) {
	var attributeName string
	var attributeType string
	var formatter = Formatter().Make()
	var attributes = instance.GetAttributes()
	if attributes != nil {
		var iterator = attributes.GetSequence().GetIterator()
		for iterator.HasNext() {
			var attribute = iterator.GetNext()
			var identifier = attribute.GetIdentifier()
			var abstraction AbstractionLike
			switch {
			case sts.HasPrefix(identifier, "Get"):
				attributeName = sts.TrimPrefix(identifier, "Get")
				abstraction = attribute.GetAbstraction()
			case sts.HasPrefix(identifier, "Is"):
				attributeName = sts.TrimPrefix(identifier, "Is")
				abstraction = attribute.GetAbstraction()
			case sts.HasPrefix(identifier, "Was"):
				attributeName = sts.TrimPrefix(identifier, "Was")
				abstraction = attribute.GetAbstraction()
			case sts.HasPrefix(identifier, "Has"):
				attributeName = sts.TrimPrefix(identifier, "Has")
				abstraction = attribute.GetAbstraction()
			default:
				if attributeName == v.makePrivate(sts.TrimPrefix(identifier, "Set")) {
					// This attribute was already added.
					continue
				}
				attributeName = sts.TrimPrefix(identifier, "Set")
				var parameter = attribute.GetParameter()
				abstraction = parameter.GetAbstraction()
			}
			attributeName = v.makePrivate(attributeName)
			attributeType = formatter.FormatAbstraction(abstraction)
			catalog.SetValue(attributeName, attributeType)
		}
	}
}

func (v *generator_) extractParameterAttributes(
	parameters ParametersLike,
	catalog col.CatalogLike[string, string],
) {
	var formatter = Formatter().Make()
	var iterator = parameters.GetSequence().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var attributeName = attribute.GetIdentifier()
		attributeName = sts.TrimSuffix(attributeName, "_")
		var abstraction = attribute.GetAbstraction()
		var attributeType = formatter.FormatAbstraction(abstraction)
		catalog.SetValue(attributeName, attributeType)
	}
}

func (v *generator_) generateAbstractionMethods(
	aspect AspectLike,
	abstraction AbstractionLike,
) string {
	var formatter = Formatter().Make()
	var aspectDeclaration = aspect.GetDeclaration()
	var genericTypes = aspectDeclaration.GetParameters()
	var concreteTypes = abstraction.GetArguments()
	var abstractionMethods string
	var aspectMethods = aspect.GetMethods()
	if aspectMethods == nil {
		return abstractionMethods
	}
	var iterator = aspectMethods.GetSequence().GetIterator()
	for iterator.HasNext() {
		var aspectMethod = iterator.GetNext()
		var methodName = aspectMethod.GetIdentifier()
		var methodParameters = aspectMethod.GetParameters()
		var parameters string
		if methodParameters != nil {
			if genericTypes != nil {
				// Replace the generic type names from the aspect definition
				// with the actual types defined in the instance interface.
				methodParameters = v.replaceParameterTypes(
					genericTypes,
					concreteTypes,
					methodParameters,
				)
			}
			parameters = formatter.FormatParameters(methodParameters)
		}
		var resultType string
		var body = methodBodyTemplate_
		var methodResult = aspectMethod.GetResult()
		if methodResult != nil {
			if genericTypes != nil {
				// Replace the generic type names from the aspect definition
				// with the actual types defined in the instance interface.
				methodResult = v.replaceResultTypes(
					genericTypes,
					concreteTypes,
					methodResult,
				)
			}
			resultType = " " + formatter.FormatResult(methodResult)
			if methodResult.GetAbstraction() != nil {
				body = resultBodyTemplate_
			} else {
				body = returnBodyTemplate_
			}
		}
		var method = instanceMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		abstractionMethods += method + "\n"
	}
	return abstractionMethods
}

func (v *generator_) generateAbstractions(
	package_ PackageLike,
	instanceInterface InstanceLike,
) string {
	var formatter = Formatter().Make()
	var result string
	var abstractions = instanceInterface.GetAbstractions()
	if abstractions == nil {
		return result
	}
	var iterator = abstractions.GetSequence().GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		var prefix = abstraction.GetPrefix()
		var identifier = abstraction.GetIdentifier()
		var aspectName = formatter.FormatAbstraction(abstraction)
		var methods string
		if prefix == nil {
			// We only know the method signatures for the local aspects.
			var aspect = v.retrieveAspect(package_, identifier)
			methods = v.generateAbstractionMethods(aspect, abstraction)
		}
		var instanceAspect = instanceAspectTemplate_
		instanceAspect = sts.ReplaceAll(instanceAspect, "<AspectName>", aspectName)
		instanceAspect = sts.ReplaceAll(instanceAspect, "<Methods>", methods)
		result += instanceAspect
	}
	return result
}

func (v *generator_) generateAttributeAssignments(
	classInterface ClassLike,
	constructor ConstructorLike,
) string {
	var assignments string
	var identifier = constructor.GetIdentifier()
	if !sts.HasPrefix(identifier, "MakeWith") {
		return assignments
	}
	var parameters = constructor.GetParameters()
	var iterator = parameters.GetSequence().GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var parameterName = parameter.GetIdentifier()
		var attributeName = sts.TrimSuffix(parameterName, "_")
		var assignment = attributeAssignmentTemplate_
		assignment = sts.ReplaceAll(assignment, "<AttributeName>", attributeName)
		assignment = sts.ReplaceAll(assignment, "<ParameterName>", parameterName)
		assignments += assignment
	}
	assignments += "\n\t"
	return assignments
}

func (v *generator_) generateAttributeMethods(instanceInterface InstanceLike) string {
	var formatter = Formatter().Make()
	var methods string
	var instanceAttributes = instanceInterface.GetAttributes()
	if instanceAttributes == nil {
		return methods
	}
	var iterator = instanceAttributes.GetSequence().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var methodName = attribute.GetIdentifier()
		var attributeName string
		var body string

		var parameter string
		var attributeParameter = attribute.GetParameter()
		var parameterName string
		if attributeParameter != nil {
			attributeName = sts.TrimPrefix(methodName, "Set")
			parameterName = attributeParameter.GetIdentifier()
			parameter = formatter.FormatParameter(attributeParameter)
			body = setterBodyTemplate_
		}

		var resultType string
		var abstraction = attribute.GetAbstraction()
		if abstraction != nil {
			switch {
			case sts.HasPrefix(methodName, "Get"):
				attributeName = sts.TrimPrefix(methodName, "Get")
			case sts.HasPrefix(methodName, "Is"):
				attributeName = sts.TrimPrefix(methodName, "Is")
			case sts.HasPrefix(methodName, "Was"):
				attributeName = sts.TrimPrefix(methodName, "Was")
			case sts.HasPrefix(methodName, "Has"):
				attributeName = sts.TrimPrefix(methodName, "Has")
			}
			resultType = " " + formatter.FormatAbstraction(abstraction)
			body = getterBodyTemplate_
		}

		attributeName = v.makePrivate(attributeName)
		body = sts.ReplaceAll(body, "<AttributeName>", attributeName)
		body = sts.ReplaceAll(body, "<ParameterName>", parameterName)
		var method = instanceMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameter)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateClass(
	directory string,
	package_ PackageLike,
	classInterface ClassLike,
	instanceInterface InstanceLike,
) {
	var class = classTemplate_

	var notice = package_.GetNotice().GetComment()
	class = sts.ReplaceAll(class, "<Notice>", notice)

	var header = v.generateHeader(package_)
	class = sts.ReplaceAll(class, "<Header>", header)

	var classAccess = v.generateClassAccess(classInterface)
	class = sts.ReplaceAll(class, "<Access>", classAccess)

	var classMethods = v.generateClassMethods(classInterface)
	class = sts.ReplaceAll(class, "<Class>", classMethods)

	var instanceMethods = v.generateInstanceMethods(
		package_,
		classInterface,
		instanceInterface,
	)
	class = sts.ReplaceAll(class, "<Instance>", instanceMethods)

	var classDeclaration = classInterface.GetDeclaration()
	var classIdentifier = classDeclaration.GetIdentifier()
	var className = sts.TrimSuffix(classIdentifier, "ClassLike")
	class = sts.ReplaceAll(class, "<ClassName>", className)
	class = sts.ReplaceAll(class, "<TargetName>", v.makePrivate(className))

	var parameters string
	var arguments string
	var classParameters = classDeclaration.GetParameters()
	if classParameters != nil {
		var formatter = Formatter().Make()
		parameters = "[" + formatter.FormatParameters(classParameters) + "]"
		arguments = "[" + formatter.FormatParameterNames(classParameters) + "]"
	}
	class = sts.ReplaceAll(class, "[<Parameters>]", parameters)
	class = sts.ReplaceAll(class, "[<Arguments>]", arguments)

	var imports = v.generateImports(package_, class)
	class = sts.ReplaceAll(class, "<Imports>", imports)

	var fileName = sts.ToLower(className)
	var classFile = directory + fileName + ".go"
	v.outputClass(classFile, class)
}

func (v *generator_) generateClasses(directory string, package_ PackageLike) {
	var interfaces = package_.GetInterfaces()
	if interfaces == nil {
		return
	}
	var classes = interfaces.GetClasses()
	if classes == nil {
		return
	}
	var classIterator = classes.GetSequence().GetIterator()
	var instances = interfaces.GetInstances()
	var instanceIterator = instances.GetSequence().GetIterator()
	for classIterator.HasNext() {
		var classInterface = classIterator.GetNext()
		var instanceInterface = instanceIterator.GetNext()
		v.generateClass(directory, package_, classInterface, instanceInterface)
	}
}

func (v *generator_) generateClassAccess(classInterface ClassLike) string {
	var declaration = classInterface.GetDeclaration()
	var parameters = declaration.GetParameters()
	var reference = classReferenceTemplate_
	var function = classFunctionTemplate_
	if parameters != nil {
		reference = genericReferenceTemplate_
		function = genericFunctionTemplate_
	}
	var access = classAccessTemplate_
	access = sts.ReplaceAll(access, "<Reference>", reference)
	access = sts.ReplaceAll(access, "<Function>", function)
	return access + "\n"
}

func (v *generator_) generateClassConstants(classInterface ClassLike) string {
	var formatter = Formatter().Make()
	var constants string
	var classConstants = classInterface.GetConstants()
	if classConstants == nil {
		constants = "\n\t// TBA - Add private class constants.\n"
		return constants
	}
	var iterator = classConstants.GetSequence().GetIterator()
	for iterator.HasNext() {
		var classConstant = iterator.GetNext()
		var constantIdentifier = classConstant.GetIdentifier()
		var constantAbstraction = classConstant.GetAbstraction()
		var constantName = v.makePrivate(constantIdentifier)
		var constantType = formatter.FormatAbstraction(constantAbstraction)
		var constant = classConstantTemplate_
		constant = sts.ReplaceAll(constant, "<ConstantName>", constantName)
		constant = sts.ReplaceAll(constant, "<ConstantType>", constantType)
		constants += constant
	}
	constants += "\n"
	return constants
}

func (v *generator_) generateClassMethods(classInterface ClassLike) string {
	var methods = classMethodsTemplate_
	var target = v.generateClassTarget(classInterface)
	methods = sts.ReplaceAll(methods, "<Target>", target)
	var constantMethods = v.generateConstantMethods(classInterface)
	methods = sts.ReplaceAll(methods, "<Constants>", constantMethods)
	var constructorMethods = v.generateConstructorMethods(classInterface)
	methods = sts.ReplaceAll(methods, "<Constructors>", constructorMethods)
	var functionMethods = v.generateFunctionMethods(classInterface)
	methods = sts.ReplaceAll(methods, "<Functions>", functionMethods)
	return methods
}

func (v *generator_) generateClassTarget(classInterface ClassLike) string {
	var target = classTargetTemplate_
	var constants = v.generateClassConstants(classInterface)
	target = sts.ReplaceAll(target, "<Constants>", constants) + "\n"
	return target
}

func (v *generator_) generateConstantMethods(classInterface ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var classConstants = classInterface.GetConstants()
	if classConstants == nil {
		return methods
	}
	var iterator = classConstants.GetSequence().GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		var methodName = constant.GetIdentifier()
		var constantName = v.makePrivate(methodName)
		var abstraction = constant.GetAbstraction()
		var resultType = " " + formatter.FormatAbstraction(abstraction)
		var body = constantBodyTemplate_
		body = sts.ReplaceAll(body, "<ConstantName>", constantName)
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", "")
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateConstructorMethods(classInterface ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var iterator = classInterface.GetConstructors().GetSequence().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var methodName = constructor.GetIdentifier()
		var ConstructorParameters = constructor.GetParameters()
		var parameters string
		if ConstructorParameters != nil {
			parameters = formatter.FormatParameters(ConstructorParameters)
		}
		var abstraction = constructor.GetAbstraction()
		var resultType = " " + formatter.FormatAbstraction(abstraction)
		var assignments = v.generateAttributeAssignments(classInterface, constructor)
		var body = constructorBodyTemplate_
		body = sts.ReplaceAll(body, "<Assignments>", assignments)
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateFunctionMethods(classInterface ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var classFunctions = classInterface.GetFunctions()
	if classFunctions == nil {
		return methods
	}
	var iterator = classFunctions.GetSequence().GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		var identifier = function.GetIdentifier()
		var functionParameters = function.GetParameters()
		var parameters string
		if functionParameters != nil {
			parameters = formatter.FormatParameters(functionParameters)
		}
		var result = function.GetResult()
		var resultType = " " + formatter.FormatResult(result)
		var body = functionBodyTemplate_
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", identifier)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateHeader(package_ PackageLike) string {
	var packageName = package_.GetHeader().GetIdentifier()
	var header = headerTemplate_
	header = sts.ReplaceAll(header, "<PackageName>", packageName) + "\n"
	return header
}

func (v *generator_) generateImports(package_ PackageLike, class string) string {
	var imports string
	var modules string
	var packageImports = package_.GetImports()
	if packageImports != nil {
		var packageModules = packageImports.GetModules()
		if packageModules != nil {
			var iterator = packageModules.GetSequence().GetIterator()
			for iterator.HasNext() {
				var packageModule = iterator.GetNext()
				var identifier = packageModule.GetIdentifier()
				var repository = packageModule.GetRepository()
				if sts.Contains(class, identifier+".") {
					modules += "\n\t" + identifier + " " + repository
				}
			}
		}
	}
	if sts.Contains(class, "syn.") {
		modules += "\n\tfmt \"fmt\""
		modules += "\n\tsyn \"sync\""
	}
	if len(modules) > 0 {
		modules += "\n"
	}
	imports = importsTemplate_
	imports = sts.ReplaceAll(imports, "<Modules>", modules) + "\n"
	return imports
}

func (v *generator_) generateInstanceAttributes(
	classInterface ClassLike,
	instanceInterface InstanceLike,
) string {
	var attributes string
	var catalog = col.Catalog[string, string]().Make()
	v.extractInstanceAttributes(instanceInterface, catalog)
	v.extractConstructorAttributes(classInterface, catalog)
	if catalog.IsEmpty() {
		attributes = "\n\t// TBA - Add private instance attributes.\n"
		return attributes
	}
	var iterator = catalog.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var attributeName = association.GetKey()
		var attributeType = association.GetValue()
		var attribute = instanceAttributeTemplate_
		attribute = sts.ReplaceAll(attribute, "<AttributeName>", attributeName)
		attribute = sts.ReplaceAll(attribute, "<AttributeType>", attributeType)
		attributes += attribute
	}
	attributes += "\n"
	return attributes
}

func (v *generator_) generateInstanceMethods(
	package_ PackageLike,
	classInterface ClassLike,
	instanceInterface InstanceLike,
) string {
	var instanceMethods = instanceMethodsTemplate_
	var target = v.generateInstanceTarget(classInterface, instanceInterface)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Target>", target)
	var attributes = v.generateAttributeMethods(instanceInterface)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Attributes>", attributes)
	var abstractions = v.generateAbstractions(package_, instanceInterface)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Abstractions>", abstractions)
	var methods = v.generatePublicMethods(instanceInterface)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Methods>", methods)
	return instanceMethods
}

func (v *generator_) generateInstanceTarget(
	classInterface ClassLike,
	instanceInterface InstanceLike,
) string {
	var target = instanceTargetTemplate_
	var attributes = v.generateInstanceAttributes(classInterface, instanceInterface)
	target = sts.ReplaceAll(target, "<Attributes>", attributes) + "\n"
	return target
}

func (v *generator_) generateModel(directory string, package_ PackageLike) {
	var formatter = Formatter().Make()
	var source = formatter.FormatPackage(package_)
	var bytes = []byte(source)
	var modelFile = directory + "Model.go"
	var err = osx.WriteFile(modelFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) generatePublicMethods(instanceInterface InstanceLike) string {
	var formatter = Formatter().Make()
	var publicMethods string
	var instanceMethods = instanceInterface.GetMethods()
	if instanceMethods == nil {
		return publicMethods
	}
	var iterator = instanceMethods.GetSequence().GetIterator()
	for iterator.HasNext() {
		var publicMethod = iterator.GetNext()
		var methodName = publicMethod.GetIdentifier()
		var methodParameters = publicMethod.GetParameters()
		var parameters string
		if methodParameters != nil {
			parameters = formatter.FormatParameters(methodParameters)
		}
		var body = methodBodyTemplate_
		var result = publicMethod.GetResult()
		var resultType string
		if result != nil {
			if result.GetAbstraction() != nil {
				body = resultBodyTemplate_
			} else {
				body = returnBodyTemplate_
			}
			resultType = " " + formatter.FormatResult(result)
		}
		var method = instanceMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		publicMethods += method + "\n"
	}
	return publicMethods
}

func (v *generator_) makePrivate(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToLower(runes[0])
	return string(runes)
}

func (v *generator_) outputClass(classFile, class string) {
	var _, err = osx.ReadFile(classFile)
	if err == nil {
		// Don't overwrite an existing class file.
		fmt.Printf(
			"The class file %q already exists, leaving it alone.\n",
			classFile,
		)
		return
	}
	err = osx.WriteFile(classFile, []byte(class), 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) parseModel(directory string, copyright string) PackageLike {
	v.createDirectory(directory)
	var modelFile = directory + "Model.go"
	var bytes, err = osx.ReadFile(modelFile)
	if err != nil {
		// The file does not yet exist so create one.
		fmt.Printf(
			"The model file %q does not exist, creating a template for it.\n",
			modelFile,
		)
		// Center the copyright string.
		var maximum = 78
		var length = len(copyright)
		if length > maximum {
			copyright = copyright[:maximum]
			length = maximum
		}
		var padding = (maximum - length) / 2
		for range padding {
			copyright = " " + copyright + " "
		}
		if len(copyright) < maximum {
			copyright = " " + copyright
		}
		copyright = "." + copyright + "."

		// Create a new model template.
		var template = sts.ReplaceAll(modelTemplate_, "<Copyright>", copyright)
		bytes = []byte(template[1:]) // Remove leading "\n".
		err = osx.WriteFile(modelFile, bytes, 0644)
		if err != nil {
			panic(err)
		}
		return nil
	}
	var source = string(bytes)
	var parser = Parser().Make()
	var package_ = parser.ParseSource(source)
	return package_
}

func (v *generator_) replaceGenericType(
	genericTypes ParametersLike,
	concreteTypes ArgumentsLike,
	abstraction AbstractionLike,
) AbstractionLike {
	var formatter = Formatter().Make()
	var prefix = abstraction.GetPrefix()
	var identifier = abstraction.GetIdentifier()
	var arguments = abstraction.GetArguments()
	var genericIterator = genericTypes.GetSequence().GetIterator()
	var concreteIterator = concreteTypes.GetSequence().GetIterator()
	for genericIterator.HasNext() {
		var genericType = genericIterator.GetNext()
		var genericName = genericType.GetIdentifier()
		var concreteType = concreteIterator.GetNext()
		if identifier == genericName {
			identifier = formatter.FormatAbstraction(concreteType)
			break
		}
	}
	var sequence = col.List[AbstractionLike]().Make()
	if arguments != nil {
		var argumentIterator = arguments.GetSequence().GetIterator()
		for argumentIterator.HasNext() {
			var argument = argumentIterator.GetNext()
			argument = v.replaceGenericType(
				genericTypes,
				concreteTypes,
				argument,
			)
			sequence.AppendValue(argument)
		}
		arguments = Abstractions().MakeWithAttributes(sequence)
	}
	abstraction = Abstraction().MakeWithAttributes(prefix, identifier, arguments)
	return abstraction
}

func (v *generator_) replaceParameterTypes(
	genericTypes ParametersLike,
	concreteTypes ArgumentsLike,
	methodParameters ParametersLike,
) ParametersLike {
	var sequence = col.List[ParameterLike]().Make()
	var parameterIterator = methodParameters.GetSequence().GetIterator()
	for parameterIterator.HasNext() {
		var methodParameter = parameterIterator.GetNext()
		var parameterName = methodParameter.GetIdentifier()
		var parameterType = methodParameter.GetAbstraction()
		parameterType = v.replaceGenericType(
			genericTypes,
			concreteTypes,
			parameterType,
		)
		methodParameter = Parameter().MakeWithAttributes(parameterName, parameterType)
		sequence.AppendValue(methodParameter)
	}
	methodParameters = Parameters().MakeWithAttributes(sequence)
	return methodParameters
}

func (v *generator_) replaceResultTypes(
	genericTypes ParametersLike,
	concreteTypes ArgumentsLike,
	methodResult ResultLike,
) ResultLike {
	var resultAbstraction = methodResult.GetAbstraction()
	if resultAbstraction != nil {
		resultAbstraction = v.replaceGenericType(
			genericTypes,
			concreteTypes,
			resultAbstraction,
		)
		methodResult = Result().MakeWithAbstraction(resultAbstraction)
	} else {
		var resultParameters = methodResult.GetParameters()
		resultParameters = v.replaceParameterTypes(
			genericTypes,
			concreteTypes,
			resultParameters,
		)
		methodResult = Result().MakeWithParameters(resultParameters)
	}
	return methodResult
}

func (v *generator_) retrieveAspect(
	package_ PackageLike,
	identifier string,
) AspectLike {
	var iterator = package_.GetInterfaces().GetAspects().GetSequence().GetIterator()
	for iterator.HasNext() {
		var aspect = iterator.GetNext()
		var declaration = aspect.GetDeclaration()
		if declaration.GetIdentifier() == identifier {
			return aspect
		}
	}
	var message = fmt.Sprintf(
		"Missing the following aspect definition: %v",
		identifier,
	)
	panic(message)
}
