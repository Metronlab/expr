package main

import (
	"fmt"
	"github.com/metronlab/expr/constants"
	"go/format"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var data string
	echo := func(s string, xs ...interface{}) {
		data += fmt.Sprintf(s, xs...) + "\n"
	}

	echo(`// Code generated by vm/generate/main.go. DO NOT EDIT.`)
	echo(``)
	echo(`package vm`)
	echo(`import (`)
	echo(`"fmt"`)
	echo(`"reflect"`)
	echo(`)`)

	types := []string{
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"float32",
		"float64",
	}

	helpers := []struct {
		name, op        string
		noFloat, string bool
	}{
		{
			name:   "equal",
			op:     constants.OpEqual,
			string: true,
		},
		{
			name:   "less",
			op:     constants.OpLess,
			string: true,
		},
		{
			name:   "more",
			op:     constants.OpGreater,
			string: true,
		},
		{
			name:   "lessOrEqual",
			op:     constants.OpLessOrEqual,
			string: true,
		},
		{
			name:   "moreOrEqual",
			op:     constants.OpGreaterOrEqual,
			string: true,
		},
		{
			name:   "add",
			op:     constants.OpAdd,
			string: true,
		},
		{
			name: "subtract",
			op:   constants.OpSubtract,
		},
		{
			name: "multiply",
			op:   constants.OpMultiply,
		},
		{
			name: "divide",
			op:   constants.OpDivide,
		},
		{
			name:    "modulo",
			op:      constants.OpModulo,
			noFloat: true,
		},
		{
			name:    "bitwiseAnd",
			op:      constants.OpBitwiseAnd,
			noFloat: true,
		},
		{
			name:    "bitwiseOr",
			op:      constants.OpBitwiseOr,
			noFloat: true,
		},
		{
			name:    "bitwiseXor",
			op:      constants.OpBitwiseXor,
			noFloat: true,
		},
		{
			name:    "bitwiseLeftShift",
			op:      constants.OpBitwiseLShift,
			noFloat: true,
		},
		{
			name:    "bitwiseRightShift",
			op:      constants.OpBitwiseRShift,
			noFloat: true,
		},
	}

	for _, helper := range helpers {
		name := helper.name
		op := helper.op
		echo(`func %v(a, b interface{}) interface{} {`, name)
		echo(`switch x := a.(type) {`)
		for i, a := range types {
			if helper.noFloat && strings.HasPrefix(a, "float") {
				continue
			}
			echo(`case %v:`, a)
			echo(`switch y := b.(type) {`)
			for j, b := range types {
				if helper.noFloat && strings.HasPrefix(b, "float") {
					continue
				}
				echo(`case %v:`, b)
				if i == j {
					if op == constants.OpBitwiseAnd {
						echo(`return x & y`)
					} else if op == constants.OpBitwiseOr {
						echo(`return x | y`)
					} else if op == constants.OpBitwiseXor {
						echo(`return x ^ y`)
					} else if op == constants.OpBitwiseLShift {
						echo(`return x << y`)
					} else if op == constants.OpBitwiseRShift {
						echo(`return x >> y`)
					} else {
						echo(`return x %v y`, op)
					}
				}
				if i < j {
					if op == constants.OpBitwiseAnd {
						echo(`return %v(x) & y`, b)
					} else if op == constants.OpBitwiseOr {
						echo(`return %v(x) | y`, b)
					} else if op == constants.OpBitwiseXor {
						echo(`return %v(x) ^ y`, b)
					} else if op == constants.OpBitwiseLShift {
						echo(`return %v(x) << y`, b)
					} else if op == constants.OpBitwiseRShift {
						echo(`return %v(x) >> y`, b)
					} else {
						echo(`return %v(x) %v y`, b, op)
					}
				}
				if i > j {
					if op == constants.OpBitwiseAnd {
						echo(`return x & %v(y)`, a)
					} else if op == constants.OpBitwiseOr {
						echo(`return x | %v(y)`, a)
					} else if op == constants.OpBitwiseXor {
						echo(`return x ^ %v(y)`, a)
					} else if op == constants.OpBitwiseLShift {
						echo(`return x << %v(y)`, a)
					} else if op == constants.OpBitwiseRShift {
						echo(`return x >> %v(y)`, a)
					} else {
						echo(`return x %v %v(y)`, op, a)
					}
				}
			}
			echo(`}`)
		}
		if helper.string {
			echo(`case string:`)
			echo(`switch y := b.(type) {`)
			echo(`case string: return x %v y`, op)
			echo(`}`)
		}
		echo(`}`)
		if name == "equal" {
			echo(`if isNil(a) && isNil(b) { return true }`)
			echo(`return reflect.DeepEqual(a, b)`)
		} else {
			echo(`panic(fmt.Sprintf("invalid operation: %%T %%v %%T", a, "%v", b))`, op)
		}
		echo(`}`)
		echo(``)
	}

	b, err := format.Source([]byte(data))
	check(err)
	err = ioutil.WriteFile("helpers.go", b, 0644)
	check(err)
}
