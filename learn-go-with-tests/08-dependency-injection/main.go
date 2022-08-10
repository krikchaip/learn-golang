package main

import (
	dependency_injection "08-dependency-injection/lib"
	di "08-dependency-injection/lib" // name is too long :(
)

func main() {
	dependency_injection.GreetStdout()
	di.GreetHttp()
}
