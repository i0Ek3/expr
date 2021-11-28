# expr

Expressor in Go.

## Getting Started

### Install

`go get https://github.com/i0Ek3/expr`

### Usage

```Go
package main

import (
    "github.com/i0Ek3/expr"
)

func main() {
    // ...
	var env  Env
    var vars map[Var]bool

    expr.Eval(env)
    expr.Check(vars)

    // ...
}

```

## Credit

File parse.go modified from gopl.io.
