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
	sts "strings"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
	queueSize_: 16,
	stackSize_: 4,
}

// Function

func Parser() ParserClassLike {
	return parserClass
}

// CLASS METHODS

// Target

type parserClass_ struct {
	queueSize_ int
	stackSize_ int
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{
		next_: col.Stack[TokenLike]().MakeWithCapacity(c.stackSize_),
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of unprocessed retrieved tokens.
}

// Public

func (v *parser_) ParseSource(source string) PackageLike {
	// The scanner runs in a separate Go routine.
	v.source_ = sts.ReplaceAll(source, "\t", "    ") // F'ing tabs!
	v.tokens_ = col.Queue[TokenLike]().MakeWithCapacity(parserClass.queueSize_)
	Scanner().MakeFromSource(v.source_, v.tokens_)

	// Attempt to parse Package.
	var package_, token, ok = v.parsePackage()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("package",
			"$source",
			"$package",
		)
		panic(message)
	}

	// Attempt to parse the end-of-file marker.
	_, token, ok = v.parseToken(EOFToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("EOF",
			"$source",
			"$package",
		)
		panic(message)
	}

	// Found Package.
	return package_
}

// Private

/*
This private instance method returns an error message containing the context for
a parsing error.
*/
func (v *parser_) formatError(token TokenLike) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		token,
	)
	var lines = sts.Split(v.source_, "\n")
	var line = token.GetLine()

	// Append the source line with the error in it.
	message += "\033[36m"
	if line > 1 {
		message += fmt.Sprintf("%04d: ", line-1) + string(lines[line-2]) + "\n"
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count = 0
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < len(lines) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"

	return message
}

/*
This private instance method is useful when creating scanner and parser error
messages that include the required grammatical rules.
*/
func (v *parser_) generateGrammar(expected string, symbols ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, symbol := range symbols {
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			symbol,
			grammar[symbol],
		)
	}
	return message
}

/*
This private instance method attempts to read the next token from the token
stream and return it.
*/
func (v *parser_) getNextToken() TokenLike {
	var ok bool
	var token TokenLike

	// Check for any unprocessed tokens.
	if !v.next_.IsEmpty() {
		token = v.next_.RemoveTop()
		//fmt.Println(token)
		return token
	}

	// Read a new token from the token stream.
	token, ok = v.tokens_.RemoveHead()
	if !ok {
		panic("The token channel terminated without an EOF token.")
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError(token)
		panic(message)
	}
	//fmt.Println(token)
	return token
}

func (v *parser_) parseAbstraction() (
	abstraction AbstractionLike,
	token TokenLike,
	ok bool,
) {
	var arguments ArgumentsLike
	var delimiterToken, identifierToken TokenLike
	var identifier string
	var prefix PrefixLike

	// Attempt to parse an optional prefix.
	prefix, _, ok = v.parsePrefix()
	if ok {
		// Attempt to parse an identifier.
		identifier, token, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("arguments",
				"$abstraction",
				"$prefix",
				"$arguments",
			)
			panic(message)
		}
	} else {
		// Attempt to parse an identifier.
		identifier, identifierToken, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			// This is not an abstraction.
			return abstraction, identifierToken, false
		}
		_, delimiterToken, ok = v.parseToken(DelimiterToken, "(")
		if ok {
			// The identifier is the next method name not an abstraction.
			v.putBack(delimiterToken)
			v.putBack(identifierToken)
			return abstraction, identifierToken, false
		}

	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if ok {
		// Attempt to parse a sequence of arguments.
		arguments, token, ok = v.parseArguments()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("arguments",
				"$abstraction",
				"$prefix",
				"$arguments",
			)
			panic(message)
		}

		// Attempt to parse a delimiter.
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("]",
				"$abstraction",
				"$prefix",
				"$arguments",
			)
			panic(message)
		}
	}

	// Found an abstraction.
	abstraction = Abstraction().MakeWithAttributes(prefix, identifier, arguments)
	return abstraction, token, true
}

func (v *parser_) parseAbstractions() (
	abstractions AbstractionsLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Abstractions")
	if !ok {
		// This is not a sequence of abstractions.
		return abstractions, token, false
	}

	// Attempt to parse at least one abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("abstraction",
			"$abstractions",
			"$abstraction",
		)
		panic(message)
	}
	var sequence = col.List[AbstractionLike]().Make()
	for ok {
		sequence.AppendValue(abstraction)
		abstraction, token, ok = v.parseAbstraction()
	}

	// Found a sequence of abstractions.
	abstractions = Abstractions().MakeWithAttributes(sequence)
	return abstractions, token, true
}

