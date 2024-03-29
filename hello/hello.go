// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a trivial example of a main package.
package main

import (
	"fmt"

	//"code.google.com/p/go.example/newmath"
	// error msg:
	/*
	$ go get github.com/aspes/go.example/hello
	go: missing Mercurial command. See http://golang.org/s/gogetcmd
	package github.com/aspes/go.example/hello
        imports code.google.com/p/go.example/newmath: exec: "hg": executable file not found in %PATH%
        */
	
	//"newmath"
	// error msg:
	/* 
	$ go get github.com/aspes/go.example/hello
	package github.com/aspes/go.example/hello
        	imports newmath: unrecognized import path "newmath"
	*/
	
	"github.com/aspes/go.example/newmath"
	// now remote package install ok
	// but be careful! go get checks local copy first.
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))
}
