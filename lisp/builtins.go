package lisp

import (
	"fmt"
	"io"
	"os"
)

type function func(*scope, []sexpr) sexpr

// Circumvent lame initialization loop detection. An explicit init() allows
// builtinDefine et al to reference global.
func init() {
	globalData := map[sym]sexpr{
		// Misc. primitives (primitives.go)
		"if":     primitive(primitiveIf),
		"for":    primitive(primitiveFor),
		"lambda": primitive(primitiveLambda),
		"let":    primitive(primitiveLet),
		"define": primitive(primitiveDefine),
		"quote":  primitive(primitiveQuote),
		"begin":  primitive(primitiveBegin),

		// Nil
		"nil": Nil,

		// Misc
		"read":  function(builtinRead),
		"eval":  function(builtinEval),
		"apply": function(builtinApply),
		"print": function(builtinPrint),

		// Cons manipulation (cons.go)
		"cons": function(builtinCons),
		"car":  function(builtinCar),
		"cdr":  function(builtinCdr),

		// Boolean arithmetic (bool.go)
		"not": function(builtinNot),

		// Arithmetic (math.go)
		"=":  function(builtinEq),
		"/=": function(builtinNe),
		"<":  function(builtinLt),
		">":  function(builtinGt),
		"<=": function(builtinLe),
		">=": function(builtinGe),
		"+":  function(builtinAdd),
		"-":  function(builtinSub),
		"*":  function(builtinMul),
		"/":  function(builtinDiv),
		"%":  function(builtinMod),

		// Go runtime (compat.go)
		"import": function(builtinImport),

		// Panics (panic.go)
		"recover": function(builtinRecover),
		"panic":   function(builtinPanic),
	}

	global = &scope{globalData, nil}
}

// (read)
//
// Reads one s-expression from standard input.
func builtinRead(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 0 {
		panic("Invalid number of arguments")
	}
	v, err := parse(GetRuneScanner(os.Stdin))
	if err != nil && err != io.EOF {
		panic(err)
	} else if err == io.EOF {
		panic(sym("eof"))
	}
	return v
}

// (eval expr)
//
// Evaluates an s-expression.
func builtinEval(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 1 {
		panic("Invalid number of arguments")
	}
	return eval(sc, ss[0]) // TODO custom scope
}

// (print expr)
//
// Prints an s-expression.
func builtinPrint(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 1 {
		panic("Invalid number of arguments")
	}
	fmt.Printf("%s\n", asString(ss[0]))
	return Nil
}

// (apply func '(arg1 ...))
func builtinApply(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	return apply(sc, ss[0], flatten(ss[1]))
}