func (v *parser_) parseArguments() (
	arguments ArgumentsLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike

	// Attempt to parse at least one abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		// This is not a sequence of arguments.
		return arguments, token, false
	}
	var sequence = col.List[AbstractionLike]().Make()
	for ok {
		sequence.AppendValue(abstraction)
		_, token, ok = v.parseToken(DelimiterToken, ",")
		if ok {
			abstraction, token, ok = v.parseAbstraction()
		}
	}

	// Found a sequence of arguments.
	arguments = Arguments().MakeWithAttributes(sequence)
	return arguments, token, true
}

func (v *parser_) parseAspect() (
	aspect AspectLike,
	token TokenLike,
	ok bool,
) {
	var declaration DeclarationLike
	var methods MethodsLike

	// Attempt to parse a declaration.
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an aspect.
		return aspect, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"interface"`,
			"$aspect",
			"$declaration",
			"$methods",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("{",
			"$aspect",
			"$declaration",
			"$methods",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of methods.
	methods, _, _ = v.parseMethods()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("}",
			"$aspect",
			"$declaration",
			"$methods",
		)
		panic(message)
	}

	// Found an aspect.
	aspect = Aspect().MakeWithAttributes(declaration, methods)
	return aspect, token, true
}

func (v *parser_) parseAspects() (
	aspects AspectsLike,
	token TokenLike,
	ok bool,
) {
	var aspect AspectLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Aspects")
	if !ok {
		// This is not a sequence of aspects.
		return aspects, token, false
	}

	// Attempt to parse at least one aspect.
	aspect, token, ok = v.parseAspect()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("aspect",
			"$aspects",
			"$aspect",
		)
		panic(message)
	}
	var sequence = col.List[AspectLike]().Make()
	for ok {
		sequence.AppendValue(aspect)
		aspect, token, ok = v.parseAspect()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The aspects must be sorted using their lowercase identifiers.
			var firstAspect = first.(AspectLike)
			var firstString = sts.ToLower(firstAspect.GetDeclaration().GetIdentifier())
			var secondAspect = second.(AspectLike)
			var secondString = sts.ToLower(secondAspect.GetDeclaration().GetIdentifier())
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of aspects.
	aspects = Aspects().MakeWithAttributes(sequence)
	return aspects, token, true
}

func (v *parser_) parseAttribute() (
	attribute AttributeLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike
	var identifier string
	var parameter ParameterLike

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a attribute.
		return attribute, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$attribute",
			"$parameter",
			"$abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional parameter.
	parameter, _, _ = v.parseParameter()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$attribute",
			"$parameter",
			"$abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional abstraction.
	abstraction, _, _ = v.parseAbstraction()

	// Found a attribute.
	attribute = Attribute().MakeWithAttributes(identifier, parameter, abstraction)
	return attribute, token, true
}

func (v *parser_) parseAttributes() (
	attributes AttributesLike,
	token TokenLike,
	ok bool,
) {
	var attribute AttributeLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Attributes")
	if !ok {
		// This is not a sequence of attributes.
		return attributes, token, false
	}

	// Attempt to parse at least one attribute.
	attribute, token, ok = v.parseAttribute()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("attribute",
			"$attributes",
			"$attribute",
		)
		panic(message)
	}
	var sequence = col.List[AttributeLike]().Make()
	for ok {
		sequence.AppendValue(attribute)
		attribute, token, ok = v.parseAttribute()
	}

	// Found a sequence of attributes.
	attributes = Attributes().MakeWithAttributes(sequence)
	return attributes, token, true
}

func (v *parser_) parseClass() (
	class ClassLike,
	token TokenLike,
	ok bool,
) {
	var constants ConstantsLike
	var constructors ConstructorsLike
	var declaration DeclarationLike
	var functions FunctionsLike

	// Attempt to parse a declaration.
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a class.
		return class, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"interface"`,
			"$class",
			"$declaration",
			"$constants",
			"$constructors",
			"$functions",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("{",
			"$class",
			"$declaration",
			"$constants",
			"$constructors",
			"$functions",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of constants.
	constants, _, _ = v.parseConstants()

	// Attempt to parse an optional sequence of constructors.
	constructors, _, _ = v.parseConstructors()

	// Attempt to parse an optional sequence of functions.
	functions, _, _ = v.parseFunctions()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("}",
			"$class",
			"$declaration",
			"$constants",
			"$constructors",
			"$functions",
		)
		panic(message)
	}

	// Found a class.
	class = Class().MakeWithAttributes(declaration, constants, constructors, functions)
	return class, token, true
}

