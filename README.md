<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Package Framework

### Overview
This project provides a class-based framework for jump-starting Go package
development.  The class-based model follows the Crater Dog Technologies™
(i.e. craterdog)
[Go Coding Conventions](https://github.com/craterdog/go-package-framework/wiki).

### Getting Started
The general development process—in a nutshell—is as follows:
 1. Install the
    [go-package-tools](https://github.com/craterdog/go-package-tools) module.
 1. Run the `bin/generate` program to generate a `Package.go` class model
    template file in your package directory.
 1. Fill in the `Package.go` class model template with the abstract types and
    interfaces for the classes that your package will provide.
 1. Run the `bin/generate` program again to generate a concrete class file for each
    of the corresponding abstract classes defined in your `Package.go` file.
 1. Insert the method implementations for the class methods and instance methods
    associated with each concrete class.

### Contributing
Project contributors are always welcome. Check out the contributing guidelines
[here](https://github.com/craterdog/go-package-framework/blob/main/.github/CONTRIBUTING.md).

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
