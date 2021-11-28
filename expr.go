package expr

import (
    "fmt"
    "math"
    "strings"
)

type Expr interface {
    // Eval returns the value on the context Env
    Eval(env Env) float64

    // Check checks the error in the expression,
    // and add them into map vars
    Check(vars map[Var]bool) error
}

// Var indicates a variable
type Var string

// literal indicates a literal
type literal float64

// Env indicates a context to map Var to float64 
type Env map[Var]float64

// unary indicates `op x`, which op represents +-!
type unary struct {
    op rune
    x  Expr
}

// binary indicates `x op y`, which op represents +-*/
type binary struct {
    op   rune
    x, y Expr
}

// call indicates the expression of function call
type call struct {
    fn   string
    args []Expr
}

func (v Var) Eval(env Env) float64 {
    return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
    vars[v] = true
    return nil
}

func (l literal) Eval(env Env) float64 {
    return float64(l)
}

func (_ literal) Check(vars map[Var]bool) error {
    return nil
}

func (u unary) Eval(env Env) float64 {
    switch u.op {
    case '+':
        return +u.x.Eval(env)
    case '-':
        return -u.x.Eval(env)
    case '!':
        return 0
    default:
        panic(fmt.Sprintf("Unsupported unary operator: %q", u.op))
    }
}

func (u unary) Check(vars map[Var]bool) error {
    if !strings.ContainsRune("+-!", u.op) {
        return fmt.Errorf("Unexpected unary op %q", u.op)
    }
    return u.x.Check(vars)
}

func (b binary) Eval(env Env) float64 {
    switch b.op {
    case '+':
        return b.x.Eval(env) + b.y.Eval(env)
    case '-':
        return b.x.Eval(env) - b.y.Eval(env)
    case '*':
        return b.x.Eval(env) * b.y.Eval(env)
    case '/':
        return b.x.Eval(env) / b.y.Eval(env)
    default:
        panic(fmt.Sprintf("Unsupported binary operator: %q", b.op))
    }
}

func (b binary) Check(vars map[Var]bool) error {
    if !strings.ContainsRune("+-*/", b.op) {
        return fmt.Errorf("Unexpected binary op %q", b.op)
    }
    if err := b.x.Check(vars); err != nil {
        return err
    }
    return b.y.Check(vars)
}

func (c call) Eval(env Env) float64 {
    switch c.fn {
    case "pow":
        return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
    case "log":
        return math.Log(c.args[0].Eval(env))
    case "sqrt":
        return math.Sqrt(c.args[0].Eval(env))
    case "sin":
        return math.Sin(c.args[0].Eval(env))
    case "cos":
        return math.Cos(c.args[0].Eval(env))
    case "tan":
        return math.Tan(c.args[0].Eval(env))
    default:
        panic(fmt.Sprintf("Unsupported function call: %s", c.fn))
    }
}

var params = map[string]int{"pow": 2, "log": 1, "sqrt": 1, "sin": 1, "cos": 1, "tan": 1}

func (c call) Check(vars map[Var]bool) error {
    arity, ok := params[c.fn]
    if !ok {
        return fmt.Errorf("Unknown function %q", c.fn)
    }
    if len(c.args) != arity {
        return fmt.Errorf("Call to %s has %d args, want %d", c.fn, len(c.args), arity)
    }
    for _, arg := range c.args {
        if err := arg.Check(vars); err != nil {
            return err
        }
    }
    return nil
}