func (v *parser_) parseClasses() (
	classes ClassesLike,
	token TokenLike,
	ok bool,
) {
	var class ClassLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Classes")
	if !ok {
		// This is not a sequence of classes.
		return classes, token, false
	}

	// Attempt to parse at least one class.
	class, token, ok = v.parseClass()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("class",
			"$classes",
			"$class",
		)
		panic(message)
	}
	var sequence = col.List[ClassLike]().Make()
	for ok {
		sequence.AppendValue(class)
		class, token, ok = v.parseClass()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The classes must be sorted using their lowercase identifiers.
			var firstClass = first.(ClassLike)
			var firstString = firstClass.GetDeclaration().GetIdentifier()
			firstString = sts.ToLower(sts.TrimSuffix(firstString, "ClassLike"))
			var secondClass = second.(ClassLike)
			var secondString = secondClass.GetDeclaration().GetIdentifier()
			secondString = sts.ToLower(sts.TrimSuffix(secondString, "ClassLike"))
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of classes.
	classes = Classes().MakeWithAttributes(sequence)
	return classes, token, true
}

func (v *parser_) parseConstant() (
	constant ConstantLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike
	var identifier string

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a constant.
		return constant, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$constant",
			"$abstraction",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$constant",
			"$abstraction",
		)
		panic(message)
	}

	// Attempt to parse an abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("abstraction",
			"$constant",
			"$abstraction",
		)
		panic(message)
	}

	// Found a constant.
	constant = Constant().MakeWithAttributes(identifier, abstraction)
	return constant, token, true
}

func (v *parser_) parseConstants() (
	constants ConstantsLike,
	token TokenLike,
	ok bool,
) {
	var constant ConstantLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Constants")
	if !ok {
		// This is not a sequence of constants.
		return constants, token, false
	}

	// Attempt to parse at least one constant.
	constant, token, ok = v.parseConstant()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("constant",
			"$constants",
			"$constant",
		)
		panic(message)
	}
	var sequence = col.List[ConstantLike]().Make()
	for ok {
		sequence.AppendValue(constant)
		constant, token, ok = v.parseConstant()
	}

	// Found a sequence of constants.
	constants = Constants().MakeWithAttributes(sequence)
	return constants, token, true
}

func (v *parser_) parseConstructor() (
	constructor ConstructorLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike
	var identifier string
	var parameters ParametersLike

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a constructor.
		return constructor, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$constructor",
			"$parameters",
			"$abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	parameters, _, _ = v.parseParameters()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$constructor",
			"$parameters",
			"$abstraction",
		)
		panic(message)
	}

	// Attempt to parse an abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("abstraction",
			"$constructor",
			"$parameters",
			"$abstraction",
		)
		panic(message)
	}

	// Found a constructor.
	constructor = Constructor().MakeWithAttributes(identifier, parameters, abstraction)
	return constructor, token, true
}

func (v *parser_) parseConstructors() (
	constructors ConstructorsLike,
	token TokenLike,
	ok bool,
) {
	var constructor ConstructorLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Constructors")
	if !ok {
		// This is not a sequence of constructors.
		return constructors, token, false
	}

	// Attempt to parse at least one constructor.
	constructor, token, ok = v.parseConstructor()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("constructor",
			"$constructors",
			"$constructor",
		)
		panic(message)
	}
	var sequence = col.List[ConstructorLike]().Make()
	for ok {
		sequence.AppendValue(constructor)
		constructor, token, ok = v.parseConstructor()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The constructors must be sorted using their lowercase identifiers.
			var firstConstructor = first.(ConstructorLike)
			var firstString = sts.ToLower(firstConstructor.GetIdentifier())
			var secondConstructor = second.(ConstructorLike)
			var secondString = sts.ToLower(secondConstructor.GetIdentifier())
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of constructors.
	constructors = Constructors().MakeWithAttributes(sequence)
	return constructors, token, true
}

