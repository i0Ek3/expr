package expr

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(pi * pi)", Env{"pi": math.Pi}, "3.14159"},
		{"sqrt(four * nine)", Env{"four": 4, "nine": 9}, "6"},

		{"pow(x, 3) + pow(y, 3)", Env{"x": 3, "y": 4}, "91"},
		{"pow(x, 1) + pow(y, 2)", Env{"x": 3, "y": 4}, "19"},

		{"log(x)", Env{"x": 2}, "0.693147"},
		{"log(x, y)", Env{"x": 2, "y": 4}, "0.693147"},

		{"sin(pi)", Env{"pi": math.Pi}, "1.22465e-16"},
		{"cos(pi)", Env{"pi": math.Pi}, "-1"},
		{"tan(pi)", Env{"pi": math.Pi}, "-1.22465e-16"},

		{"5/9 * (F-32)", Env{"F": 41}, "5"},
		{"5/9 * (F-32)", Env{"F": 32}, "0"},
		{"5/9 * (F-32)", Env{"F": 212}, "100"},

		{"-x", Env{"x": 1}, "-1"},
		{"-x", Env{"x": -1}, "1"},
		{"+x", Env{"x": 1}, "1"},
		{"+x", Env{"x": -1}, "-1"},
		{"!x", Env{"x": 1}, "0"},
		{"!x", Env{"x": -1}, "0"},

		// unsupported function
		//{"asin(pi)", Env{"pi": math.Pi}, "0"},
	}

	var preExpr string
	for _, test := range tests {
		if test.expr != preExpr {
			fmt.Printf("\n%s\n", test.expr)
			preExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		//asrt.Asrt(t, got, test.want)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, but want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
