package main

import (
	"08-dependency-injection/lib"    // using a full package name
	di "08-dependency-injection/lib" // using another alias
)

func main() {
	dependency_injection.GreetStdout()
	di.GreetHttp()
}