func (v *parser_) parseDeclaration() (
	declaration DeclarationLike,
	token TokenLike,
	ok bool,
) {
	var comment string
	var identifier string
	var parameters ParametersLike

	// Attempt to parse a comment.
	comment, token, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not a declaration.
		return declaration, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "type")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"type"`,
			"$declaration",
			"$parameters",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("IDENTIFIER",
			"$declaration",
			"$parameters",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if ok {
		parameters, token, ok = v.parseParameters()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("parameters",
				"$declaration",
				"$parameters",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("]",
				"$declaration",
				"$parameters",
			)
			panic(message)
		}
	}

	// Found a declaration.
	declaration = Declaration().MakeWithAttributes(comment, identifier, parameters)
	return declaration, token, true
}

func (v *parser_) parseEnumeration() (
	enumeration EnumerationLike,
	token TokenLike,
	ok bool,
) {
	var values ValuesLike

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "const")
	if !ok {
		// This is not an enumeration.
		return enumeration, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$enumeration",
			"$values",
		)
		panic(message)
	}

	// Attempt to parse a sequence of values.
	values, token, ok = v.parseValues()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("values",
			"$enumeration",
			"$values",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$enumeration",
			"$values",
		)
		panic(message)
	}

	// Found an enumeration.
	enumeration = Enumeration().MakeWithAttributes(values)
	return enumeration, token, true
}

func (v *parser_) parseFunction() (
	function FunctionLike,
	token TokenLike,
	ok bool,
) {
	var identifier string
	var parameters ParametersLike
	var result ResultLike

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a function.
		return function, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$function",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	parameters, _, _ = v.parseParameters()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$function",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Attempt to parse a result.
	result, token, ok = v.parseResult()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("result",
			"$function",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Found a function.
	function = Function().MakeWithAttributes(identifier, parameters, result)
	return function, token, true
}

func (v *parser_) parseFunctions() (
	functions FunctionsLike,
	token TokenLike,
	ok bool,
) {
	var function FunctionLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Functions")
	if !ok {
		// This is not a sequence of functions.
		return functions, token, false
	}

	// Attempt to parse at least one function.
	function, token, ok = v.parseFunction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("function",
			"$functions",
			"$function",
		)
		panic(message)
	}
	var sequence = col.List[FunctionLike]().Make()
	for ok {
		sequence.AppendValue(function)
		function, token, ok = v.parseFunction()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The functions must be sorted using their lowercase identifiers.
			var firstFunction = first.(FunctionLike)
			var firstString = sts.ToLower(firstFunction.GetIdentifier())
			var secondFunction = second.(FunctionLike)
			var secondString = sts.ToLower(secondFunction.GetIdentifier())
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of functions.
	functions = Functions().MakeWithAttributes(sequence)
	return functions, token, true
}

func (v *parser_) parseFunctional() (
	functional FunctionalLike,
	token TokenLike,
	ok bool,
) {
	var declaration DeclarationLike
	var parameters ParametersLike
	var result ResultLike

	// Attempt to parse a declaration.
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a functional.
		return functional, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "func")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"func"`,
			"$functional",
			"$declaration",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$functional",
			"$declaration",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	parameters, _, _ = v.parseParameters()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$functional",
			"$declaration",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Attempt to parse a result.
	result, token, ok = v.parseResult()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("result",
			"$functional",
			"$declaration",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Found a functional.
	functional = Functional().MakeWithAttributes(declaration, parameters, result)
	return functional, token, true
}

func (v *parser_) parseFunctionals() (
	functionals FunctionalsLike,
	token TokenLike,
	ok bool,
) {
	var functional FunctionalLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Functionals")
	if !ok {
		// This is not a sequence of functionals.
		return functionals, token, false
	}

	// Attempt to parse at least one functional.
	functional, token, ok = v.parseFunctional()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("functional",
			"$functionals",
			"$functional",
		)
		panic(message)
	}
	var sequence = col.List[FunctionalLike]().Make()
	for ok {
		sequence.AppendValue(functional)
		functional, token, ok = v.parseFunctional()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The functionals must be sorted using their lowercase identifiers.
			var firstFunctional = first.(FunctionalLike)
			var firstString = sts.ToLower(firstFunctional.GetDeclaration().GetIdentifier())
			var secondFunctional = second.(FunctionalLike)
			var secondString = sts.ToLower(secondFunctional.GetDeclaration().GetIdentifier())
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of functionals.
	functionals = Functionals().MakeWithAttributes(sequence)
	return functionals, token, true
}

func (v *parser_) parseHeader() (
	header HeaderLike,
	token TokenLike,
	ok bool,
) {
	var comment string
	var identifier string

	// Attempt to parse a comment.
	comment, token, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not a header.
		return header, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "package")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"package"`,
			"$header",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"IDENTIFIER"`,
			"$header",
		)
		panic(message)
	}

	// Found a header.
	header = Header().MakeWithAttributes(comment, identifier)
	return header, token, true
}

func (v *parser_) parseImports() (
	imports ImportsLike,
	token TokenLike,
	ok bool,
) {
	var modules ModulesLike

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "import")
	if !ok {
		// This is not a sequence of imports.
		return imports, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$imports",
			"$module",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of modules.
	modules, _, _ = v.parseModules()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$imports",
			"$module",
		)
		panic(message)
	}

	// Found a sequence of imports.
	imports = Imports().MakeWithAttributes(modules)
	return imports, token, true
}

func (v *parser_) parseInstance() (
	instance InstanceLike,
	token TokenLike,
	ok bool,
) {
	var abstractions AbstractionsLike
	var attributes AttributesLike
	var declaration DeclarationLike
	var methods MethodsLike

	// Attempt to parse a declaration.
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an instance.
		return instance, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"interface"`,
			"$instance",
			"$declaration",
			"$attributes",
			"$abstractions",
			"$methods",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("{",
			"$instance",
			"$declaration",
			"$attributes",
			"$abstractions",
			"$methods",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of attributes.
	attributes, _, _ = v.parseAttributes()

	// Attempt to parse an optional sequence of abstractions.
	abstractions, _, _ = v.parseAbstractions()

	// Attempt to parse an optional sequence of methods.
	methods, _, _ = v.parseMethods()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("}",
			"$instance",
			"$declaration",
			"$attributes",
			"$abstractions",
			"$methods",
		)
		panic(message)
	}

	// Found an instance.
	instance = Instance().MakeWithAttributes(declaration, attributes, abstractions, methods)
	return instance, token, true
}

func (v *parser_) parseInstances() (
	instances InstancesLike,
	token TokenLike,
	ok bool,
) {
	var instance InstanceLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Instances")
	if !ok {
		// This is not a sequence of instances.
		return instances, token, false
	}

	// Attempt to parse at least one instance.
	instance, token, ok = v.parseInstance()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("instance",
			"$instances",
			"$instance",
		)
		panic(message)
	}
	var sequence = col.List[InstanceLike]().Make()
	for ok {
		sequence.AppendValue(instance)
		instance, token, ok = v.parseInstance()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The instances must be sorted using their lowercase identifiers.
			var firstInstance = first.(InstanceLike)
			var firstString = firstInstance.GetDeclaration().GetIdentifier()
			firstString = sts.ToLower(sts.TrimSuffix(firstString, "Like"))
			var secondInstance = second.(InstanceLike)
			var secondString = secondInstance.GetDeclaration().GetIdentifier()
			secondString = sts.ToLower(sts.TrimSuffix(secondString, "Like"))
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of instances.
	instances = Instances().MakeWithAttributes(sequence)
	return instances, token, true
}

func (v *parser_) parseInterfaces() (
	interfaces InterfacesLike,
	token TokenLike,
	ok bool,
) {
	var aspects AspectsLike
	var classes ClassesLike
	var instances InstancesLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// INTERFACES")
	if !ok {
		// This is not a sequence of interfaces.
		return interfaces, token, false
	}

	// Attempt to parse an optional sequence of aspects.
	aspects, _, _ = v.parseAspects()

	// Attempt to parse an optional sequence of classes.
	classes, _, _ = v.parseClasses()

	// Attempt to parse an optional sequence of instances.
	instances, _, _ = v.parseInstances()

	// Found a sequence of interfaces.
	interfaces = Interfaces().MakeWithAttributes(aspects, classes, instances)
	return interfaces, token, true
}

func (v *parser_) parseMethod() (
	method MethodLike,
	token TokenLike,
	ok bool,
) {
	var identifier string
	var parameters ParametersLike
	var result ResultLike

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a method.
		return method, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"$method",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	parameters, _, _ = v.parseParameters()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$method",
			"$parameters",
			"$result",
		)
		panic(message)
	}

	// Attempt to parse an optional result.
	result, _, _ = v.parseResult()

	// Found a method.
	method = Method().MakeWithAttributes(identifier, parameters, result)
	return method, token, true
}

func (v *parser_) parseMethods() (
	methods MethodsLike,
	token TokenLike,
	ok bool,
) {
	var method MethodLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Methods")
	if !ok {
		// This is not a sequence of methods.
		return methods, token, false
	}

	// Attempt to parse at least one method.
	method, token, ok = v.parseMethod()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("method",
			"$methods",
			"$method",
		)
		panic(message)
	}
	var sequence = col.List[MethodLike]().Make()
	for ok {
		sequence.AppendValue(method)
		method, token, ok = v.parseMethod()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The methods must be sorted using their lowercase identifiers.
			var firstMethod = first.(MethodLike)
			var firstString = sts.ToLower(firstMethod.GetIdentifier())
			var secondMethod = second.(MethodLike)
			var secondString = sts.ToLower(secondMethod.GetIdentifier())
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of methods.
	methods = Methods().MakeWithAttributes(sequence)
	return methods, token, true
}

func (v *parser_) parseModule() (
	module ModuleLike,
	token TokenLike,
	ok bool,
) {
	var identifier string
	var repository string

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a module.
		return module, token, false
	}

	// Attempt to parse a repository.
	repository, token, ok = v.parseToken(TextToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"TEXT"`,
			"$module",
		)
		panic(message)
	}

	// Found a module.
	module = Module().MakeWithAttributes(identifier, repository)
	return module, token, true
}

func (v *parser_) parseModules() (
	modules ModulesLike,
	token TokenLike,
	ok bool,
) {
	var module ModuleLike

	// Attempt to parse at least one module.
	module, token, ok = v.parseModule()
	if !ok {
		// This is not a sequence of modules.
		return modules, token, false
	}
	var sequence = col.List[ModuleLike]().Make()
	for ok {
		sequence.AppendValue(module)
		module, _, ok = v.parseModule()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The modules must be sorted using their repository names.
			var firstModule = first.(ModuleLike)
			var firstString = firstModule.GetRepository()
			var secondModule = second.(ModuleLike)
			var secondString = secondModule.GetRepository()
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of modules.
	modules = Modules().MakeWithAttributes(sequence)
	return modules, token, true
}

func (v *parser_) parseNotice() (
	notice NoticeLike,
	token TokenLike,
	ok bool,
) {
	var comment string

	// Attempt to parse a comment.
	comment, token, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not a notice.
		return notice, token, false
	}

	// Found a notice.
	notice = Notice().MakeWithAttributes(comment)
	return notice, token, true
}

func (v *parser_) parsePackage() (
	package_ PackageLike,
	token TokenLike,
	ok bool,
) {
	var notice NoticeLike
	var header HeaderLike
	var types TypesLike
	var interfaces InterfacesLike
	var imports ImportsLike

	// Attempt to parse a notice.
	notice, token, ok = v.parseNotice()
	if !ok {
		// This is not Package.
		return package_, token, false
	}

	// Attempt to parse a header.
	header, token, ok = v.parseHeader()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("header",
			"$package",
			"$notice",
			"$header",
			"$imports",
			"$types",
			"$interfaces",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of imports.
	imports, _, _ = v.parseImports()

	// Attempt to parse an optional sequence of types.
	types, _, _ = v.parseTypes()

	// Attempt to parse an optional sequence of interfaces.
	interfaces, _, _ = v.parseInterfaces()

	// Found a package.
	package_ = Package().MakeWithAttributes(notice, header, imports, types, interfaces)
	return package_, token, true
}

func (v *parser_) parseParameter() (
	parameter ParameterLike,
	token TokenLike,
	ok bool,
) {
	var identifier string
	var abstraction AbstractionLike

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a parameter.
		return parameter, token, false
	}

	// Attempt to parse an abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("abstraction",
			"$parameter",
			"$abstraction",
		)
		panic(message)
	}

	// Found a parameter.
	parameter = Parameter().MakeWithAttributes(identifier, abstraction)
	return parameter, token, true
}

func (v *parser_) parseParameters() (
	parameters ParametersLike,
	token TokenLike,
	ok bool,
) {
	var parameter ParameterLike

	// Attempt to parse at least one parameter.
	parameter, token, ok = v.parseParameter()
	if !ok {
		// This is not a sequence of parameters.
		return parameters, token, false
	}
	var sequence = col.List[ParameterLike]().Make()
	for ok {
		sequence.AppendValue(parameter)
		_, token, ok = v.parseToken(DelimiterToken, ",")
		if ok {
			parameter, token, ok = v.parseParameter()
		}
	}

	// Found a sequence of parameters.
	parameters = Parameters().MakeWithAttributes(sequence)
	return parameters, token, true
}

func (v *parser_) parsePrefix() (
	prefix PrefixLike,
	token TokenLike,
	ok bool,
) {
	var delimiterToken, identifierToken TokenLike
	var identifier string
	var prefixType PrefixType

	// Attempt to parse an array prefix.
	_, delimiterToken, ok = v.parseToken(DelimiterToken, "[")
	if ok {
		// Attempt to parse a delimiter.
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if ok {
			prefixType = ArrayPrefix
			prefix = Prefix().MakeWithAttributes(identifier, prefixType)
			return prefix, token, true
		}
		v.putBack(delimiterToken)
		return prefix, token, false
	}

	// Attempt to parse a map prefix.
	_, _, ok = v.parseToken(IdentifierToken, "map")
	if ok {
		_, token, ok = v.parseToken(DelimiterToken, "[")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("[",
				"$prefix",
			)
			panic(message)
		}
		identifier, token, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("IDENTIFIER",
				"$prefix",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("]",
				"$prefix",
			)
			panic(message)
		}
		prefixType = MapPrefix
		prefix = Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// Attempt to parse a channel prefix.
	_, token, ok = v.parseToken(IdentifierToken, "chan")
	if ok {
		prefixType = ChannelPrefix
		prefix = Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// Attempt to parse an alias prefix.
	identifier, identifierToken, ok = v.parseToken(IdentifierToken, "")
	if ok {
		_, token, ok = v.parseToken(DelimiterToken, ".")
		if !ok {
			v.putBack(identifierToken)
			return prefix, token, false
		}
		prefixType = AliasPrefix
		prefix = Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// This is not a prefix.
	return prefix, identifierToken, false
}

func (v *parser_) parseResult() (
	result ResultLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike
	var parameters ParametersLike

	// Attempt to parse an abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if ok {
		// Found an abstraction result.
		result = Result().MakeWithAbstraction(abstraction)
		return result, token, true
	}

	// Attempt to parse a sequence of parameters.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if ok {
		parameters, token, ok = v.parseParameters()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("parameters",
				"$result",
				"$abstraction",
				"$parameters",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, ")")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar(")",
				"$result",
				"$abstraction",
				"$parameters",
			)
			panic(message)
		}

		// Found a named parameters result.
		result = Result().MakeWithParameters(parameters)
		return result, token, true
	}

	// This is not a result.
	return result, token, false
}

func (v *parser_) parseSpecialization() (
	specialization SpecializationLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike
	var declaration DeclarationLike
	var enumeration EnumerationLike

	// Attempt to parse a declaration.
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a specialization.
		return specialization, token, false
	}

	// Attempt to parse an abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("abstraction",
			"$specialization",
			"$declaration",
			"$abstraction",
			"$enumeration",
		)
		panic(message)
	}

	// Attempt to parse an optional enumeration.
	enumeration, token, _ = v.parseEnumeration()

	// Found a specialization.
	specialization = Specialization().MakeWithAttributes(declaration, abstraction, enumeration)
	return specialization, token, true
}

func (v *parser_) parseSpecializations() (
	specializations SpecializationsLike,
	token TokenLike,
	ok bool,
) {
	var specialization SpecializationLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Specializations")
	if !ok {
		// This is not a sequence of specializations.
		return specializations, token, false
	}

	// Attempt to parse at least one specialization.
	specialization, token, ok = v.parseSpecialization()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("specialization",
			"$specializations",
			"$specialization",
		)
		panic(message)
	}
	var sequence = col.List[SpecializationLike]().Make()
	for ok {
		sequence.AppendValue(specialization)
		specialization, token, ok = v.parseSpecialization()
	}
	sequence.SortValuesWithRanker(
		func(first, second col.Value) int {
			// The specializations must be sorted using their lowercase identifiers.
			var firstSpecialization = first.(SpecializationLike)
			var firstString = sts.ToLower(firstSpecialization.GetDeclaration().GetIdentifier())
			var secondSpecialization = second.(SpecializationLike)
			var secondString = sts.ToLower(secondSpecialization.GetDeclaration().GetIdentifier())
			switch {
			case firstString < secondString:
				return -1
			case firstString > secondString:
				return 1
			default:
				return 0
			}
		},
	)

	// Found a sequence of specializations.
	specializations = Specializations().MakeWithAttributes(sequence)
	return specializations, token, true
}

func (v *parser_) parseToken(expectedType TokenType, expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token.
	token = v.getNextToken()
	value = token.GetValue()
	if token.GetType() == expectedType {
		var constrained = len(expectedValue) > 0
		if !constrained || value == expectedValue {
			// Found the expected token.
			return value, token, true
		}
	}

	// This is not the right token.
	v.putBack(token)
	return "", token, false
}

func (v *parser_) parseTypes() (
	types TypesLike,
	token TokenLike,
	ok bool,
) {
	var functionals FunctionalsLike
	var specializations SpecializationsLike

	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// TYPES")
	if !ok {
		// This is not a sequence of types.
		return types, token, false
	}

	// Attempt to parse an optional sequence of specializations.
	specializations, _, _ = v.parseSpecializations()

	// Attempt to parse an optional sequence of functionals.
	functionals, _, _ = v.parseFunctionals()

	// Found a sequence of types.
	types = Types().MakeWithAttributes(specializations, functionals)
	return types, token, true
}

func (v *parser_) parseValues() (
	values ValuesLike,
	token TokenLike,
	ok bool,
) {
	var abstraction AbstractionLike
	var identifier string

	// Attempt to parse an identifier.
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a sequence of values.
		return values, token, false
	}

	// Attempt to parse an abstraction.
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("values",
			"$values",
			"$parameter",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "=")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("=",
			"$values",
			"$parameter",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	_, token, ok = v.parseToken(IdentifierToken, "iota")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("iota",
			"$values",
			"$parameter",
		)
		panic(message)
	}

	// Attempt to parse a sequence of identifiers.
	var sequence = col.List[string]().Make()
	for ok {
		sequence.AppendValue(identifier)
		identifier, token, ok = v.parseToken(IdentifierToken, "")
	}

	// Found a sequence of values.
	values = Values().MakeWithAttributes(sequence, abstraction)
	return values, token, true
}

func (v *parser_) putBack(token TokenLike) {
	//fmt.Printf("Put Back %v\n", token)
	v.next_.AddValue(token)
}

var grammar = map[string]string{
	"$abstraction":     `prefix? IDENTIFIER ("[" arguments "]")?`,
	"$abstractions":    `"// Abstractions" abstraction+`,
	"$arguments":       `abstraction ("," abstraction)* ","?`,
	"$aspect":          `declaration "interface" "{" methods? "}"`,
	"$aspects":         `"// Aspects" aspect+`,
	"$attribute":       `IDENTIFIER "(" parameter? ")" abstraction?`,
	"$attributes":      `"// Attributes" attribute+`,
	"$class":           `declaration "interface" "{" constants? constructors? functions? "}"`,
	"$classes":         `"// Classes" class+`,
	"$constant":        `IDENTIFIER "(" ")" abstraction`,
	"$constants":       `"// Constants" constant+`,
	"$constructor":     `IDENTIFIER "(" parameters? ")" abstraction`,
	"$constructors":    `"// Constructors" constructor+`,
	"$declaration":     `COMMENT "type" IDENTIFIER ("[" parameters "]")?`,
	"$enumeration":     `"const" "(" values ")"`,
	"$function":        `IDENTIFIER "(" parameters? ")" result`,
	"$functional":      `declaration "func" "(" parameters? ")" result`,
	"$functionals":     `"// Functionals" functional+`,
	"$functions":       `"// Functions" function+`,
	"$header":          `COMMENT "package" IDENTIFIER`,
	"$imports":         `"import" "(" modules? ")"`,
	"$instance":        `declaration "interface" "{" attributes? abstractions? methods? "}"`,
	"$instances":       `"// Instances" instance+`,
	"$interfaces":      `"// INTERFACES" aspects? classes? instances?`,
	"$method":          `IDENTIFIER "(" parameters? ")" result?`,
	"$methods":         `"// Methods" method+`,
	"$module":          `IDENTIFIER TEXT`,
	"$modules":         `module+`,
	"$notice":          `COMMENT`,
	"$package":         `notice header imports? types? interfaces?`,
	"$parameter":       `IDENTIFIER abstraction`,
	"$parameters":      `parameter ("," parameter)* ","?`,
	"$prefix":          `"[" "]" | "map" "[" IDENTIFIER "]" | "chan" | IDENTIFIER "."`,
	"$result":          `abstraction | "(" parameters ")"`,
	"$source":          `package EOF  ! Terminated with an end-of-file marker.`,
	"$specialization":  `declaration abstraction enumeration?`,
	"$specializations": `"// Specializations" specialization+`,
	"$types":           `"// TYPES" specializations? functionals?`,
	"$values":          `parameter "=" "iota" IDENTIFIER*`,
}
